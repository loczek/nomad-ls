package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"unicode/utf8"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/schema"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

type Service struct {
	con       jsonrpc2.Conn
	parser    hclparse.Parser
	schemaMap map[string]*hcl.BodySchema
	// files  map[string]string
	logger slog.Logger
}

// type File struct {
// 	URI     string
// 	content string
// }

func New(con jsonrpc2.Conn, logger slog.Logger) Service {
	return Service{
		con:       con,
		parser:    *hclparse.NewParser(),
		schemaMap: schema.SchemaMap,
		// files:  map[string]string{},
		logger: logger,
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

		// hclsyntax.pars

		pos := hcl.InitialPos
		pos.Byte = int(byteOffset)

		// x := collectBlockTypes(body, s.schemaMap)
		x := CollectBlockTypes(body, hcl.Pos{
			Line:   int(params.Position.Line),
			Column: int(params.Position.Character),
			Byte:   pos.Byte,
		}, s.schemaMap)

		s.logger.Info(fmt.Sprintf("arr: %v", x))

		// y := int32(x.Position.Line);

		// x := hcl.Pos{
		// 	Line:   0,
		// 	Column: 0,
		// 	Byte:   0,
		// }

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

		reply(ctx, protocol.CompletionList{
			IsIncomplete: false,
			Items: []protocol.CompletionItem{
				{
					Label: "test_one",
				},
				{
					Label: "test_two",
				},
				{
					Label: "test_three",
				},
			},
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

func CollectBlockTypes(body hcl.Body, pos hcl.Pos, schemaMap map[string]*hcl.BodySchema) []string {
	var blockTypes []string

	dfs(body, schemaMap, &blockTypes, pos, schema.JobConfigSchema)

	return blockTypes
}

func dfs(body hcl.Body, schemaMap map[string]*hcl.BodySchema, arr *[]string, pos hcl.Pos, currSchema *hcl.BodySchema) {
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
			*arr = append(*arr, k)
			// log.Printf("block '%s': %+v", k, b)
			// log.Printf("block '%s' body: %+v", k, b.Body)

			dfs(b.Body, schemaMap, arr, pos, schemaMap[k])
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

	log.Printf("%v", runes)

	var runeIndex uint
	var line uint
	var bytesCount uint

	for line < uint(pos.Line) {
		fmt.Println("line")
		if runes[runeIndex] == '\n' {
			line += 1
		}
		fmt.Println(runeIndex)
		bytesCount += uint(utf8.RuneLen(runes[runeIndex]))
		runeIndex += 1
	}

	var j uint

	for j < uint(pos.Character) {
		fmt.Println("char")
		bytesCount += uint(utf8.RuneLen(runes[runeIndex]))
		runeIndex += 1
		j += 1
	}

	return bytesCount
}
