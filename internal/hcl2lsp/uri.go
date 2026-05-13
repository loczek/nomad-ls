package hcl2lsp

import "go.lsp.dev/protocol"

func FileName(params protocol.TextDocumentIdentifier) string {
	return params.URI.Filename()
}

func FileNameVersioned(params protocol.VersionedTextDocumentIdentifier) string {
	return params.URI.Filename()
}

func FileNameItem(params protocol.TextDocumentItem) string {
	return params.URI.Filename()
}
