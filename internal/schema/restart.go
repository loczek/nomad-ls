package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// TODO: find actuall default values for all of the attributes
var RestartSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"attempts": {
			Description: lang.Markdown("Specifies the number of restarts allowed in the configured interval. Defaults vary by job type, see below for more information."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
		},
		"delay": {
			Description: lang.Markdown("Specifies the duration to wait before restarting a task. This is specified using a label suffix like \"30s\" or \"1h\". A random jitter of up to 25% is added to the delay."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("15s"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"interval": {
			Description: lang.Markdown("Specifies the duration which begins when the first task starts and ensures that only `attempts` number of restarts happens within it. If more than `attempts` number of failures happen, behavior is controlled by `mode`. This is specified using a label suffix like \"30s\" or \"1h\". Defaults vary by job type, see below for more information.\"30s\" or \"1h\". A random jitter of up to 25% is added to the delay."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"mode": {
			Description: lang.Markdown("Controls the behavior when the task fails more than `attempts` times in an interval. For a detailed explanation of these values and their behavior, please see the [mode values section](https://developer.hashicorp.com/nomad/docs/job-specification/restart#mode-values)."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("fail"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"render_templates": {
			Description: lang.Markdown("Specifies whether to re-render all templates when a task is restarted. If set to `true`, all templates will be re-rendered when the task restarts. This can be useful for re-fetching Vault secrets, even if the lease on the existing secrets has not yet expired."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
	},
}
