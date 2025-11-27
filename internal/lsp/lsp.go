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

		diag, err := s.HandleTextDocumentDidChange(ctx, &params)

		if diag != nil {
			protocolDiagnostics := []protocol.Diagnostic{}

			for _, v := range *diag {
				protocolDiagnostics = append(protocolDiagnostics, protocol.Diagnostic{
					Source: "nomad-ls",
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      uint32(v.Subject.Start.Line - 1),
							Character: uint32(v.Subject.Start.Column - 1),
						},
						End: protocol.Position{
							Line:      uint32(v.Subject.End.Line - 1),
							Character: uint32(v.Subject.End.Column - 1),
						},
					},
					Message: v.Detail,
				})
			}

			log.Printf("diagnostics: %+v", protocolDiagnostics)
			s.con.Notify(context.Background(), "textDocument/publishDiagnostics", protocol.PublishDiagnosticsParams{
				URI:         params.TextDocument.URI,
				Version:     uint32(params.TextDocument.Version),
				Diagnostics: protocolDiagnostics,
			})
		}

		return nil, err
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
	langSchema *hclschema.BodySchema,
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
					return langSchema.Blocks[k].Description.Value
				}
				continue
			}

			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				ans = dfs(b.Body, schemaMap, pos, schemaMap[k], langSchema.Blocks[k].Body)
			}
		}
	}

	for k, v := range bodyContent.Attributes {
		if v.NameRange.ContainsPos(pos) {
			return langSchema.Attributes[k].Description.Value
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

	var matchingBlocks uint

	for k, v := range blocksByType {
		for _, b := range v {
			blockRange := b.Body.(*hclsyntax.Body).SrcRange
			if !blockRange.ContainsPos(pos) {
				continue
			}

			matchingBlocks += 1

			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				dfs2(b.Body, blocks, schemaMap, pos, schemaMap[k], langSchema.Blocks[k].Body)
			}
		}
	}

	if matchingBlocks == 0 {
		var blocksByTypeArr []protocol.CompletionItem

		for k, v := range langSchema.Blocks {
			if len(v.Labels) != 0 {
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: asBlock(k),
					Kind:       protocol.CompletionItemKindInterface,
					// Kind:       protocol.CompletionItemKindClass,
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			} else {
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: asAnonymousBlock(k),
					Kind:       protocol.CompletionItemKindInterface,
					// Kind:       protocol.CompletionItemKindClass,
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			}
		}

		for k, v := range langSchema.Attributes {
			if v.Constraint == nil {
				continue
			}

			h := v.Constraint.(*hclschema.LiteralType)

			if h == nil {
				continue
			}

			log.Printf("attr: %s", k)
			log.Printf("%+v", bodyContent.Attributes)

			if bodyContent.Attributes[k] != nil {
				continue
			}

			switch h.Type {
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

func CollectDiagnistics(body hcl.Body, schemaMap map[string]*hcl.BodySchema) *hcl.Diagnostics {
	var diags hcl.Diagnostics

	diags = diags.Extend(dfs3(body, &diags, schemaMap, schema.SchemaMapBetter["root"], &schema.RootBodySchema))

	return &diags
}

func dfs3(body hcl.Body, diags *hcl.Diagnostics, schemaMap map[string]*hcl.BodySchema, currSchema *hcl.BodySchema, langSchema *hclschema.BodySchema) hcl.Diagnostics {
	if currSchema == nil {
		return make(hcl.Diagnostics, 0)
	}

	bodyContent, allDiags := body.Content(currSchema)
	blocksByType := bodyContent.Blocks.ByType()

	for k, v := range blocksByType {
		for _, b := range v {
			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				allDiags = allDiags.Extend(dfs3(b.Body, diags, schemaMap, schemaMap[k], langSchema.Blocks[k].Body))
			} else if langSchema.Blocks[k] != nil && langSchema.Blocks[k].DependentBody != nil {
				log.Printf("found config!")
				if bodyContent.Attributes["driver"] != nil {
					driver, _ := bodyContent.Attributes["driver"].Expr.Value(&hcl.EvalContext{})

					log.Printf("driver: %s", driver.AsString())

					schemaMapDependentKey := fmt.Sprintf("%s:%s", k, driver.AsString())

					log.Printf("map key: %s", schemaMapDependentKey)

					// log.Printf("driver: %+v", bodyContent.Attributes["xd"].)
					// langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(bodyContent.Attributes["driver"].Name)]
					allDiags = allDiags.Extend(dfs3(b.Body, diags, schemaMap, schemaMap[schemaMapDependentKey], langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(driver.AsString())]))
				}
			}
		}
	}

	return allDiags
}
