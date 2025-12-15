package lsp

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/loczek/nomad-ls/internal/schema"
	"go.lsp.dev/protocol"
)

const (
	LOKI_NOMAD_FILE_PATH              = "./testdata/loki.nomad.hcl"
	GENERIC_NOMAD_FILE_PATH           = "./testdata/generic.nomad.hcl"
	INVALID_ATTRIBUTE_NOMAD_FILE_PATH = "./testdata/invalid_attribute.nomad.hcl"
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

func TestServiceBlockHoverInformation(t *testing.T) {
	tests := []struct {
		name           string
		filePath       string
		pos            protocol.Position
		expectedPrefix string
	}{
		{
			name:           "loki nomad file",
			filePath:       LOKI_NOMAD_FILE_PATH,
			pos:            protocol.Position{Line: 28, Character: 5},
			expectedPrefix: "Specifies integrations with Noma",
		},
		{
			name:           "generic nomad file",
			filePath:       GENERIC_NOMAD_FILE_PATH,
			pos:            protocol.Position{Line: 0, Character: 5},
			expectedPrefix: "A less precise block for decla",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hclFile := LoadSampleFile(tt.filePath)

			predictedCount := CalculateByteOffset(tt.pos, hclFile.Bytes)

			blocks := CollectHoverInfo(hclFile.Body, hcl.Pos{
				Line:   int(tt.pos.Line),
				Column: int(tt.pos.Character),
				Byte:   int(predictedCount),
			}, schema.SchemaMapBetter)

			t.Logf("blocks: %v", blocks)

			x := blocks[len(blocks)-1]

			if !strings.HasPrefix(x, tt.expectedPrefix) {
				t.Errorf("wrong hover information '%s'", x)
			}
		})
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

func TestMetaBlockAllowsAnyAttribute(t *testing.T) {
	hclFile := LoadSampleFile(GENERIC_NOMAD_FILE_PATH)

	diags := CollectDiagnostics(hclFile.Body, schema.SchemaMapBetter)

	// Filter for errors only (ignore warnings)
	var errors hcl.Diagnostics
	for _, d := range *diags {
		if d.Severity == hcl.DiagError {
			errors = append(errors, d)
		}
	}

	if len(errors) > 0 {
		for _, d := range errors {
			t.Errorf("unexpected diagnostic: %s at %v", d.Summary, d.Subject)
		}
	}
}

func TestInvalidAttributeGeneratesDiagnostic(t *testing.T) {
	hclFile := LoadSampleFile(INVALID_ATTRIBUTE_NOMAD_FILE_PATH)

	diags := CollectDiagnostics(hclFile.Body, schema.SchemaMapBetter)

	// Filter for errors only
	var errors hcl.Diagnostics
	for _, d := range *diags {
		if d.Severity == hcl.DiagError {
			errors = append(errors, d)
		}
	}

	if len(errors) == 0 {
		t.Error("expected diagnostic error for invalid attribute, but got none")
	}

	// Verify the error is about an unsupported attribute
	found := false
	for _, d := range errors {
		if strings.Contains(d.Summary, "Unsupported") || strings.Contains(d.Summary, "unsupported") {
			found = true
			t.Logf("correctly detected invalid attribute: %s", d.Summary)
			break
		}
	}

	if !found {
		t.Errorf("expected 'unsupported attribute' error, got: %v", errors)
	}
}

func LoadSampleFile(path string) *hcl.File {
	parser := hclparse.NewParser()

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parser.ParseHCL(file, "nomad-job")

	hclFile := parser.Files()["nomad-job"]

	return hclFile
}
