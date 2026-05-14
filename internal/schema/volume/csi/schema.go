package csi

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/schema/volume"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"capacity_min": {
			Description: lang.Markdown("Option for requesting a minimum capacity, in bytes. The capacity of a volume may be the physical size of a disk, or a quota, depending on the storage provider. The specific size of the resulting volume is somewhere between `capacity_min` and `capacity_max`; the exact behavior is up to the storage provider. If you want to specify an exact size, you should set `capacity_min` and `capacity_max` to the same value. Accepts human-friendly suffixes such as `\"100GiB\"`. This field may not be supported by all storage providers. Increasing this value and reissuing `volume create` or `volume register` may expand the volume, if the CSI plugin supports it."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"capacity_max": {
			Description: lang.Markdown("Option for requesting a maximum capacity, in bytes. The capacity of a volume may be the physical size of a disk, or a quota, depending on the storage provider. The specific size of the resulting volume is somewhere between `capacity_min` and `capacity_max`; the exact behavior is up to the storage provider. If you want to specify an exact size, you should set `capacity_min` and `capacity_max` to the same value. Accepts human-friendly suffixes such as `\"100GiB\"`. This field may not be supported by all storage providers."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"clone_id": {
			Description: lang.Markdown("If the storage provider supports cloning, the external ID of the volume to clone when creating this volume. If omitted, the volume is created from scratch. The `clone_id` cannot be set if the `snapshot_id` field is set. Only allowed on volume creation."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"external_id": {
			Description: lang.Markdown("The ID of the physical volume from the storage provider. For example, the volume ID of an AWS EBS volume or Digital Ocean volume. Only allowed on volume registration."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"id": {
			Description: lang.Markdown("The unique ID of the volume. This is how the [`volume.source`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#source) field in a job specification refers to the volume."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"name": {
			Description: lang.Markdown("The name of the volume. On volume creation, the external storage provider may use this field to tag the volume or as an idempotency token, so it must be unique across all namespaces."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"namespace": {
			Description: lang.Markdown("The namespace of the volume. This field overrides the namespace provided by the `-namespace` flag or `NOMAD_NAMESPACE` environment variable. Defaults to `\"default\"` if unset."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"plugin_id": {
			Description: lang.Markdown("The ID of the [CSI plugin](https://developer.hashicorp.com/nomad/docs/job-specification/csi_plugin) that manages this volume."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"snapshot_id": {
			Description: lang.Markdown("If the storage provider supports snapshots, the external ID of the snapshot to restore when creating this volume. If omitted, the volume is created from scratch. The snapshot_id cannot be set if the clone_id field is set. Only allowed on volume creation."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"type": {
			Description: lang.Markdown("The type of volume. Must be `\"csi\"` for CSI volumes."),
			Constraint:  schema.LiteralValue{Value: cty.StringVal("csi")},
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
		"context": {
			Description: lang.Markdown("An optional key-value map of strings passed directly to the CSI plugin to validate the volume. The details of these parameters are specific to each storage provider, so consult the specific plugin documentation for more information. Only allowed on volume registration. Note that, like the rest of the volume specification, this block is declarative, and an update replaces it in its entirety, therefore all parameters need to be specified."),
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
		"parameters": {
			Description: lang.Markdown("An optional key-value map of strings passed directly to the CSI plugin to configure the volume. The details of these parameters are specific to each storage provider, so consult the specific plugin documentation for more information."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
					IsOptional: true,
				},
			},
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
