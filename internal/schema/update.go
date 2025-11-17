package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// TODO: check docs
var UpdateSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"max_parallel": {
			Description: lang.Markdown(`Specifies the number of allocations within a task
group that can be destructively updated at the same time. The task groups themselves are
updated in parallel. In-place updates are performed all at once.

- max_parallel = 0 - Specifies that the allocation should use forced updates instead of deployments

~> **Note:** System jobs only support one allocation per node. When canary is set, system job updates make up to max_parallel destructive updates *or* destructively update enough allocations to place on canary percent of feasible nodes, whichever is lower. You should adjust the max_parallel value to allow deploying all desired canaries.`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(1),
			},
		},
		"health_check": {
			Description: lang.Markdown(`Specifies the mechanism in which allocations health is determined. The potential values are:

- "checks" - Specifies that the allocation should be considered healthy when all of its tasks are running and their associated [checks][] are healthy, and unhealthy if any of the tasks fail or not all checks become healthy. This is a superset of "task_states" mode.
- "task_states" - Specifies that the allocation should be considered healthy when all its tasks are running and unhealthy if tasks fail.
- "manual" - Specifies that Nomad should not automatically determine health and that the operator will specify allocation health using the [HTTP API](https://developer.hashicorp.com/nomad/api-docs/deployments#set-allocation-health-in-deployment).`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("checks"),
			},
		},
		"min_healthy_time": {
			Description: lang.PlainText(`Specifies the minimum time the allocation must be in the healthy state before it is marked as healthy and unblocks further allocations from being updated. This is specified using a label suffix like "30s" or "15m".`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("10s"),
			},
		},
		"healthy_deadline": {
			Description: lang.PlainText(`Specifies the deadline in which the allocation must be marked as healthy after which the allocation is automatically transitioned to unhealthy. This is specified using a label suffix like "2m" or "1h". If progress_deadline is non-zero, it must be greater than healthy_deadline. Otherwise the progress_deadline may fail a deployment before an allocation reaches its healthy_deadline.`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("5m"),
			},
		},
		"progress_deadline": {
			Description: lang.PlainText(`Specifies the deadline in which an allocation must be marked as healthy. The deadline begins when the first allocation for the deployment is created and is reset whenever an allocation as part of the deployment transitions to a healthy state or when a deployment is manually promoted. If no allocation transitions to the healthy state before the progress deadline, the deployment is marked as failed. If the progress_deadline is set to 0, the first allocation to be marked as unhealthy causes the deployment to fail. This is specified using a label suffix like "2m" or "1h".`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("10m"),
			},
		},
		"auto_revert": {
			Description: lang.PlainText(`Specifies if the job should auto-revert to the last stable job on deployment failure. A job is marked as stable if all the allocations as part of its deployment were marked healthy.`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
		},
		"auto_promote": {
			Description: lang.PlainText(`Specifies if the job should auto-promote to the canary version when all canaries become healthy during a deployment. Defaults to false which means canaries must be manually updated with the nomad deployment promote command. If a job has multiple task groups, all must be set to auto_promote = true in order for the deployment to be promoted automatically.`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
		},
		"canary": {
			Description: lang.Markdown(`Specifies that changes to the job that would result in destructive updates should create the specified number of canaries without stopping any previous allocations. Once the operator determines the canaries are healthy, they can be promoted which unblocks a rolling update of the remaining allocations at a rate of max_parallel. Canary deployments cannot be used with volumes when per_alloc = true.

In system jobs, the canary setting indicates the percentage of feasible nodes to which Nomad makes destructive allocation updates. System jobs do not support more than one allocation per node, so effectively setting canary to a positive integer means this percentage of feasible nodes gets a new version of the job if the update is destructive. Non-destructive updates ignore the canary field. Setting canary to 100 updates the job on all nodes. Percentage of nodes is always rounded up to the nearest integer. If canary is set, nodes that register during a deployment do not receive placements until after the deployment is promoted.

~> **Note:** Updates when canary is set make up to max_parallel destructive updates *or* destructively update enough allocations to place on canary percent of feasible nodes, whichever is lower. You should adjust the max_parallel value to allow deploying all desired canaries.`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
		},
		"stagger": {
			Description: lang.PlainText(`Specifies the delay between each set of max_parallel updates when updating system jobs. This setting is being deprecated, and is equivalent to min_healthy_time.`),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("30s"),
			},
		},
	},
}
