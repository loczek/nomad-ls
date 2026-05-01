package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ParameterizedSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		// TODO: shoud be required
		"meta_optional": {
			Description: lang.Markdown("Specifies the set of metadata keys that may be provided when dispatching against the job."),
			DefaultValue: schema.DefaultValue{
				Value: cty.ListVal([]cty.Value{cty.StringVal("")}),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.List(cty.String)},
				schema.AnyExpression{OfType: cty.List(cty.String)},
			},
			IsOptional: true,
		},
		// TODO: shoud be required
		"meta_required": {
			Description: lang.Markdown("Specifies the set of metadata keys that must be provided when dispatching against the job."),
			DefaultValue: schema.DefaultValue{
				Value: cty.ListVal([]cty.Value{cty.StringVal("")}),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.List(cty.String)},
				schema.AnyExpression{OfType: cty.List(cty.String)},
			},
			IsOptional: true,
		},
		"payload": {
			Description: lang.Markdown("Specifies the requirement of providing a payload when dispatching against the parameterized job. The maximum size of a `payload` is 16 KiB. The options for this field are:\n- `optional` - A payload is optional when dispatching against the job.\n- `required` - A payload must be provided when dispatching against the job.\n- `forbidden` - A payload is forbidden when dispatching against the job."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("optional"),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
	},
}
