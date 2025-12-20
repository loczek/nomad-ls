package lsp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"go.lsp.dev/protocol"
)

func (s *Service) HandleInitialize(ctx context.Context, params *protocol.InitializedParams) (*protocol.InitializeResult, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("could not read build info")
	}

	return &protocol.InitializeResult{
		ServerInfo: &protocol.ServerInfo{
			Name:    "nomad-ls",
			Version: info.Main.Version,
		},
		Capabilities: protocol.ServerCapabilities{
			CompletionProvider: &protocol.CompletionOptions{},
			HoverProvider:      &protocol.HoverOptions{},
			TextDocumentSync: &protocol.TextDocumentSyncOptions{
				Change: protocol.TextDocumentSyncKindFull,
			},
			DocumentFormattingProvider: &protocol.DocumentFormattingOptions{},
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
	})

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
	})

	return &protocol.CompletionList{
		IsIncomplete: false,
		Items:        completions,
	}, nil
}

func (s *Service) HandleTextDocumentDidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (*hcl.Diagnostics, error) {
	file, diags := s.parser.ParseHCL([]byte(params.TextDocument.Text), params.TextDocument.URI.Filename())

	schemaDiags := CollectDiagnostics(file.Body)

	allDiags := diags.Extend(*schemaDiags)

	s.logger.Info(fmt.Sprintf("%+v", params))

	return &allDiags, nil
}

func (s *Service) HandleTextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (*hcl.Diagnostics, error) {
	changesCount := len(params.ContentChanges)

	if changesCount > 0 {
		_, diags := s.parser.UpdateHCL([]byte(params.ContentChanges[changesCount-1].Text), params.TextDocument.URI.Filename())

		s.logger.Info(fmt.Sprintf("text: %+v", params))

		file := s.parser.Files()[params.TextDocument.URI.Filename()]

		body := file.Body

		schemaDiags := CollectDiagnostics(body)

		allDiags := diags.Extend(*schemaDiags)

		s.logger.Info(fmt.Sprintf("diags: %+v", allDiags))

		return &allDiags, nil
	}

	return nil, nil
}

func (s *Service) HandleTextDocumentDidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	s.parser.RemoveHCL(params.TextDocument.URI.Filename())

	s.logger.Info(fmt.Sprintf("%+v", s.parser.Files()))

	s.logger.Info(fmt.Sprintf("%+v", params))

	return nil
}

func (s *Service) HandleTextDocumentFormatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	filename := params.TextDocument.URI.Filename()

	var edits []protocol.TextEdit

	if file, ok := s.parser.Files()[filename]; ok {
		outBytes := hclwrite.Format(file.Bytes)

		if !bytes.Equal(file.Bytes, outBytes) {
			startPos := protocol.Position{Line: 0, Character: 0}
			endPos := getLastPostionFromBytes(file.Bytes)

			edits = append(edits, protocol.TextEdit{
				Range: protocol.Range{
					Start: startPos,
					End:   endPos,
				},
				NewText: string(outBytes),
			})
		}
	}

	return edits, nil
}

func getLastPostionFromBytes(src []byte) protocol.Position {
	runes := []rune(string(src))

	var runeIndex uint
	var line uint
	var character uint

	for runeIndex < uint(len(runes)) {
		character += 1
		if runes[runeIndex] == '\n' {
			line += 1
			character = 0
		}
		runeIndex += 1
	}

	return protocol.Position{
		Line:      uint32(line),
		Character: uint32(character),
	}
}
