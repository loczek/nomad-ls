package parser

import (
	"sync"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Parser struct {
	files map[string]*hcl.File
	mu    sync.Mutex
}

func NewParser() *Parser {
	return &Parser{
		files: map[string]*hcl.File{},
		mu:    sync.Mutex{},
	}
}

func (p *Parser) ParseHCL(src []byte, filename string) (*hcl.File, hcl.Diagnostics) {
	p.mu.Lock()
	defer p.mu.Unlock()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	p.files[filename] = file
	return file, diags
}

func (p *Parser) UpdateHCL(src []byte, filename string) (*hcl.File, hcl.Diagnostics) {
	p.mu.Lock()
	defer p.mu.Unlock()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	p.files[filename] = file
	return file, diags
}

func (p *Parser) RemoveHCL(filename string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.files, filename)
}

func (p *Parser) Files() map[string]*hcl.File {
	return p.files
}
