package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var PeriodicSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cron": {
			Description: lang.Markdown("Specifies a cron expression configuring the interval to launch the job. In addition to [cron-specific formats](https://github.com/hashicorp/cronexpr#implementation), this option also includes predefined expressions such as `@daily` or `@weekly`. Either `cron` or `crons` must be set, but not both."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.ListVal([]cty.Value{cty.StringVal("")}),
			},
			IsDeprecated: true,
			Constraint:   &schema.LiteralType{Type: cty.List(cty.String)},
		},
		"crons": {
			Description: lang.Markdown("A list of cron expressions configuring the intervals the job is launched at. The job runs at the next earliest time that matches any of the expressions. Supports predefined expressions such as `@daily` and `@weekly`. Refer to [the documentation](https://github.com/hashicorp/cronexpr#implementation) for full details about the supported cron specs and the predefined expressions. Either `cron` or `crons` must be set, but not both."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.ListVal([]cty.Value{cty.StringVal("")}),
			},
			Constraint: &schema.LiteralType{Type: cty.List(cty.String)},
		},
		"prohibit_overlap": {
			Description: lang.Markdown("Specifies if this job should wait until previous instances of this job have completed. This only applies to this job; it does not prevent other periodic jobs from running at the same time."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
		"time_zone": {
			Description: lang.Markdown("Specifies the time zone to evaluate the next launch interval against. [Daylight Saving Time](https://developer.hashicorp.com/nomad/docs/job-specification/periodic#daylight-saving-time) affects scheduling, so please ensure the [behavior below](https://developer.hashicorp.com/nomad/docs/job-specification/periodic#daylight-saving-time) meets your needs. The time zone must be parsable by Golang's [LoadLocation](https://golang.org/pkg/time/#LoadLocation)."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("UTC"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"enabled": {
			Description: lang.Markdown("Specifies if this job should run. This not only prevents this job from running on the `cron` schedule but prevents force launches."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(true),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
	},
}
