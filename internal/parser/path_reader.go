package parser

import (
	"context"

	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/validator"
	funcs "github.com/loczek/nomad-ls/internal/function"
	nomadSchema "github.com/loczek/nomad-ls/internal/schema"
)

var _ decoder.PathReader = (*Parser)(nil)

func (p *Parser) PathContext(path lang.Path) (*decoder.PathContext, error) {
	return &decoder.PathContext{
		Schema:           &nomadSchema.RootBodySchema,
		ReferenceOrigins: p.RefOrigins,
		ReferenceTargets: p.RefTargets,
		Files:            p.files,
		Functions:        funcs.Functions,
		Validators: []validator.Validator{
			validator.BlockLabelsLength{},
			validator.DeprecatedAttribute{},
			validator.DeprecatedBlock{},
			validator.MaxBlocks{},
			validator.MinBlocks{},
			validator.MissingRequiredAttribute{},
			validator.UnexpectedAttribute{},
			validator.UnexpectedBlock{},
		},
	}, nil
}

func (p *Parser) Paths(ctx context.Context) []lang.Path {
	var paths []lang.Path

	for path := range p.files {
		paths = append(paths, lang.Path{
			Path:       path,
			LanguageID: "nomad",
		})
	}

	return paths
}
