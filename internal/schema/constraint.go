package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ConstraintSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"attribute": {
			Description: lang.Markdown("Specifies the name or reference of the attribute to examine for the constraint. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars)."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsRequired: true,
		},
		// TODO: update docs
		"operator": {
			Description: lang.Markdown("Specifies the comparison operator. If the operator is one of `>, >=, <, <=`, the ordering is compared numerically if the operands are both integers or both floats, and lexically otherwise"),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("="),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsRequired: true,
		},
		"value": {
			Description: lang.Markdown("Specifies the value to compare the attribute against using the specified operation. This can be a literal value, another attribute, or any [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars). The value field is required except for when using the `is_set`, `is_not_set`, `distinct_hosts`, or `distinct_property` operators."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsRequired: true,
		},
	},
}
