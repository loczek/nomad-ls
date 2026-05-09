package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"

	"github.com/loczek/nomad-ls/internal/store"
)

type Service struct {
	con    jsonrpc2.Conn
	store  store.Store
	logger slog.Logger
}

func New(con jsonrpc2.Conn, logger slog.Logger) Service {
	return Service{
		con:    con,
		store:  store.NewStore(),
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

		diags, err := s.HandleTextDocumentDidOpen(ctx, &params)

		if diags != nil {
			s.con.Notify(context.Background(), protocol.MethodTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
				URI:         params.TextDocument.URI,
				Version:     uint32(params.TextDocument.Version),
				Diagnostics: *diags,
			})
		}

		return nil, err
	case protocol.MethodTextDocumentDidChange:
		params := protocol.DidChangeTextDocumentParams{}
		err := json.Unmarshal(req.Params(), &params)
		if err != nil {
			return nil, err
		}

		diags, err := s.HandleTextDocumentDidChange(ctx, &params)

		if diags != nil {
			s.con.Notify(context.Background(), protocol.MethodTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
				URI:         params.TextDocument.URI,
				Version:     uint32(params.TextDocument.Version),
				Diagnostics: *diags,
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
	default:
		return nil, fmt.Errorf("Received unimplementhed method: %s", req.Method())
	}

}
