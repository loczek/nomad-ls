package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var VolumeSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description: lang.Markdown("Specifies the type of a given volume. The valid volume types are `host` and `csi`. Setting the `host` value can request either statically configured host volumes or dynamic host volumes, depending on what is available on a given node."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"source": {
			Description: lang.Markdown("The name of the volume to request. When using `host_volume`'s this should match the published name of the host volume. When using `csi` volumes, this should match the ID of the registered volume."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"read_only": {
			Description: lang.Markdown("Specifies that the group only requires read only access to a volume and is used as the default value for the `volume_mount -> read_only` configuration. This value is also used for validating `host_volume` ACLs and for scheduling when a matching `host_volume` requires `read_only` usage."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
		"sticky": {
			Description: lang.Markdown("Specifies that this volume sticks to the allocation that uses it. Upon every reschedule and replacement, the task group always receives the volume with the same ID, if available. Use sticky volumes for [stateful deployments](https://developer.hashicorp.com/nomad/docs/concepts/stateful-deployments). You may only use the `sticky` field for dynamic host volumes. For CSI volumes, the `per_alloc` field provides similar functionality"),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
		"per_alloc": {
			Description: lang.Markdown("Specifies that the `source` of the volume should have the suffix `[n]`, where `n` is the allocation index. This allows mounting a unique volume per allocation, so long as the volume's source is named appropriately. For example, with the source `myvolume` and `per_alloc = true`, the allocation named `myjob.mygroup.mytask[0]` will require a volume ID `myvolume[0]`. The `per_alloc` field cannot be true for system jobs, sysbatch jobs, or jobs that use canaries. `per_alloc` is mutually exclusive with the `sticky` property. Use `per_alloc` only with CSI volumes and `sticky` only with dynamic host volumes."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
		"access_mode": {
			Description: lang.Markdown("Defines whether a volume should be available concurrently. The `access_mode` and `attachment_mode` together must exactly match one of the volume's `capability` blocks.\n\nFor CSI volumes the `access_mode` is required. Can be one of the following:\n\n- [`\"single-node-reader-only\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#single-node-reader-only)\n- [`\"single-node-writer\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#single-node-writer)\n- [`\"multi-node-reader-only\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#multi-node-reader-only)\n- [`\"multi-node-single-writer\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#multi-node-single-writer)\n- [`\"multi-node-multi-writer\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#multi-node-multi-writer)\n\nMost CSI plugins support only single-node modes. Consult the documentation of the storage provider and CSI plugin.\n    \nFor dynamic host volumes the `access_mode` is optional. Can be one of the following:\n\n- [`\"single-node-writer\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#single-node-writer-1)\n- [`\"single-node-reader-only\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#single-node-reader-only-1)\n- [`\"single-node-single-writer\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#single-node-single-writer)\n- [`\"single-node-multi-writer\"`](https://developer.hashicorp.com/nomad/docs/job-specification/volume#single-node-multi-writer)\n\nDefaults to `single-node-writer` unless `read_only = true`, in which case it defaults to `single-node-reader-only`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"attachment_mode": {
			Description: lang.Markdown("The storage API used by the volume. One of `\"file-system\"` or `\"block-device\"`. The `access_mode` and `attachment_mode` together must exactly match one of the volume's `capability` blocks.\n-For CSI volumes the `attachment_mode` field is required. Most storage providers support `\"file-system\"`, to mount volumes using the CSI filesystem API. Some storage providers support `\"block-device\"`, which mounts the volume with the CSI block device API within the container.\n-For dynamic host volumes the `attachment_mode` field is optional and defaults to `\"file-system\"`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		// TODO: check if constraint is correct
		// TODO: change mount_flags to cty.List?
		"mount_options": {
			Description: lang.Markdown("Options for mounting CSI volumes that have the `file-system` [attachment mode](https://developer.hashicorp.com/nomad/commands/volume/register#attachment_mode). These options override the `mount_options` field from [volume registration](https://developer.hashicorp.com/nomad/commands/volume/register#mount_options). Consult the documentation for your storage provider and CSI plugin as to whether these options are required or necessary."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.MapVal(map[string]cty.Value{
					"fs_type":     cty.StringVal("ext4"),
					"mount_flags": cty.StringVal("[\"ro\", \"noatime\"]"),
				}),
			},
			Constraint: &schema.LiteralType{Type: cty.Map(cty.String)},
		},
	},
}
