package variable

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/scope"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"variable": {
			Address: &schema.BlockAddrSchema{
				Steps: []schema.AddrStep{
					schema.StaticStep{Name: "var"},
					schema.LabelStep{Index: 0},
				},
				FriendlyName: "variable",
				ScopeId:      scope.VariableScope,
				AsReference:  true,
				AsTypeOf: &schema.BlockAsTypeOf{
					AttributeExpr: "type",
				},
			},
			Description: VariableSchema.Description,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: VariableSchema,
		},
	},
}

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
			Constraint:  schema.TypeDeclaration{},
			IsOptional:  true,
		},
		"description": {
			Description: lang.Markdown("variable description"),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
	},
}
