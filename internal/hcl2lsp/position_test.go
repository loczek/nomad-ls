package hcl2lsp

import (
	"os"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/loczek/nomad-ls/internal/languages"
	"github.com/loczek/nomad-ls/internal/store"
	"go.lsp.dev/protocol"
)

const (
	GENERIC_NOMAD_FILE_PATH = "./testdata/generic.nomad.hcl"
)

func TestConvertProtocolPosition(t *testing.T) {
	hclFile := LoadSampleFile(GENERIC_NOMAD_FILE_PATH)

	actuallPos := hcl.Pos{Line: 14, Column: 3, Byte: 182}
	predictedPos := Position(protocol.Position{Line: 13, Character: 2}, hclFile.Bytes)

	if actuallPos != predictedPos {
		t.Errorf("expected: %v, recieved: %v", actuallPos, predictedPos)
	}
}

func LoadSampleFile(path string) *hcl.File {
	parser := parser.NewParser()

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	doc := store.NewDocument(languages.NomadJob)

	doc.ParseHCL(file, "name")

	return doc.HCLFile
}
