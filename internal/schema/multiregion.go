package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var MultiregionSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"strategy": {
			Description: lang.PlainText("Specifies a rollout strategy for the regions."),
			Body:        StrategySchema,
		},
		"region": {
			Description: lang.PlainText("Specifies the parameters for a specific region. This can be specified multiple times to define the set of regions for the multi-region deployment. Regions are ordered; depending on the rollout strategy Nomad may roll out to each region in order or to several at a time."),
			Body:        RegionSchema,
		},
	},
}

var StrategySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"max_parallel": {
			Description: lang.Markdown("Specifies the maximum number of region deployments that a multi-region will have in a running state at a time. By default, Nomad will deploy all regions simultaneously."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"on_failure": {
			Description:  lang.Markdown("Specifies the behavior when a region deployment fails. Available options are `fail_all`, `fail_local`, or the default (empty `\"\"`). This field and its interactions with the job's [`update` block](https://developer.hashicorp.com/nomad/docs/job-specification/update) is described in the [examples](https://developer.hashicorp.com/nomad/docs/job-specification/multiregion#examples) below."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}

var RegionSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"count": {
			Description: lang.Markdown("Specifies a count override for task groups in the region. If a task group specifies a `count = 0`, its count will be replaced with this value. If a task group specifies its own `count` or omits the `count` field, this value will be ignored. This value must be non-negative."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"datacenters": {
			Description: lang.Markdown("A list of datacenters in the region which are eligible for task placement. If not provided, the `datacenters` field of the job will be used."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"node_pool": {
			Description: lang.Markdown("The node pool to be used in this region. It overrides the job-level node_pool and the namespace default node pool."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"meta": {
			Description: lang.PlainText("The meta block allows for user-defined arbitrary key-value pairs. The meta specified for each region will be merged with the meta block at the job level."),
			Body:        MetaSchema,
		},
	},
}
