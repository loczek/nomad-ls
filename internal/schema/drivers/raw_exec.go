package drivers

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var RawExecDriverSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"args": {
			Description: lang.Markdown("A list of arguments to the `command`. References to environment variables or any [interpretable Nomad variables](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation) will be interpreted before launching the task."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.List(cty.String)},
				schema.AnyExpression{OfType: cty.List(cty.String)},
			},
			IsOptional: true,
		},
		"command": {
			Description: lang.Markdown("The command to execute. Must be provided. If executing a binary that exists on the host, the path must be absolute. If executing a binary that is downloaded from an [`artifact`](https://developer.hashicorp.com/nomad/docs/job-specification/artifact), the path can be relative from the allocation's root directory."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsRequired: true,
		},
		// TODO: check if type is correct
		"cgroup_v1_override": {
			Description: lang.Markdown("A map of controller names to paths. The task will be added to these cgroups. The task will fail if these cgroups do not exist. **WARNING:** May conflict with other Nomad driver's cgroups and have unintended side effects."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Map(cty.String)},
				schema.AnyExpression{OfType: cty.Map(cty.String)},
			},
			IsOptional: true,
		},
		// TODO: check if type is correct
		"cgroup_v2_override": {
			Description: lang.Markdown("A map of controller names to paths. The task will be added to these cgroups. The task will fail if these cgroups do not exist. **WARNING:** May conflict with other Nomad driver's cgroups and have unintended side effects."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Map(cty.String)},
				schema.AnyExpression{OfType: cty.Map(cty.String)},
			},
			IsOptional: true,
		},
		"oom_score_adj": {
			Description:  lang.Markdown("A positive integer to indicate the likelihood of the task being OOM killed (valid only for Linux). Defaults to 0."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
		"work_dir": {
			Description: lang.Markdown("Sets a custom working directory for the task. This must be an absolute path. This will also change the working directory when using `nomad alloc exec`."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"denied_envvars": {
			Description: lang.Markdown("Passes a list of environment variables that the driver should scrub from the task environment. Supports globbing, with \"*\" wildcard accepted as prefix and/or suffix."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.List(cty.String)},
				schema.AnyExpression{OfType: cty.List(cty.String)},
			},
			IsOptional: true,
		},
	},
}
