package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var DrainOnShutdownSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"deadline": {
			Description:  lang.Markdown("Set the deadline by which all allocations must be moved off the client. Remaining allocations after the deadline are removed from the client, regardless of their `migrate` block. Defaults to 1 hour."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("1h")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"force": {
			Description:  lang.Markdown("Setting to `true` drains all the allocations on the client immediately, ignoring the `migrate` block. Note if you have multiple allocations for the same job on the draining client without additional allocations on other clients, this will result in an outage for that job until the drain is complete."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"ignore_system_jobs": {
			Description:  lang.Markdown("Setting to true allows the drain to complete without stopping system job allocations. By default system jobs (and CSI plugins) are stopped last."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
