package lsp

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/schema"
	"go.lsp.dev/protocol"
)

const (
	LOKI_NOMAD_FILE_PATH = "./testdata/loki.nomad.hcl"
)

func TestByteCount(t *testing.T) {
	hclFile := LoadSampleFile(LOKI_NOMAD_FILE_PATH)

	actuallCount := len(hclFile.Bytes)

	pos := protocol.Position{Line: 100, Character: 0}

	predictedCount := CalculateByteOffset(pos, hclFile.Bytes)

	if actuallCount != int(predictedCount) {
		t.Errorf("expected: %d, recieved: %d", actuallCount, predictedCount)
	}
}

func TestSimpleParse(t *testing.T) {
	hclFile := LoadSampleFile(LOKI_NOMAD_FILE_PATH)

	bc, _ := hclFile.Body.Content(schema.SchemaMapBetter["root"])
	// t.Logf("%+v", bc.Blocks.ByType())
	// t.Logf("ed: %+v", bc)

	x := hclFile.Body.(*hclsyntax.Body)
	// t.Logf("body syntax cast: %#v", x.SrcRange.)

	t.Logf("body range: %d:%d to %d:%d", x.SrcRange.Start.Line, x.SrcRange.Start.Column, x.SrcRange.End.Line, x.SrcRange.End.Column)

	block := bc.Blocks[2]
	bc, _ = block.Body.Content(schema.SchemaMapBetter["job"])
	// t.Logf("%+v", bc.Blocks.ByType())
	// t.Logf("ed: %#v", bc)
	t.Logf("job block: %#v", block.DefRange)

	block = bc.Blocks[0]
	bc, _ = block.Body.Content(schema.SchemaMapBetter["group"])
	// t.Logf("%+v", bc.Blocks.ByType())
	// t.Logf("ed: %#v", bc)
	t.Logf("group block: %#v", block.DefRange)

	// t.Logf("%+v", attr)
	// t.Logf("blocks: %v", blocks)
}

func TestServiceBlockHoverInformation(t *testing.T) {
	hclFile := LoadSampleFile(LOKI_NOMAD_FILE_PATH)

	pos := protocol.Position{Line: 28, Character: 5}

	predictedCount := CalculateByteOffset(pos, hclFile.Bytes)

	blocks := CollectHoverInfo(hclFile.Body, hcl.Pos{
		Line:   int(pos.Line),
		Column: int(pos.Character),
		Byte:   int(predictedCount),
	}, schema.SchemaMapBetter)

	t.Logf("blocks: %v", blocks)

	x := blocks[len(blocks)-1]

	if !strings.HasPrefix(x, "Specifies integrations with Noma") {
		t.Errorf("wrong hover information %s", x)
	}
}

func TestBlockCompletion(t *testing.T) {
	hclFile := LoadSampleFile(LOKI_NOMAD_FILE_PATH)

	pos := protocol.Position{Line: 14, Character: 0}

	predictedCount := CalculateByteOffset(pos, hclFile.Bytes)

	blocks := CollectCompletions(hclFile.Body, hcl.Pos{
		Line:   int(pos.Line),
		Column: int(pos.Character),
		Byte:   int(predictedCount),
	}, schema.SchemaMapBetter)

	t.Logf("blocks: %v", blocks)

	if len(blocks) == 0 {
		t.Errorf("blocks empty")
	}
}

func LoadSampleFile(path string) *hcl.File {
	parser := hclparse.NewParser()

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parser.ParseHCL(file, "loki")

	hclFile := parser.Files()["loki"]

	return hclFile
}
