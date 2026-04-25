package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var DispatchPayloadSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"file": {
			Description: lang.Markdown("Specifies the file name to write the content of dispatch payload to. The file is written relative to the [task's local directory](https://developer.hashicorp.com/nomad/docs/reference/runtime-environment-settings#local)."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
	},
}
