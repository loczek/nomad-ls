package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"unicode/utf8"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/parser"
	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/zclconf/go-cty/cty"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

type Service struct {
	con       jsonrpc2.Conn
	parser    parser.Parser
	schemaMap map[string]*hcl.BodySchema
	logger    slog.Logger
}

func New(con jsonrpc2.Conn, logger slog.Logger) Service {
	return Service{
		con:       con,
		parser:    *parser.NewParser(),
		schemaMap: schema.SchemaMapBetter,
		logger:    logger,
	}
}

func (s *Service) Handle(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) (any, error) {
	switch req.Method() {
	case protocol.MethodInitialize:
		params := protocol.InitializedParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		return s.HandleInitialize(ctx, &params)
	case protocol.MethodTextDocumentHover:
		params := protocol.HoverParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		s.logger.Info(fmt.Sprintf("%+v", params))

		return s.HandleTextDocumentHover(ctx, &params)
	case protocol.MethodTextDocumentCompletion:
		params := protocol.CompletionParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		s.logger.Info(fmt.Sprintf("%+v", params))

		return s.HandleTextDocumentCompletion(ctx, &params)
	case protocol.MethodTextDocumentDidOpen:
		params := protocol.DidOpenTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		return nil, s.HandleTextDocumentDidOpen(ctx, &params)
	case protocol.MethodTextDocumentDidChange:
		params := protocol.DidChangeTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		return nil, s.HandleTextDocumentDidChange(ctx, &params)
	case protocol.MethodTextDocumentDidClose:
		params := protocol.DidCloseTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		return nil, s.HandleTextDocumentDidClose(ctx, &params)
	case protocol.MethodShutdown:
		ctx.Done()
		return nil, nil
	}

	return nil, nil
}

func CollectHoverInfo(body hcl.Body, pos hcl.Pos, schemaMap map[string]*hcl.BodySchema) []string {
	return []string{dfs(body, schemaMap, pos, schema.SchemaMapBetter["root"], &schema.RootBodySchema)}
}

func dfs(
	body hcl.Body,
	schemaMap map[string]*hcl.BodySchema,
	pos hcl.Pos,
	currSchema *hcl.BodySchema,
	nonHCLSchema *hclschema.BodySchema,
) string {
	if currSchema == nil {
		return ""
	}

	bodyContent, _ := body.Content(currSchema)
	blocksByType := bodyContent.Blocks.ByType()

	ans := ""

	for k, v := range blocksByType {
		for _, b := range v {
			blockRange := b.Body.(*hclsyntax.Body).SrcRange
			if !blockRange.ContainsPos(pos) {
				blockRange := b.TypeRange
				if blockRange.ContainsPos(pos) {
					return nonHCLSchema.Blocks[k].Description.Value
				}
				continue
			}

			if nonHCLSchema.Blocks[k] != nil && nonHCLSchema.Blocks[k].Body != nil {
				ans = dfs(b.Body, schemaMap, pos, schemaMap[k], nonHCLSchema.Blocks[k].Body)
			}
		}
	}

	for k, v := range bodyContent.Attributes {
		if v.NameRange.ContainsPos(pos) {
			return nonHCLSchema.Attributes[k].Description.Value
		}
	}

	return ans
}

func CollectCompletions(body hcl.Body, pos hcl.Pos, schemaMap map[string]*hcl.BodySchema) []protocol.CompletionItem {
	var blocks []protocol.CompletionItem

	dfs2(body, &blocks, schemaMap, pos, schema.SchemaMapBetter["root"], &schema.RootBodySchema)

	return blocks
}

func dfs2(body hcl.Body, blocks *[]protocol.CompletionItem, schemaMap map[string]*hcl.BodySchema, pos hcl.Pos, currSchema *hcl.BodySchema, langSchema *hclschema.BodySchema) {
	if currSchema == nil {
		return
	}

	bodyContent, _ := body.Content(currSchema)
	blocksByType := bodyContent.Blocks.ByType()
	log.Printf("block by type: %#v", blocksByType)

	var matchingBlocks uint

	for k, v := range blocksByType {
		for _, b := range v {
			blockRange := b.Body.(*hclsyntax.Body).SrcRange
			if !blockRange.ContainsPos(pos) {
				continue
			}

			matchingBlocks += 1

			// if langSchema.Blocks[k].Body != nil {

			// 	// for _, z := range langSchema.Blocks {
			// 	// 	arr = append(arr, )
			// 	// }
			// 	arr = append(arr, langSchema.BlockTypes()...)
			// }

			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				dfs2(b.Body, blocks, schemaMap, pos, schemaMap[k], langSchema.Blocks[k].Body)
			}
		}
	}

	if matchingBlocks == 0 {
		var blocksByTypeArr []protocol.CompletionItem

		for k := range langSchema.Blocks {
			blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
				Label:      k,
				InsertText: asAnonymousBlock(k),
				Kind:       protocol.CompletionItemKindInterface,
				// Kind:       protocol.CompletionItemKindClass,
				InsertTextFormat: protocol.InsertTextFormatSnippet,
			})
		}

		for k, v := range langSchema.Attributes {
			if v.DefaultValue == nil {
				continue
			}

			z := v.DefaultValue.(*hclschema.DefaultValue)

			if z == nil {
				continue
			}

			switch z.Value.Type() {
			case cty.String:
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: fmt.Sprintf("%s = \"$0\"", k),
					Kind:       protocol.CompletionItemKindVariable,
					Documentation: protocol.MarkupContent{
						Kind:  protocol.Markdown,
						Value: v.Description.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			case cty.List(cty.String):
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: fmt.Sprintf("%s = [\"$0\"]", k),
					Kind:       protocol.CompletionItemKindVariable,
					Documentation: protocol.MarkupContent{
						Kind:  protocol.Markdown,
						Value: v.Description.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			default:
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: fmt.Sprintf("%s = ", k),
					Kind:       protocol.CompletionItemKindVariable,
					Documentation: protocol.MarkupContent{
						Kind:  protocol.Markdown,
						Value: v.Description.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			}
		}

		// *blocks = append(*blocks, langSchema.AttributeNames()...)
		*blocks = append(*blocks, blocksByTypeArr...)
	}

	log.Printf("matching blocks: %d", matchingBlocks)
}

func asBlock(name string) string {
	return fmt.Sprintf("%s \"$1\" {\n\t$0\n}", name)
}

func asAnonymousBlock(name string) string {
	return fmt.Sprintf("%s {\n\t$0\n}", name)
}

func CalculateByteOffset(pos protocol.Position, src []byte) uint {
	runes := []rune(string(src))

	var runeIndex uint
	var line uint
	var bytesCount uint

	for line < uint(pos.Line) && runeIndex < uint(len(runes)) {
		if runes[runeIndex] == '\n' {
			line += 1
		}
		bytesCount += uint(utf8.RuneLen(runes[runeIndex]))
		runeIndex += 1
	}

	var j uint

	for j < uint(pos.Character) && runeIndex < uint(len(runes)) {
		bytesCount += uint(utf8.RuneLen(runes[runeIndex]))
		runeIndex += 1
		j += 1
	}

	return bytesCount
}
