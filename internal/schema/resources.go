package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ResourcesSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cpu": {
			Description: lang.Markdown("Specifies the CPU required to run this task in MHz."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(100),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
		},
		"cores": {
			Description: lang.Markdown("Specifies the number of CPU cores to reserve specifically for the task. This may not be used with `cpu`. The behavior of setting `cores` is specific to each task driver (e.g. [docker](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/docker#cpu), [exec](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#cpu))."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
			IsOptional: true,
		},
		"memory": {
			Description: lang.Markdown("Specifies the memory required in MB."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(300),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
		},
		"memory_max": {
			Description: lang.Markdown("Optionally, specifies the maximum memory the task may use, if the client has excess memory capacity, in MB. See [Memory Oversubscription](https://developer.hashicorp.com/nomad/docs/job-specification/resources#memory-oversubscription) for more details."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(300),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
			IsOptional: true,
		},
		"secret": {
			Description: lang.Markdown("Specifies the size of the [`secrets/`](https://developer.hashicorp.com/nomad/docs/reference/runtime-environment-settings#secrets) directory in MB, on platforms where the directory is a tmpfs. If set, the scheduler adds the `secrets` value to the `memory` value when allocating resources on a client, and this value will be included in the allocated resources shown by the `nomad alloc status` and `nomad node status` commands. If unset, the client will allocate 1 MB of tmpfs space and it will not be counted for scheduling purposes or included in allocated resources. You should not set this value if the workload will be placed on a platform where tmpfs is unsupported, because it will still be counted for scheduling purposes."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
	},
	// TODO: add numa block
	Blocks: map[string]*schema.BlockSchema{
		"device": {
			Description: lang.Markdown("test"),
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: DeviceSchema,
		},
	},
}

var DeviceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("Specifies the device required. The following inputs are valid:\n\n- `<device_type>`: If a single value is given, it is assumed to be the device type, such as \"gpu\", or \"fpga\"\n- `<vendor>/<device_type>`: If two values are given separated by a /, the given device type will be selected, constraining on the provided vendor. Examples include \"nvidia/gpu\" or \"amd/gpu\"\n- `<vendor>/<device_type>/<model>`: If three values are given separated by a /, the given device type will be selected, constraining on the provided vendor, and model name. Examples include \"nvidia/gpu/1080ti\" or \"nvidia/gpu/2080ti\""),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"count": {
			Description: lang.Markdown("Specifies the number of instances of the given device that are required."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(1),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
			IsOptional: true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"constraint": {
			Description: lang.Markdown("Constraints to restrict which devices are eligible. This can be provided multiple times to define additional constraints. See below for available attributes."),
			Body:        ConstraintSchema,
			MinItems:    1,
		},
		"affinity": {
			Description: lang.Markdown(" Affinity to specify a preference for which devices get selected. This can be provided multiple times to define additional affinities. See below for available attributes."),
			Body:        AffinitySchema,
		},
	},
}
