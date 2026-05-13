package job

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var LifecycleSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		// TODO: update docs
		"hook": {
			Description: lang.Markdown("Specifies when a task should be run within the lifecycle of a group. The following hooks are available:\n- `prestart` - Will be started immediately. The main tasks will not start until all prestart tasks with sidecar = false have completed successfully."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("default"),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsRequired: true,
		},
		"sidecar": {
			Description: lang.Markdown("Controls whether a task is ephemeral or long-lived within the task group. If a lifecycle task is ephemeral (`sidecar = false`), the task will not be restarted after it completes successfully. If a lifecycle task is long-lived (`sidecar = true`) and terminates, it will be restarted as long as the allocation is running."),
			DefaultValue: schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
	},
}
