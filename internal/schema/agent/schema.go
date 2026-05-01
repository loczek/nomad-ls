package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("The name of the Quota. Nomad uses name to connect the quota a [`Namespace`](https://developer.hashicorp.com/nomad/docs/other-specifications/namespace)."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"description": {
			Description: lang.Markdown("A human-readable description."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	// Blocks: map[string]*schema.BlockSchema{
	// 	"limit": {
	// 		Description: QuotaLimitSchema.Description,
	// 		Body:        QuotaLimitSchema,
	// 	},
	// },
}
