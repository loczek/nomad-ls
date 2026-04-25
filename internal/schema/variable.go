package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// TODO: update docs
var VariableSchema = &schema.BodySchema{
	Description: lang.Markdown("Input variables serve as parameters for a Nomad job, allowing aspects of the job to be customized without altering the job's own source code.\nWhen you declare variables in the same file as the job specification, you can set their values using CLI options and environment variables."),
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description: lang.Markdown("The type of HCL variable: `string`, `number`, `bool`."),
			Constraint:  schema.TypeDeclaration{},
			IsOptional:  true,
		},
		"default": {
			Description: lang.Markdown("The default value used when no value for this variable is provided."),
			Constraint:  schema.LiteralType{Type: cty.DynamicPseudoType},
			IsOptional:  true,
		},
		"description": {
			Description: lang.Markdown("variable description"),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
