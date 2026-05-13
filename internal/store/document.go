package store

import (
	"sync"

	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/reference"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/languages"
	"github.com/loczek/nomad-ls/internal/references"
)

// Document represents a open file in memory
type Document struct {
	HCLFile *hcl.File

	RefTargets reference.Targets
	RefOrigins reference.Origins

	Language languages.LanguageID
	Version  int32

	mu sync.Mutex
}

func NewDocument(language languages.LanguageID) *Document {
	return &Document{
		HCLFile:    &hcl.File{},
		RefTargets: make(reference.Targets, 0),
		RefOrigins: make(reference.Origins, 0),
		Language:   language,
		mu:         sync.Mutex{},
	}
}

func (f *Document) ParseHCL(src []byte, filename string) (*hcl.File, hcl.Diagnostics) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	f.HCLFile = file

	return file, diags
}

func (f *Document) UpdateHCL(src []byte, filename string) (*hcl.File, hcl.Diagnostics) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	f.HCLFile = file

	return file, diags
}

func (f *Document) UpdateReferences(pathDecoder *decoder.PathDecoder, fileName string) error {
	targets, err := pathDecoder.CollectReferenceTargets()
	if err != nil {
		return err
	}

	f.RefTargets = append(targets, references.CommonBuiltinReferences()...)

	origins, err := pathDecoder.CollectReferenceOrigins()
	if err != nil {
		return err
	}

	f.RefOrigins = origins

	return nil
}
