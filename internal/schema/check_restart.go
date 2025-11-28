package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var CheckRestartSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"limit": {
			Description:  lang.Markdown("Restart task when a health check has failed `limit` times. For example 1 causes a restart on the first failure. The default, `0`, disables health check based restarts. Failures must be consecutive. A single passing check will reset the count, so flapping services may not be restarted."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
		},
		"grace": {
			Description:  lang.Markdown("Duration to wait after a task starts or restarts before checking its health."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("1s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"ignore_warnings": {
			Description:  lang.Markdown("By default checks with both `critical` and `warning` statuses are considered unhealthy. Setting `ignore_warnings = true` treats a `warning` status like `passing` and will not trigger a restart. Only available in the Consul service provider."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
