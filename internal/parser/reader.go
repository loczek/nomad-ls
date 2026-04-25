package parser

import (
	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/loczek/nomad-ls/internal/languages"
	"github.com/loczek/nomad-ls/internal/references"
)

func (p *Parser) UpdateReferences(fileName string) error {
	dec := decoder.NewDecoder(p)

	pd, err := dec.Path(lang.Path{
		Path:       fileName,
		LanguageID: languages.NomadJob.String(),
	})
	if err != nil {
		return err
	}

	targets, err := pd.CollectReferenceTargets()
	if err != nil {
		return err
	}

	p.RefTargets = append(targets, references.CommonBuiltinReferences()...)

	origins, err := pd.CollectReferenceOrigins()
	if err != nil {
		return err
	}

	p.RefOrigins = origins

	return nil
}
