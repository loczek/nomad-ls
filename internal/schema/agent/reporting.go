package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ReportingSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"snapshot_retention_time": {
			Description:  lang.Markdown("Configures the maximum amount of time that Nomad retains a utilization reporting snapshot in the Nomad state store. You can export these snapshots with the nomad operator utilization command."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("9600h")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"disable_product_usage_reporting": {
			Description:  lang.Markdown("Specifies whether detailed product usage metrics should be disabled. Review the full list of metrics before disabling this option, and share any concerns you may have with your account manager."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"license": {
			Description: LicenseSchema.Description,
			Body:        LicenseSchema,
		},
	},
}

var LicenseSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Specifies whether automated license utilization reporting should be enabled and run. License utilization metrics are still gathered for offline reporting."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
