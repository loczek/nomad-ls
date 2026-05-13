package acl

import (
	"github.com/hashicorp/hcl-lang/schema"
)

var RootSchema = schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"namespace": {
			Description: NamespaceSchema.Description,
			Body:        NamespaceSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
		"node": {
			Description: NodeSchema.Description,
			Body:        NodeSchema,
		},
		"node_pool": {
			Description: NodePoolSchema.Description,
			Body:        NodePoolSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
		"agent": {
			Description: AgentSchema.Description,
			Body:        AgentSchema,
		},
		"operator": {
			Description: OperatorSchema.Description,
			Body:        OperatorSchema,
		},
		"quota": {
			Description: QuotaSchema.Description,
			Body:        QuotaSchema,
		},
		"host_volume": {
			Description: HostVolumeSchema.Description,
			Body:        HostVolumeSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
		"plugin": {
			Description: PluginSchema.Description,
			Body:        PluginSchema,
		},
		"sentinel": {
			Description: SentinelSchema.Description,
			Body:        SentinelSchema,
		},
	},
}
