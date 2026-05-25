package job

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var WeightsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"passing": {
			Description:  lang.Markdown("The weight of services in passing state."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(1)},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
		"warning": {
			Description:  lang.Markdown("The weight of services in warning state."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(1)},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
	},
}
