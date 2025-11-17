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
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/zclconf/go-cty/cty"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

type Service struct {
	con       jsonrpc2.Conn
	parser    hclparse.Parser
	schemaMap map[string]*hcl.BodySchema
	logger    slog.Logger
}

func New(con jsonrpc2.Conn, logger slog.Logger) Service {
	return Service{
		con:       con,
		parser:    *hclparse.NewParser(),
		schemaMap: schema.SchemaMapBetter,
		logger:    logger,
	}
}

func (s *Service) Handle(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) {
	switch req.Method() {
	case protocol.MethodInitialize:
		reply(ctx, protocol.InitializeResult{
			ServerInfo: &protocol.ServerInfo{
				Name:    "nomad-ls",
				Version: "0.0.1",
			},
			Capabilities: protocol.ServerCapabilities{
				CompletionProvider: &protocol.CompletionOptions{},
				HoverProvider:      &protocol.HoverOptions{},
				TextDocumentSync: &protocol.TextDocumentSyncOptions{
					Change: protocol.TextDocumentSyncKindFull,
				},
			},
		}, nil)
	case protocol.MethodTextDocumentHover:
		params := protocol.HoverParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			reply(ctx, nil, err)
			return
		}

		s.logger.Info(fmt.Sprintf("%+v", params))

		file := s.parser.Files()[params.TextDocument.URI.Filename()]

		body := file.Body

		byteOffset := CalculateByteOffset(params.Position, s.parser.Files()[params.TextDocument.URI.Filename()].Bytes)

		pos := hcl.InitialPos
		pos.Byte = int(byteOffset)

		x := CollectHoverInfo(body, hcl.Pos{
			Line:   int(params.Position.Line),
			Column: int(params.Position.Character),
			Byte:   pos.Byte,
		}, s.schemaMap)

		s.logger.Info(fmt.Sprintf("arr: %v", x))

		if len(x) == 0 {
			reply(ctx, nil, nil)
			return
		}

		reply(ctx, protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.PlainText,
				Value: fmt.Sprintf("%s", x[len(x)-1]),
			},
			// Range: &protocol.Range{Start: protocol.Position{Line: 0, Character: 0}, End: protocol.Position{Line: 0, Character: 0}},
		}, nil)
	case protocol.MethodTextDocumentCompletion:
		params := protocol.CompletionParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			reply(ctx, nil, err)
			return
		}

		s.logger.Info(fmt.Sprintf("%+v", params))

		file := s.parser.Files()[params.TextDocument.URI.Filename()]

		body := file.Body

		byteOffset := CalculateByteOffset(params.Position, s.parser.Files()[params.TextDocument.URI.Filename()].Bytes)

		pos := hcl.InitialPos
		pos.Byte = int(byteOffset)

		completions := CollectCompletions(body, hcl.Pos{
			Line:   int(params.Position.Line),
			Column: int(params.Position.Character),
			Byte:   pos.Byte,
		}, s.schemaMap)

		// completionItems := make([]protocol.CompletionItem, 0)

		// for _, comp := range completions {
		// 	completionItems = append(completionItems, protocol.CompletionItem{
		// 		Label:      comp,
		// 		Kind:       protocol.CompletionItemKindClass,
		// 		InsertText: fmt.Sprintf("%s {\n$0\n}", comp),
		// 	})
		// }

		reply(ctx, protocol.CompletionList{
			IsIncomplete: false,
			Items:        completions,
		}, nil)
	case protocol.MethodTextDocumentDidOpen:
		s.logger.Info("did open")

		params := protocol.DidOpenTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			reply(ctx, nil, err)
			return
		}

		s.parser.ParseHCL([]byte(params.TextDocument.Text), params.TextDocument.URI.Filename())

		s.logger.Info(fmt.Sprintf("%+v", params))

		reply(ctx, nil, nil)
	case protocol.MethodTextDocumentDidChange:
		s.logger.Info("did change")

		params := protocol.DidChangeTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			reply(ctx, nil, err)
			return
		}

		changesCount := len(params.ContentChanges)

		if changesCount > 0 {
			// s.logger.Info(fmt.Sprintf("%s", params.ContentChanges[changesCount-1].Text))

			delete(s.parser.Files(), params.TextDocument.URI.Filename())

			s.parser.ParseHCL([]byte(params.ContentChanges[changesCount-1].Text), params.TextDocument.URI.Filename())

			s.logger.Info(fmt.Sprintf("%+v", params))
		}

		reply(ctx, nil, nil)
	case protocol.MethodTextDocumentDidClose:
		s.logger.Info("did close")

		params := protocol.DidCloseTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			reply(ctx, nil, err)
			return
		}

		delete(s.parser.Files(), params.TextDocument.URI.Filename())

		s.logger.Info(fmt.Sprintf("%+v", s.parser.Files()))

		s.logger.Info(fmt.Sprintf("%+v", params))

		reply(ctx, nil, nil)
	case protocol.MethodShutdown:
		ctx.Done()
	}
}

