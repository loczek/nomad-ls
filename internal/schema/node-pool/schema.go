package nodePool

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"node_pool": {
			Description: NodePoolSchema.Description,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: NodePoolSchema,
		},
	},
}

var NodePoolSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"description": {
			Description: lang.Markdown("Sets a human readable description for the node pool."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"node_identity_ttl": {
			Description: lang.Markdown("Sets the TTL for node identities issued to nodes in this node pool. The value must be a valid duration string (e.g. \"30m\", \"1h\", \"24h\"). If not set, the default value is \"24h\""),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"meta": {
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Description: lang.Markdown("Sets optional metadata on the node pool, defined as key-value pairs. The scheduler does not use node pool metadata as part of scheduling."),
					Constraint:  schema.LiteralType{Type: cty.String},
					IsOptional:  true,
				},
			},
		},
		"scheduler_config": {
			Description: SchedulerConfigSchema.Description,
			Body:        SchedulerConfigSchema,
		},
	},
}

var SchedulerConfigSchema = &schema.BodySchema{
	Description: lang.Markdown("scheduler config docs" + schemautils.Divider + schemautils.EnterpriseOnly),
	Attributes: map[string]*schema.AttributeSchema{
		"description": {
			Description: lang.Markdown("Sets a human readable description for the node pool."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"node_identity_ttl": {
			Description: lang.Markdown("Sets the TTL for node identities issued to nodes in this node pool. The value must be a valid duration string (e.g. \"30m\", \"1h\", \"24h\"). If not set, the default value is \"24h\""),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
