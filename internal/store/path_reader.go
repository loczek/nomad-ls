package store

import (
	"context"
	"errors"

	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/validator"
	"github.com/hashicorp/hcl/v2"
	funcs "github.com/loczek/nomad-ls/internal/function"
	"github.com/loczek/nomad-ls/internal/languages"
)

var _ decoder.PathReader = (*Store)(nil)

// PathContext implements [decoder.PathReader].
func (p *Store) PathContext(path lang.Path) (*decoder.PathContext, error) {
	langID := languages.LanguageID(path.LanguageID)
	langSchema := languages.ToSchema(langID)

	file, ok := p.files[path.Path]
	if !ok {
		return nil, errors.New("file not found")
	}

	return &decoder.PathContext{
		Schema:           &langSchema,
		ReferenceOrigins: file.RefOrigins,
		ReferenceTargets: file.RefTargets,
		Files: map[string]*hcl.File{
			path.Path: file.HCLFile,
		},
		Functions: funcs.Functions,
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

// Paths implements [decoder.PathReader].
func (p *Store) Paths(ctx context.Context) []lang.Path {
	var paths []lang.Path

	for path, val := range p.files {
		paths = append(paths, lang.Path{
			Path:       path,
			LanguageID: string(val.Language),
		})
	}

	return paths
}
