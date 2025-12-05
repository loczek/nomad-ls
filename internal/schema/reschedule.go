package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// TODO: find actuall default values for all of the attributes
var RescheduleSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"attempts": {
			Description: lang.Markdown("Specifies the number of reschedule attempts allowed in the configured interval. Defaults vary by job type, see below for more information."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(15),
			},
			Constraint: &schema.LiteralType{Type: cty.Number},
		},
		"interval": {
			Description: lang.Markdown("Specifies the sliding window which begins when the first reschedule attempt starts and ensures that only `attempts` number of reschedule happen within it. If more than `attempts` number of failures happen with this interval, Nomad will not reschedule any more."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"delay": {
			Description: lang.Markdown("Specifies the duration to wait before attempting to reschedule a failed task. This is specified using a label suffix like \"30s\" or \"1h\". Delay cannot be less than 5 seconds."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"delay_function": {
			Description: lang.Markdown("Specifies the function that is used to calculate subsequent reschedule delays. The initial delay is specified by the delay parameter. `delay_function` has three possible values which are described below.\n- `constant` - The delay between reschedule attempts stays constant at the delay value.\n- `exponential` - The delay between reschedule attempts doubles.\n- `fibonacci` - The delay between reschedule attempts is calculated by adding the two most recent delays applied. For example if delay is set to 5 seconds, the next five reschedule attempts will be delayed by 5 seconds, 5 seconds, 10 seconds, 15 seconds, and 25 seconds respectively."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"max_delay": {
			Description: lang.Markdown("is an upper bound on the delay beyond which it will not increase. This parameter is used when `delay_function` is `exponential` or `fibonacci`, and is ignored when `constant `delay is used."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"unlimited": {
			Description: lang.Markdown("`unlimited` enables unlimited reschedule attempts. If this is set to `true` the `attempts` and `interval` fields are not used. The [`progress_deadline`](https://developer.hashicorp.com/nomad/docs/job-specification/update#progress_deadline) parameter within the update block is still adhered to when this is set to `true`, meaning no more reschedule attempts are triggered once the [`progress_deadline`](https://developer.hashicorp.com/nomad/docs/job-specification/update#progress_deadline) is reached."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(true),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
	},
}
