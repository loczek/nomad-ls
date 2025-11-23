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

	// if existing := p.files[filename]; existing != nil {
	// 	return existing, nil
	// }

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Byte: 0, Line: 1, Column: 1})
	p.files[filename] = file
	return file, diags
}

func (p *Parser) UpdateHCL(src []byte, filename string) (*hcl.File, hcl.Diagnostics) {
	p.mu.Lock()
	defer p.mu.Unlock()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Byte: 0, Line: 1, Column: 1})
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
