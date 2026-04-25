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
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"percent": {
			Description: lang.PlainText("Specifies the percentage associated with the target value."),
			DefaultValue: schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
	},
}
