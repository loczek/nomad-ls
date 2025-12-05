package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var MigrateSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"max_parallel": {
			Description: lang.Markdown("Specifies the number of allocations that can be migrated at the same time. This number must be less than the total [`count`](https://developer.hashicorp.com/nomad/docs/job-specification/group#count) for the group as `count - max_parallel` will be left running during migrations."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(1),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
		},
		"health_checks": {
			Description: lang.Markdown("Specifies the mechanism in which allocations health is determined. The potential values are:\n- \"checks\" - Specifies that the allocation should be considered healthy when all of its tasks are running and their associated checks are healthy, and unhealthy if any of the tasks fail or not all checks become healthy. This is a superset of \"task_states\" mode.\n- \"task_states\" - Specifies that the allocation should be considered healthy when all its tasks are running unhealthy if tasks fail."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("checks"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"min_healthy_time": {
			Description: lang.Markdown("Specifies the minimum time the allocation must be in the healthy state before it is marked as healthy and unblocks further allocations from being migrated. This is specified using a label suffix like \"30s\" or \"15m\"."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("10s"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"healthy_deadline": {
			Description: lang.Markdown("Specifies the deadline in which the allocation must be marked as healthy after which the allocation is automatically transitioned to unhealthy. This is specified using a label suffix like \"2m\" or \"1h\"."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("5m"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
	},
}
