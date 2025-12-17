package schema

import (
	"github.com/hashicorp/hcl-lang/schema"
)

var RootBodySchema = schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"variable": {
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
			Body: JobSchema,
		},
	},
}
