package hcl2lsp

import "go.lsp.dev/protocol"

func FileName(params protocol.TextDocumentPositionParams) string {
	return params.TextDocument.URI.Filename()
}
