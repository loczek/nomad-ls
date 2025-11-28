package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var GroupSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"count": {
			Description: lang.Markdown("Specifies the number of instances that should be running under for this group. This value must be non-negative. This defaults to the `min` value specified in the [`scaling`](https://developer.hashicorp.com/nomad/docs/job-specification/scaling) block, if present; otherwise, this defaults to `1`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(1),
			},
		},
		"shutdown_delay": {
			Description: lang.PlainText("Specifies the duration to wait when stopping a group's tasks. The delay occurs between Consul or Nomad service deregistration and sending each task a shutdown signal. Ideally, services would fail health checks once they receive a shutdown signal. Alternatively, shutdown_delay may be set to give in-flight requests time to complete before shutting down. A group level shutdown_delay will run regardless if there are any defined group services and only applies to these services. In addition, tasks may have their own shutdown_delay which waits between de-registering task services and stopping the task."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("0s"),
			},
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"constraint": {
			Description: lang.PlainText("This can be provided multiple times to define additional constraints."),
			Body:        ConstraintSchema,
		},
		"affinity": {
			Description: lang.PlainText("This can be provided multiple times to define preferred placement criteria."),
			Body:        AffinitySchema,
		},
		"spread": {
			Description: lang.Markdown("This can be provided multiple times to define criteria for spreading allocations across a node attribute or metadata. See the [Nomad spread reference](https://developer.hashicorp.com/nomad/docs/job-specification/spread) for more details."),
			Body:        SpreadSchema,
		},
		"consul": {
			Description: lang.Markdown("Specifies Consul configuration options specific to the group. These options will be applied to all tasks and services in the group unless a task has its own `consul` block."),
			Body:        ConsulSchema,
		},
		"ephemeral_disk": {
			Description: lang.PlainText("Specifies the ephemeral disk requirements of the group. Ephemeral disks can be marked as sticky and support live data migrations."),
			Body:        EphemeralDiskSchema,
		},
		"disconnect": {
			Description: lang.PlainText("Specifies the disconnect strategy for the server and client for all tasks in this group in case of a network partition. The tasks can be left unconnected, stopped or replaced when the client disconnects. The policy for reconciliation in case the client regains connectivity is also specified here."),
		},
		"meta": {
			Description: lang.PlainText("Specifies a key-value map that annotates the group with user-defined metadata."),
			Body:        MetaSchema,
		},
		"migrate": {
			Description: lang.PlainText("Specifies the group strategy for migrating off of draining nodes. Only service jobs with a count greater than 1 support migrate blocks."),
			Body:        MigrateSchema,
		},
		// TODO: there is no way to make the label optional so it had to be removed
		"network": {
			Description: lang.PlainText("Specifies the network requirements and configuration, including static and dynamic port allocations, for the group."),
			Body:        NetworkSchema,
		},
		"reschedule": {
			Description: lang.PlainText("Allows to specify a rescheduling strategy. Nomad will then attempt to schedule the task on another node if any of the group allocation statuses become \"failed\"."),
			Body:        RescheduleSchema,
		},
		"restart": {
			Description: lang.PlainText("Specifies the restart policy for all tasks in this group. If omitted, a default policy exists for each job type, which can be found in the [restart block documentation](https://developer.hashicorp.com/nomad/docs/job-specification/restart)."),
			Body:        RestartSchema,
		},
		"service": {
			Description: lang.Markdown("Specifies integrations with Nomad or [Consul](https://developer.hashicorp.com/nomad/docs/configuration/consul) for service discovery. Nomad automatically registers each service when an allocation is started and de-registers them when the allocation is destroyed."),
			Body:        ServiceSchema,
		},
		// TODO: make it required
		"task": {
			Description: lang.PlainText("Specifies one or more tasks to run within this group. This can be specified multiple times, to add a task as part of the group."),
			Labels: []*schema.LabelSchema{
				{
					Name: "name",
				},
			},
			Body: TaskSchema,
		},
		"update": {
			Description: lang.PlainText("Specifies the task's update strategy. When omitted, a default update strategy is applied."),
			Body:        UpdateSchema,
		},
		"vault": {
			Description: lang.PlainText("Specifies the set of Vault policies required by all tasks in this group. Overrides a `vault` block set at the `job` level."),
		},
		"volume": {
			Description: lang.PlainText("Specifies the volumes that are required by tasks within the group."),
			Body:        VolumeSchema,
			Labels: []*schema.LabelSchema{
				{
					Name: "name",
				},
			},
		},
	},
}
