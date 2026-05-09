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
	"github.com/loczek/nomad-ls/internal/store"
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
				Change:    protocol.TextDocumentSyncKindFull,
				OpenClose: true,
			},
			DocumentFormattingProvider: &protocol.DocumentFormattingOptions{},
		},
	}, nil
}

func (s *Service) HandleTextDocumentHover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	fileName := hcl2lsp.FileName(params.TextDocument)
	file, err := s.store.GetFile(fileName)
	if err != nil {
		return nil, err
	}

	pos := hcl2lsp.Position(params.Position, file.HCLFile.Bytes)

	dec := decoder.NewDecoder(&s.store)
	langPath := lang.Path{
		Path:       fileName,
		LanguageID: string(file.Language),
	}

	pathDec, err := dec.Path(langPath)
	if err != nil {
		panic(err)
	}

	hoverData, err := pathDec.HoverAtPos(ctx, fileName, pos)
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
	fileName := hcl2lsp.FileName(params.TextDocument)
	file, err := s.store.GetFile(fileName)
	if err != nil {
		return nil, err
	}

	pos := hcl2lsp.Position(params.Position, file.HCLFile.Bytes)

	dec := decoder.NewDecoder(&s.store)
	langPath := lang.Path{
		Path:       fileName,
		LanguageID: string(file.Language),
	}

	pathDec, err := dec.Path(langPath)
	if err != nil {
		return nil, err
	}

	cands, err := pathDec.CompletionAtPos(ctx, fileName, pos)
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
	fileName := hcl2lsp.FileNameItem(params.TextDocument)
	langID, err := languages.NewFromString(string(params.TextDocument.LanguageID))
	if err != nil {
		return nil, err
	}

	newFile := store.NewDocument(langID)
	_, diags := newFile.ParseHCL([]byte(params.TextDocument.Text), fileName)
	file := s.store.AddFile(fileName, newFile)

	dec := decoder.NewDecoder(&s.store)
	langPath := lang.Path{
		Path:       fileName,
		LanguageID: langID.String(),
	}

	dec.SetContext(decoder.NewDecoderContext())

	pathDec, err := dec.Path(langPath)
	if err != nil {
		return nil, err
	}

	file.UpdateReferences(pathDec, fileName)

	diagsPath, err := pathDec.ValidateFile(ctx, fileName)
	if err != nil {
		return nil, err
	}

	diags = diags.Extend(diagsPath)

	return &diags, nil
}

func (s *Service) HandleTextDocumentDidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (*hcl.Diagnostics, error) {
	changesCount := len(params.ContentChanges)

	if changesCount == 0 {
		return nil, nil
	}

	fileName := hcl2lsp.FileNameVersioned(params.TextDocument)
	file, err := s.store.GetFile(fileName)
	if err != nil {
		return nil, err
	}

	_, diags := file.UpdateHCL([]byte(params.ContentChanges[changesCount-1].Text), fileName)

	dec := decoder.NewDecoder(&s.store)
	langPath := lang.Path{
		Path:       fileName,
		LanguageID: string(file.Language),
	}

	dec.SetContext(decoder.NewDecoderContext())

	pathDec, err := dec.Path(langPath)
	if err != nil {
		return nil, err
	}

	file.UpdateReferences(pathDec, fileName)

	pathContext, err := s.store.PathContext(langPath)
	if err != nil {
		return nil, err
	}

	diagMap := validation.UnreferencedOrigins(ctx, pathContext)
	originsDiags := hcl.Diagnostics{}
	for _, v := range diagMap {
		originsDiags = originsDiags.Extend(v)
	}

	diags = diags.Extend(originsDiags)

	schemaDiags, err := pathDec.ValidateFile(ctx, fileName)
	if err != nil {
		return nil, err
	}

	diags = diags.Extend(schemaDiags)

	s.logger.Info(fmt.Sprintf("diags: %+v", diags))

	return &diags, nil
}

func (s *Service) HandleTextDocumentDidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	s.store.RemoveFile(hcl2lsp.FileName(params.TextDocument))

	return nil
}

func (s *Service) HandleTextDocumentFormatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	fileName := hcl2lsp.FileName(params.TextDocument)
	file, err := s.store.GetFile(fileName)
	if err != nil {
		return nil, err
	}

	outBytes := hclwrite.Format(file.HCLFile.Bytes)

	var edits []protocol.TextEdit

	if !bytes.Equal(file.HCLFile.Bytes, outBytes) {
		startPos := protocol.Position{Line: 0, Character: 0}
		endPos := getLastPostionFromBytes(file.HCLFile.Bytes)

		edits = append(edits, protocol.TextEdit{
			Range: protocol.Range{
				Start: startPos,
				End:   endPos,
			},
			NewText: string(outBytes),
		})
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
