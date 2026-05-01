package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var FingerprintSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description:  lang.Markdown("The name of the fingerprinter the configuration applies to. This field is required and is part of the block label. It supports `\"env_aws\"`, `\"env_azure\"`, `\"env_digitalocean\"`, and `\"env_gce\"`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"retry_attempts": {
			Description:  lang.Markdown("Specifies the maximum number of retry attempts to make before giving up. This field supports `-1` for unlimited retries, but we do not recommend `-1` as this may cause the client to hang indefinitely on startup if the fingerprinting source is unavailable. The default value of `0` means no retries are made."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"retry_interval": {
			Description:  lang.Markdown("Specifies the amount of time to wait between retry attempts."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("2s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add example
		"exit_on_failure": {
			Description:  lang.Markdown("Specifies whether the client should exit if fingerprinter fails to probe the envrionment endpoint service."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
