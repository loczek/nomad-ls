package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var SpreadSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"attribute": {
			Description: lang.PlainText("Specifies the name or reference of the attribute to use. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars)."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
		},
		"weight": {
			Description: lang.PlainText("Specifies a weight for the spread block. The weight is used during scoring and must be an integer between 0 to 100. Weights can be used when there is more than one spread or affinity block to express relative preference across them."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"target": {
			Description: lang.PlainText("Specifies one or more spread target percentages. If omitted, Nomad spreads evenly."),
			Body:        TargetSchema,
			Labels: []*schema.LabelSchema{
				{
					Name: "name",
				},
			},
		},
	},
}
