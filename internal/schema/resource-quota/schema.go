package schemaResourceQuota

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("The name of the Quota. Nomad uses name to connect the quota a [`Namespace`](https://developer.hashicorp.com/nomad/docs/other-specifications/namespace)."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"description": {
			Description: lang.Markdown("A human-readable description."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"limit": {
			Description: QuotaLimitSchema.Description,
			Body:        QuotaLimitSchema,
		},
	},
}

var QuotaLimitSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"region": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"region_limit": {
			Description: RegionLimitSchema.Description,
			Body:        RegionLimitSchema,
		},
	},
}

var RegionLimitSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cores": {
			Description: lang.Markdown("The limit on total number of CPU cores from all `resources.cores` in the namespace. The [CPU concepts](https://developer.hashicorp.com/nomad/docs/architecture/cpu) documentation has more details on CPU resources."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"cpu": {
			Description: lang.Markdown("The limit on total amount of CPU from all `resources.cpu` in the namespace."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"memory": {
			Description: lang.Markdown("The limit on total mount of memory in MB from all `resources.memory` in the namespace."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"memory_max": {
			Description: lang.Markdown("The limit on total mount of hard memory limits in MB from all `resources.memory_max` in the namespace."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"device": {
			Description: DeviceSchema.Description,
			Body:        DeviceSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
		"storage": {
			Description: StorageSchema.Description,
			Body:        StorageSchema,
		},
		"node_pool": {
			Description: NodePoolSchema.Description,
			Body:        NodePoolSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
	},
}

var DeviceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"count": {
			Description: lang.Markdown("How many of this device may be used."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
	},
}

var StorageSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"host_volumes": {
			Description: lang.Markdown("Maximum total size of all [dynamic host volumes](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host) in MiB. The default `0` means unlimited, and `-1` means variables are fully disabled. This field accepts human-friendly string inputs such as \"100 GiB\". The quota for host volumes is enforced at the time the volume is created via `volume create`."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"variables": {
			Description: lang.Markdown("Maximum total size of all Nomad [variables](https://developer.hashicorp.com/nomad/docs/concepts/variabless) in MiB. The default `0` means unlimited, and `-1` means variables are fully disabled. This field accepts human-friendly string inputs such as \"100 GiB\"."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
	},
}

// TODO: the actuall docs are missing for this body schema (https://github.com/hashicorp/nomad/blob/375134e88aaf4846c96761778cc91ffbd689c627/api/quota.go#L157-L167)
var NodePoolSchema = &schema.BodySchema{}
