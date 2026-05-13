package job

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var IdentitySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("The name of the workload identity, which must be unique per task. Only one `identity` block in a task can omit the `name` field."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("default"),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		// TODO: fix default value
		"aud": {
			Description: lang.Markdown("The audience field for the workload identity. This should always be set for non-default identities."),
			// DefaultValue: schema.DefaultValue{
			// 	Value: cty.List(cty.String),
			// },
			IsOptional: true,
		},
		// TODO: update default values
		"change_mode": {
			Description: lang.Markdown("Specifies the behavior Nomad should take when the token changes."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("noop"),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		// TODO: this should be required when `change_mode` is set to `signal`
		"change_signal": {
			Description: lang.Markdown("Specifies the signal to send to the task as a string like \"SIGHUP\" or \"SIGUSR1\". This option is required if the `change_mode` is `signal`."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"env": {
			Description: lang.Markdown("If true the workload identity will be available in the task's NOMAD_TOKEN environment variable."),
			DefaultValue: schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
		"file": {
			Description: lang.Markdown("If true the workload identity will be available in the task's filesystem via the path `secrets/nomad_token`. If the `task.user` parameter is set, the token file will only be readable by that user. Otherwise the file is readable by everyone but is protected by parent directory permissions."),
			DefaultValue: schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
		"filepath": {
			Description: lang.Markdown("If not empty and file is `true`, the workload identity is available at the specified location relative to the [task working directory](https://developer.hashicorp.com/nomad/docs/reference/runtime-environment-settings#task-directories) instead of the `NOMAD_SECRETS_DIR`.\"30s\" or \"1h\". You may not set a TTL on the default identity. You should always set a TTL for non-default identities.\"30s\" or \"1h\". You may not set a TTL on the default identity. You should always set a TTL for non-default identities."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"ttl": {
			Description: lang.Markdown("The lifetime of the identity before it expires. The client will renew the identity at roughly half the TTL. This is specified using a label suffix like \"30s\" or \"1h\". You may not set a TTL on the default identity. You should always set a TTL for non-default identities.\"30s\" or \"1h\". You may not set a TTL on the default identity. You should always set a TTL for non-default identities."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
	},
}
