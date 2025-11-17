package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var JobSchemaBetter = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"all_at_once": {
			Description: lang.PlainText("Controls whether the scheduler can make partial placements if optimistic scheduling resulted in an oversubscribed node. This does not control whether all allocations for the job, where all would be the desired count for each task group, must be placed atomically. This should only be used for special circumstances."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
		},
		"datacenters": {
			Description: lang.Markdown("A list of datacenters in the region which are eligible for task placement. This field allows wildcard globbing through the use of * for multi-character matching. The default value is [\"*\"], which allows the job to be placed in any available datacenter."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.ListVal([]cty.Value{cty.StringVal("*")}),
			},
		},
		"node_pool": {
			Description: lang.Markdown("Specifies the node pool to place the job in. The node pool must exist when the job is registered. Defaults to \"default\"."),
			IsOptional:  true,
		},
		"name": {
			Description: lang.PlainText("Specifies a name for the job, which otherwise defaults to the job ID."),
			IsOptional:  true,
		},
		"namespace": {
			Description: lang.PlainText("The namespace in which to execute the job."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("default"),
			},
		},
		"priority": {
			Description: lang.Markdown("Specifies the job priority which is used to prioritize scheduling and access to resources. Must be between 1 and `job_max_priority` inclusively, with a larger value corresponding to a higher priority. If value 0 is provided this will fallback to `job_default_priority`. Priority only has an effect when job preemption is enabled. It does not have an effect on which of multiple pending jobs is run first."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(50),
			},
		},
		"region": {
			Description: lang.PlainText("The region in which to execute the job."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("global"),
			},
		},
		"type": {
			Description: lang.PlainText("Specifies the Nomad scheduler to use. Nomad provides the service, `system`, `batch`, and `sysbatch` schedulers."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("service"),
			},
		},
		// TODO: Update with docs later
		"vault_token": {
			Description: lang.PlainText("Specifies the Vault token used for job submission. Strongly discouraged to place in config."),
			IsOptional:  true,
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
		},
		// TODO: Update with docs later
		"consul_token": {
			Description: lang.PlainText("Specifies the Consul token used for job submission."),
			IsOptional:  true,
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"constraint": {
			Description: lang.Markdown("This can be provided multiple times to define additional constraints. See the [Nomad constraint reference](https://developer.hashicorp.com/nomad/docs/job-specification/constraint) for more details."),
		},
		"affinity": {
			Description: lang.Markdown("This can be provided multiple times to define preferred placement criteria. See the [Nomad affinity reference](https://developer.hashicorp.com/nomad/docs/job-specification/affinity) for more details."),
		},
		"spread": {
			Description: lang.PlainText("This can be provided multiple times to define criteria for spreading allocations across a node attribute or metadata. See the [Nomad spread reference](https://developer.hashicorp.com/nomad/docs/job-specification/spread) for more details."),
			Body:        SpreadSchema,
		},
		// TODO: make it required
		"group": {
			Description: lang.PlainText("Specifies the start of a group of tasks. This can be provided multiple times to define additional groups. Group names must be unique within the job file."),
			Body:        GroupSchema,
			Labels: []*schema.LabelSchema{
				{
					Name: "name",
				},
			},
		},
		// TODO: should it be a block?
		"meta": {
			Description: lang.PlainText("Specifies a key-value map that annotates with user-defined metadata."),
		},
		"migrate": {
			Description: lang.PlainText("Specifies the groups strategy for migrating off of draining nodes. If omitted, a default migration strategy is applied. Only service jobs with a count greater than 1 support migrate blocks."),
		},
		"parameterized": {
			Description: lang.PlainText("Specifies the job as a parameterized job such that it can be dispatched against."),
		},
		"periodic": {
			Description: lang.PlainText("Allows the job to be scheduled at fixed times, dates or intervals."),
		},
		"reschedule": {
			Description: lang.PlainText("Allows to specify a rescheduling strategy. Nomad will then attempt to schedule the task on another node if any of its allocation statuses become \"failed\"."),
		},
		"update": {
			Description: lang.PlainText("Specifies the task's update strategy. When omitted, a default update strategy is applied."),
		},
		"vault": {
			Description: lang.PlainText("Specifies the set of Vault policies required by all tasks in this job."),
		},
	},
	Description: lang.PlainText("job schema docs"),
}
