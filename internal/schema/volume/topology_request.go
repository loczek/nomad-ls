package volume

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var TopologyRequestSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"required": {
			Description: lang.Markdown("On **volume creation**, the `required` topologies indicate that the volume must be created in a location accessible from at least one of the listed topologies. On **volume registration** the `required` topologies indicate that the volume was created in a location accessible from all the listed topologies."),
			Body:        TopologySchema,
		},
		"prefered": {
			Description: lang.Markdown("Indicate that you would prefer the storage provider to create the volume in one of the provided topologies. Only allowed on **volume creation**."),
			Body:        TopologySchema,
		},
	},
}

var TopologySchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"segments": {
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Description: lang.Markdown("A map of location types to their values. The specific fields required are defined by the CSI plugin."),
					Constraint:  schema.LiteralType{Type: cty.String},
					IsOptional:  true,
				},
			},
		},
	},
}
