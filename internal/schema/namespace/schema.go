package namespace

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("Specifies the namespace to create or update."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"description": {
			Description: lang.Markdown("Specifies an optional human-readable description of the namespace."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"quota": {
			Description: lang.Markdown("Specifies a quota to attach to the namespace" + schemautils.Divider + schemautils.EnterpriseOnly),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"meta": {
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Description: lang.Markdown("Optional object with string keys and values of metadata to attach to the namespace. Namespace metadata is not used by Nomad and is intended for use by operators and third party tools."),
					Constraint:  schema.LiteralType{Type: cty.String},
					IsOptional:  true,
				},
			},
		},
		"capabilities": {
			Description: CapabilitiesSchema.Description,
			Body:        CapabilitiesSchema,
		},
		"node_pool_config": {
			Description: NodePoolConfig.Description,
			Body:        NodePoolConfig,
		},
		"vault": {
			Description: VaultConfig.Description,
			Body:        VaultConfig,
		},
		"consul": {
			Description: ConsulConfig.Description,
			Body:        ConsulConfig,
		},
	},
}

var CapabilitiesSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled_task_drivers": {
			Description:  lang.Markdown("List of task drivers allowed in the namespace. If empty all task drivers are allowed."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"disabled_task_drivers": {
			Description:  lang.Markdown("List of task drivers disabled in the namespace."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"enabled_network_modes": {
			Description:  lang.Markdown("List of network modes allowed in the namespace. If empty all network modes are allowed."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"disabled_network_modes": {
			Description:  lang.Markdown("List of network modes disabled in the namespace."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}

var NodePoolConfig = &schema.BodySchema{
	Description: lang.Markdown("node pool config docs" + schemautils.Divider + schemautils.EnterpriseOnly),
	Attributes: map[string]*schema.AttributeSchema{
		"default": {
			Description:  lang.Markdown("Specifies the node pool to use for jobs or dynamic host volumes in this namespace that don't define a node pool in their specification."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"allowed": {
			Description: lang.Markdown("Specifies the node pools that jobs or dynamic host volumes in this namespace are allowed to use. By default, all node pools are allowed. If an empty list is provided only the namespace's default node pool is allowed. This field supports wildcard globbing through the use of `*` for multi-character matching. This field cannot be used with `denied`"),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"denied": {
			Description:  lang.Markdown("Specifies the node pools that jobs or dynamic host volumes in this namespace are not allowed to use. This field supports wildcard globbing through the use of `*` for multi-character matching. If specified, jobs and dynamic host volumes are allowed to use any node pool, except for those that match any of these patterns. This field cannot be used with `allowed`."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}

var VaultConfig = &schema.BodySchema{
	Description: lang.Markdown("vault docs" + schemautils.Divider + schemautils.EnterpriseOnly),
	Attributes: map[string]*schema.AttributeSchema{
		"default": {
			Description:  lang.Markdown("Specifies the Vault cluster to use for jobs in this namespace that don't define a Vault cluster in their specification."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"allowed": {
			Description: lang.Markdown("Specifies the Vault clusters that are allowed to be used by jobs in this namespace. By default, all Vault clusters are allowed. If an empty list is provided only the namespace's default Vault cluster is allowed. This field supports wildcard globbing through the use of `*` for multi-character matching. This field cannot be used with `denied`."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"denied": {
			Description:  lang.Markdown("Specifies the Vault clusters that are not allowed to be used by jobs in this namespace. This field supports wildcard globbing through the use of `*` for multi-character matching. If specified, any Vault cluster is allowed to be used, except for those that match any of these patterns. This field cannot be used with `allowed`."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}

var ConsulConfig = &schema.BodySchema{
	Description: lang.Markdown("consul docs" + schemautils.Divider + schemautils.EnterpriseOnly),
	Attributes: map[string]*schema.AttributeSchema{
		"default": {
			Description:  lang.Markdown("Specifies the Consul cluster to use for jobs in this namespace that don't define a Consul cluster in their specification."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"allowed": {
			Description: lang.Markdown("Specifies the Consul clusters that are allowed to be used by jobs in this namespace. By default, all Consul clusters are allowed. If an empty list is provided only the namespace's default Consul cluster is allowed. This field supports wildcard globbing through the use of `*` for multi-character matching. This field cannot be used with `denied`."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"denied": {
			Description:  lang.Markdown("Specifies the Consul clusters that are not allowed to be used by jobs in this namespace. This field supports wildcard globbing through the use of `*` for multi-character matching. If specified, any Consul cluster is allowed to be used, except for those that match any of these patterns. This field cannot be used with `allowed`."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}
