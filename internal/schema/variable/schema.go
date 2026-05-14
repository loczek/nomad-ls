package variable

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"path": {
			Description: lang.Markdown("The path to the variable being defined. If empty it must be specified on the command line."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"namespace": {
			Description: lang.Markdown("The namespace of the variable. May be overridden by the `-namespace` command line flag or `NOMAD_NAMESPACE` environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"items": {
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
				},
			},
		},
	},
}
