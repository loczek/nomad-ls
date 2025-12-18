package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var AffinitySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"attribute": {
			Description: lang.Markdown("Specifies the name or reference of the attribute to examine for the affinity. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars)."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.LiteralType{Type: cty.String},
		},
		// TODO: update docs
		"operator": {
			Description: lang.Markdown("Specifies the comparison operator. The ordering is compared lexically."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("="),
			},
			Constraint: schema.LiteralType{Type: cty.String},
		},
		"value": {
			Description: lang.Markdown("Specifies the value to compare the attribute against using the specified operation. This can be a literal value, another attribute, or any [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars). The `value` field is required."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.LiteralType{Type: cty.String},
			IsRequired: true,
		},
		"weight": {
			Description: lang.Markdown("Specifies a weight for the affinity. The weight is used during scoring and must be an integer between -100 to 100. Negative weights act as anti affinities, causing nodes that match them to be scored lower. Weights can be used when there is more than one affinity to express relative preference across them."),
			DefaultValue: schema.DefaultValue{
				Value: cty.NumberIntVal(50),
			},
			Constraint: schema.LiteralType{Type: cty.Number},
		},
	},
}
