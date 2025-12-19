package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ScheduleSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"cron": {
			Description: lang.Markdown("The autoscaling policy. This is opaque to Nomad, consumed and parsed only by the external autoscaler. Therefore, its contents are specific to the autoscaler; consult the [Nomad Autoscaler documentation](https://developer.hashicorp.com/nomad/tools/autoscaling/policy) for more details."),
			Body:        CronSchema,
		},
	},
}

var CronSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"start": {
			Description: lang.Markdown("When the task should be started. Specified in 6 field [cron format](https://github.com/hashicorp/cronexpr#implementation) (no seconds) without `,` or `/` characters."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"end": {
			Description: lang.Markdown("When the task should be stopped ([`kill_signal`](https://developer.hashicorp.com/nomad/docs/job-specification/task#kill_signal) and [`kill_timeout`](https://developer.hashicorp.com/nomad/docs/job-specification/task#kill_timeout) apply). Specified in 2 field [cron format](https://github.com/hashicorp/cronexpr#implementation) (minute and hour) without `,` or `/` characters."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"timezone": {
			Description:  lang.Markdown("What time zone the `start` and `end` times are specified in. Defaults to the local time zone of the Nomad Client the job is scheduled onto."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("Local")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
