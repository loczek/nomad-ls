package dynamic

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/schema/job"
	"github.com/loczek/nomad-ls/internal/schema/volume"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"capacity": {
			Description: lang.Markdown("The size of a volume in bytes. Either the physical size of a disk or a quota, depending on the plugin. This field must be between the `capacity_min` and `capacity_max` values unless they are omitted. Accepts human-friendly suffixes such as `\"100GiB\"`. Only supported for volume registration."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"capacity_min": {
			Description: lang.Markdown("Option for requesting a minimum capacity, in bytes. The capacity of a volume may be the physical size of a disk, or a quota, depending on the plugin. The specific size of the resulting volume is somewhere between `capacity_min` and `capacity_max`; the exact behavior is up to the plugin. If you want to specify an exact size, set `capacity_min` and `capacity_max` to the same value. Accepts human-friendly suffixes such as `\"100GiB\"`. Plugins that cannot restrict the size of volumes, such as the built-in [mkdir](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host#mkdir-plugin) plugin, may ignore this field."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"capacity_max": {
			Description: lang.Markdown("Option for requesting a maximum capacity, in bytes. The capacity of a volume may be the physical size of a disk, or a quota, depending on the plugin. The specific size of the resulting volume is somewhere between `capacity_min` and `capacity_max`; the exact behavior is up to the plugin. If you want to specify an exact size, set `capacity_min` and `capacity_max` to the same value. Accepts human-friendly suffixes such as `\"100GiB\"`. Plugins that cannot restrict the size of volumes, such as the built-in [mkdir](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host#mkdir-plugin) plugin, may ignore this field."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"id": {
			Description: lang.Markdown("The ID of a previously created volume to update via volume create or volume register. You should never set this field when initially creating or registering a volume, and you should only use the values returned from the Nomad API for the ID."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"host_path": {
			Description: lang.Markdown("The path on disk where the volume exists. You should set this only for volume registration. It is ignored for volume creation."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"name": {
			Description: lang.Markdown("The name of the volume, which is used as the [`volume.source`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#source) field in job specifications that claim this volume. Host volume names must be unique per node. Names are visible to any user with `node:read` ACL, even across namespaces, so they should not be treated as sensitive values."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"namespace": {
			Description: lang.Markdown("The namespace of the volume. This field overrides the namespace provided by the `-namespace` flag or `NOMAD_NAMESPACE` environment variable. Defaults to `\"default\"` if unset."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: required on registration
		"node_id": {
			Description: lang.Markdown("A specific node where you would like the volume to be created. Refer to the [volume placement](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host#volume-placement) section for details. Optional for volume creation but required for volume registration."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"node_pool": {
			Description: lang.Markdown("A specific node pool where you would like the volume to be created. Refer to the [volume placement](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host#volume-placement) section for details. Optional for volume creation or volume registration. If you also provide `node_id`, the node must be in the provided `node_pool`."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: required on creation
		"plugin_id": {
			Description: lang.Markdown("The ID of the [dynamic host volume plugin](https://developer.hashicorp.com/nomad/docs/architecture/storage/host-volumes) that manages this volume. Required for volume creation. Nomad has one built-in plugin called [`mkdir`](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host#mkdir-plugin)."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"type": {
			Description: lang.Markdown("The type of volume. Must be `\"host\"` for Dynamic Host volumes."),
			Constraint:  schema.LiteralValue{Value: cty.StringVal("host")},
			IsRequired:  true,
			IsDepKey:    true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"capability": {
			Description: lang.PlainText("Option for validating the capability of a volume."),
			Body:        CapabilitySchema,
			MinItems:    1,
		},
		"constraint": {
			Description: lang.Markdown("A restriction on the eligible nodes where a volume can be created. Refer to the [volume placement](https://developer.hashicorp.com/nomad/docs/other-specifications/volume/host#volume-placement) section for details. You can provide multiple `constraint` blocks to add more constraints. Optional for volume creation and ignored for volume registration."),
			Body:        job.ConstraintSchema,
		},
		"parameters": {
			Description: lang.Markdown("An optional key-value map of strings passed directly to the plugin to configure the volume. The details of these parameters are specific to the plugin."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
					IsOptional: true,
				},
			},
		},

		"mount_options": {
			Description: lang.PlainText("Options for mounting file-system volumes that don't already have a pre-formatted file system."),
			Body:        volume.MountOptionsSchema,
			MinItems:    1,
		},
		"secrets": {
			Description: lang.Markdown("An optional key-value map of strings used as credentials for publishing and unpublishing volumes."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
					IsOptional: true,
				},
			},
		},
		"topology_request": {
			Description: lang.Markdown("Specify locations such as region, zone, and rack where the provisioned volume must be accessible from in the case of volume creation, or the locations where the existing volume is accessible from in the case of volume registration."),
			Body:        volume.TopologyRequestSchema,
		},
	},
}
