package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"unicode/utf8"

	"github.com/hashicorp/hcl/v2"

	"github.com/loczek/nomad-ls/internal/parser"

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
		con:    con,
		parser: *parser.NewParser(),
		logger: logger,
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
		diag, err := s.HandleTextDocumentDidOpen(ctx, &params)

		if diag != nil {
			protocolDiagnostics := hcl2lsp.Diagnostics(*diag)

			s.con.Notify(context.Background(), protocol.MethodTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
				URI:         params.TextDocument.URI,
				Version:     uint32(params.TextDocument.Version),
				Diagnostics: protocolDiagnostics,
			})
		}

		return nil, err
	case protocol.MethodTextDocumentDidChange:
		params := protocol.DidChangeTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		diag, err := s.HandleTextDocumentDidChange(ctx, &params)

		if diag != nil {
			protocolDiagnostics := hcl2lsp.Diagnostics(*diag)

			s.con.Notify(context.Background(), protocol.MethodTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
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
	case protocol.MethodTextDocumentFormatting:
		params := protocol.DocumentFormattingParams{}

		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		return s.HandleTextDocumentFormatting(ctx, &params)
	case protocol.MethodShutdown:
		ctx.Done()
		return nil, nil
	}

	return nil, nil
}

func asBlock(name string, depth int) string {
	return fmt.Sprintf("%s \"${1:name}\" {\n%s$0\n}", name, strings.Repeat("\t", depth))
}

func asAnonymousBlock(name string, depth int) string {
	return fmt.Sprintf("%s {\n%s$0\n}", name, strings.Repeat("\t", depth))
}
