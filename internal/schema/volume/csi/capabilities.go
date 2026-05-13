package csi

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
					Value:       cty.StringVal("single-node-reader-only"),
					Description: lang.PlainText("Jobs can only request the volume with read-only access, and only one node can mount the volume at a time."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("single-node-writer"),
					Description: lang.PlainText("Jobs can request the volume with read/write or read-only access, and only one node can mount the volume at a time."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("multi-node-reader-only"),
					Description: lang.PlainText("Jobs can only request the volume with read-only access, but multiple nodes can mount the volume simultaneously."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("multi-node-single-writer"),
					Description: lang.PlainText("Jobs can request the volume with read/write or read-only access, but the scheduler only allows one allocation to have read/write access. Multiple nodes can mount the volume simultaneously."),
				},
				schema.LiteralValue{
					Value:       cty.StringVal("multi-node-multi-writer"),
					Description: lang.PlainText("Jobs can request the volume with read/write or read-only access, and the scheduler allows multiple allocations to have read/write access. Multiple nodes can mount the volume simultaneously."),
				},
			},
			IsRequired: true,
		},
		"attachment_mode": {
			Description: lang.Markdown("Specifies the name or reference of the attribute to examine for the affinity. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars)."),
			Constraint: schema.OneOf{
				schema.LiteralValue{Value: cty.StringVal("file-system")},
				schema.LiteralValue{Value: cty.StringVal("block-device")},
			},
			IsRequired: true,
		},
	},
}
