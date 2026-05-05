package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var EventlogSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description: lang.Markdown("Enable sending Nomad agent logs to the Windows Event Log."),
			Constraint:  schema.LiteralType{Type: cty.Bool},
			IsOptional:  true,
		},
		"level": {
			Description: lang.Markdown("Specifies the verbosity of logs the Nomad agent outputs. Valid log levels include `ERROR`, `WARN`, or `INFO` in increasing order of verbosity. Level must be of equal or less verbosity as defined for the [`log_level`](https://developer.hashicorp.com/nomad/docs/configuration#log_level) parameter."),
			Constraint: schema.OneOf{
				schema.LiteralValue{Value: cty.StringVal("ERROR")},
				schema.LiteralValue{Value: cty.StringVal("WARN")},
				schema.LiteralValue{Value: cty.StringVal("INFO")},
			},
			IsOptional: true,
		},
	},
}
