package lsp

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"go.lsp.dev/protocol"
)

func (s *Service) HandleInitialize(ctx context.Context, params *protocol.InitializedParams) (protocol.InitializeResult, error) {
	return protocol.InitializeResult{
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
	}, nil
}

func (s *Service) HandleTextDocumentHover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
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
		return nil, nil
	}

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.PlainText,
			Value: fmt.Sprintf("%s", x[len(x)-1]),
		},
	}, nil
}

func (s *Service) HandleTextDocumentCompletion(ctx context.Context, params *protocol.CompletionParams) (*protocol.CompletionList, error) {
	file := s.parser.Files()[params.TextDocument.URI.Filename()]

	if file == nil {
		return nil, errors.New("file is nil")
	}

	body := file.Body

	byteOffset := CalculateByteOffset(params.Position, s.parser.Files()[params.TextDocument.URI.Filename()].Bytes)

	pos := hcl.InitialPos
	pos.Byte = int(byteOffset)

	completions := CollectCompletions(body, hcl.Pos{
		Line:   int(params.Position.Line),
		Column: int(params.Position.Character),
		Byte:   pos.Byte,
	}, s.schemaMap)

	return &protocol.CompletionList{
		IsIncomplete: false,
		Items:        completions,
	}, nil
}

func (s *Service) HandleTextDocumentDidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	s.parser.ParseHCL([]byte(params.TextDocument.Text), params.TextDocument.URI.Filename())

	s.logger.Info(fmt.Sprintf("%+v", params))

	return nil
}

func (s *Service) HandleTextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	changesCount := len(params.ContentChanges)

	if changesCount > 0 {
		delete(s.parser.Files(), params.TextDocument.URI.Filename())

		s.parser.ParseHCL([]byte(params.ContentChanges[changesCount-1].Text), params.TextDocument.URI.Filename())

		s.logger.Info(fmt.Sprintf("%+v", params))
	}

	return nil
}

func (s *Service) HandleTextDocumentDidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	delete(s.parser.Files(), params.TextDocument.URI.Filename())

	s.logger.Info(fmt.Sprintf("%+v", s.parser.Files()))

	s.logger.Info(fmt.Sprintf("%+v", params))

	return nil
}
