package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/scope"
	"github.com/zclconf/go-cty/cty"
)

var RootBodySchema = schema.BodySchema{
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
		"variables": {
			Description: VariablesSchema.Description,
			Body:        VariablesSchema,
		},
		"locals": {
			Description: lang.Markdown("Local values assigning names to expressions, so you can use these multiple times without repetition\n" +
				"e.g. `service_name = \"forum\"`"),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Address: &schema.AttributeAddrSchema{
						Steps: []schema.AddrStep{
							schema.StaticStep{Name: "local"},
							schema.AttrNameStep{},
						},
						FriendlyName: "local",
						ScopeId:      scope.LocalScope,
						AsReference:  true,
						AsExprType:   true,
					},
					Constraint: schema.AnyExpression{OfType: cty.DynamicPseudoType},
				},
			},
		},
		"job": {
			Description: JobSchema.Description,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			MinItems: 1,
			MaxItems: 1,
			Body:     JobSchema,
		},
	},
}