func CollectHoverInfo(body hcl.Body, pos hcl.Pos, schemaMap map[string]*hcl.BodySchema) []string {
	var blockTypes []string

	dfs(body, schemaMap, &blockTypes, pos, schema.SchemaMapBetter["root"], &schema.RootBodySchema)

	return blockTypes
}

func dfs(body hcl.Body, schemaMap map[string]*hcl.BodySchema, arr *[]string, pos hcl.Pos, currSchema *hcl.BodySchema, nonHCLSchema *hclschema.BodySchema) {
	log.Printf("body: %+v", body)
	if currSchema == nil {
		return
	}

	// bodyRange := body.(*hclsyntax.Body).SrcRange
	bodyContent, _ := body.Content(currSchema)
	blocksByType := bodyContent.Blocks.ByType()
	log.Printf("block by type: %#v", blocksByType)

	for k, v := range blocksByType {
		for _, b := range v {
			blockRange := b.Body.(*hclsyntax.Body).SrcRange
			if !blockRange.ContainsPos(pos) {
				continue
			}
			log.Printf("block '%s': %+v", k, b)
			// log.Printf("block '%s' body: %+v", k, b.Body)

			*arr = append(*arr, k)
			// *arr = append(*arr, nonHCLSchema.Description.Value)
			if nonHCLSchema.Blocks[k] != nil && nonHCLSchema.Blocks[k].Body != nil {
				*arr = append(*arr, nonHCLSchema.Blocks[k].Description.Value)
				// log.Printf("%+v", nonHCLSchema.Blocks[k].Body)
				// *arr = append(*arr, nonHCLSchema.Blocks[k].Description.Value)
				dfs(b.Body, schemaMap, arr, pos, schemaMap[k], nonHCLSchema.Blocks[k].Body)
			}
		}
	}

	// attr, _ := body.JustAttributes()

	// log.Printf("attr: %#v", attr)
	// log.Printf("body: %+v", body)
	// log.Printf("body content: %+v", bodyContent)
	// log.Printf("body content blocks: %+v", bodyContent.Blocks)

	// for _, block := range bodyContent.Blocks {
	// }
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

		for k := range blocksByType {
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
					Label:            k,
					InsertText:       fmt.Sprintf("%s = \"$0\"", k),
					Kind:             protocol.CompletionItemKindVariable,
					Detail:           v.Description.Value,
					Documentation:    v.Description.Value,
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			case cty.List(cty.String):
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:            k,
					InsertText:       fmt.Sprintf("%s = [\"$0\"]", k),
					Kind:             protocol.CompletionItemKindVariable,
					Detail:           v.Description.Value,
					Documentation:    v.Description.Value,
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			default:
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:            k,
					InsertText:       fmt.Sprintf("%s = ", k),
					Kind:             protocol.CompletionItemKindVariable,
					Detail:           v.Description.Value,
					Documentation:    v.Description.Value,
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
