package schema

import (
	"github.com/hashicorp/hcl-lang/schema"
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
