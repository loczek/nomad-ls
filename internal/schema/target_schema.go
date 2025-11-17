package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var TargetSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"value": {
			Description: lang.PlainText("Specifies a target value of the attribute from a `spread` block."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
		},
		"percent": {
			Description: lang.PlainText("Specifies the percentage associated with the target value."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
		},
	},
}
