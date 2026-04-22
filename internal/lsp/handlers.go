package lsp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"go.lsp.dev/protocol"

	"github.com/loczek/nomad-ls/internal/hcl2lsp"
	"github.com/loczek/nomad-ls/internal/languages"
	"github.com/loczek/nomad-ls/internal/validation"
)

func (s *Service) HandleInitialize(ctx context.Context, params *protocol.InitializedParams) (*protocol.InitializeResult, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("could not read build info")
	}

	return &protocol.InitializeResult{
		ServerInfo: &protocol.ServerInfo{
			Name:    "nomad-ls",
			Version: strings.TrimPrefix(info.Main.Version, "v"),
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

	pos := hcl2lsp.Position(params.Position, file.Bytes)

	dec := decoder.NewDecoder(&s.parser)
	langPath := lang.Path{
		Path:       params.TextDocument.URI.Filename(),
		LanguageID: languages.NomadJob.String(),
	}

	pathDec, err := dec.Path(langPath)
	if err != nil {
		panic(err)
	}

	hoverData, err := pathDec.HoverAtPos(ctx, params.TextDocument.URI.Filename(), pos)
	if err != nil {
		return nil, err
	}

	if hoverData == nil {
		return nil, nil
	}

	var hover = hcl2lsp.Hover(hoverData)

	return &hover, nil
}

func (s *Service) HandleTextDocumentCompletion(ctx context.Context, params *protocol.CompletionParams) (*protocol.CompletionList, error) {
	file := s.parser.Files()[params.TextDocument.URI.Filename()]

	if file == nil {
		return nil, errors.New("file is not found")
	}

	pos := hcl2lsp.Position(params.Position, s.parser.Files()[params.TextDocument.URI.Filename()].Bytes)

	dec := decoder.NewDecoder(&s.parser)
	langPath := lang.Path{
		Path:       params.TextDocument.URI.Filename(),
		LanguageID: languages.NomadJob.String(),
	}

	pathDec, err := dec.Path(langPath)
	if err != nil {
		return nil, err
	}

	cands, err := pathDec.CompletionAtPos(ctx, params.TextDocument.URI.Filename(), pos)
	if err != nil {
		return nil, err
	}

	completions := hcl2lsp.Completions(cands)

	return &protocol.CompletionList{
		IsIncomplete: cands.IsComplete,
		Items:        completions,
	}, nil
}

func (s *Service) HandleTextDocumentDidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (*hcl.Diagnostics, error) {
	_, diags := s.parser.ParseHCL([]byte(params.TextDocument.Text), params.TextDocument.URI.Filename())

	dec := decoder.NewDecoder(&s.parser)
	langPath := lang.Path{
		Path:       params.TextDocument.URI.Filename(),
		LanguageID: languages.NomadJob.String(),
	}

	dec.SetContext(decoder.NewDecoderContext())

	pathDec, err := dec.Path(langPath)
	if err != nil {
		return nil, err
	}

	diags, err = pathDec.ValidateFile(ctx, params.TextDocument.URI.Filename())
	if err != nil {
		return nil, err
	}

	diags = diags.Extend(diags)

	return &diags, nil
}

func (s *Service) HandleTextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (*hcl.Diagnostics, error) {
	changesCount := len(params.ContentChanges)

	if changesCount > 0 {
		_, diags := s.parser.UpdateHCL([]byte(params.ContentChanges[changesCount-1].Text), params.TextDocument.URI.Filename())

		dec := decoder.NewDecoder(&s.parser)
		langPath := lang.Path{
			Path:       params.TextDocument.URI.Filename(),
			LanguageID: languages.NomadJob.String(),
		}

		dec.SetContext(decoder.NewDecoderContext())

		pathDec, err := dec.Path(langPath)
		if err != nil {
			return nil, err
		}

		pathContext, err := s.parser.PathContext(langPath)
		if err != nil {
			return nil, err
		}

		diagMap := validation.UnreferencedOrigins(ctx, pathContext)
		originsDiags := hcl.Diagnostics{}
		for _, v := range diagMap {
			originsDiags = originsDiags.Extend(v)
		}

		diags = diags.Extend(originsDiags)

		schemaDiags, err := pathDec.ValidateFile(ctx, params.TextDocument.URI.Filename())
		if err != nil {
			return nil, err
		}

		diags = diags.Extend(schemaDiags)

		s.logger.Info(fmt.Sprintf("diags: %+v", diags))

		return &diags, nil
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
