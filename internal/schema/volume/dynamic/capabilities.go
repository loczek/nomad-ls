package dynamic

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var CapabilitySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"access_mode": {
			Description: lang.Markdown("Specifies the name or reference of the attribute to examine for the affinity. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars)."),
			Constraint: schema.OneOf{
				schema.LiteralValue{
					Value:       cty.StringVal("single-node-writer"),
					Description: lang.PlainText("Jobs can only request the volume with read/write access."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("single-node-reader-only"),
					Description: lang.PlainText("Jobs can only request the volume with read-only access."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("single-node-single-writer"),
					Description: lang.PlainText("Jobs can request either read/write or read-only access, but the scheduler only allows one allocation to have read/write access."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("single-node-multi-writer"),
					Description: lang.PlainText("Jobs can request either read/write or read-only access, and the scheduler allows multiple allocations to have read/write access."),
				},
			},
			IsOptional: true,
		},
		"attachment_mode": {
			Description: lang.Markdown("Specifies the name or reference of the attribute to examine for the affinity. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars)."),
			Constraint: schema.OneOf{
				schema.LiteralValue{Value: cty.StringVal("file-system")},
				schema.LiteralValue{Value: cty.StringVal("block-device")},
			},
			IsOptional: true,
		},
	},
}
