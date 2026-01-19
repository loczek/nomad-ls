// Package for converting from hcl to lsp types
package hcl2lsp

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl/v2"
	"go.lsp.dev/protocol"
)

func Completions(cands lang.Candidates) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	for _, v := range cands.List {
		completions = append(completions, protocol.CompletionItem{
			Label: v.Label,
			Kind:  protocol.CompletionItemKind(v.Kind),
			TextEdit: &protocol.TextEdit{
				NewText: v.TextEdit.Snippet,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      uint32(v.TextEdit.Range.Start.Line),
						Character: uint32(v.TextEdit.Range.Start.Column),
					},
					End: protocol.Position{
						Line:      uint32(v.TextEdit.Range.End.Line),
						Character: uint32(v.TextEdit.Range.End.Column),
					},
				},
			},
			Detail: v.Detail,
			Documentation: protocol.MarkupContent{
				Kind:  "markdown",
				Value: v.Description.Value,
			},
			InsertTextFormat: protocol.InsertTextFormatSnippet,
			SortText:         v.SortText,
		})
	}

	return completions
}

func Diagnostics(diag hcl.Diagnostics) []protocol.Diagnostic {
	protocolDiagnostics := []protocol.Diagnostic{}

	for _, v := range diag {
		newDiag := protocol.Diagnostic{
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
			Message: v.Summary,
		}

		if newDiag.Message == "" {
			newDiag.Message = v.Detail
		}

		protocolDiagnostics = append(protocolDiagnostics, newDiag)
	}

	return protocolDiagnostics
}
