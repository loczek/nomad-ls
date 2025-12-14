package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var VaultSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"allow_token_expiration": {
			Description:  lang.Markdown("Specifies that Nomad clients should not attempt to renew a task's Vault token, allowing it to expire. This should only be used when a secret is requested from Vault once at the start of a task or in a short-lived prestart task. Long-running tasks should never set `allow_token_expiration=true` if they obtain Vault secrets via `template` blocks, as the Vault token will expire and the template runner will continue to make failing requests to Vault until its [`vault_retry`](https://developer.hashicorp.com/nomad/docs/configuration/client#vault_retry) attempts are exhausted, at which point the task will fail.\n\nWhen Nomad has been configured to use [Workload Identity with Vault](https://developer.hashicorp.com/nomad/docs/secure/vault/acl#nomad-workload-identities), Nomad clients will automatically detect when tokens cannot be refreshed (for example, when the Vault auth method is configured to issue batch tokens). In this case, the `allow_token_expiration` option will be implicitly set to `true` by the client. The [legacy Vault authentication workflow](https://developer.hashicorp.com/nomad/docs/v1.8.x/integrations/vault/acl) cannot automatically detect this."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"change_mode": {
			Description:  lang.Markdown("Specifies the behavior Nomad should take if the Vault token changes. The possible values are:\n\n- [`\"noop\"`](https://developer.hashicorp.com/nomad/docs/job-specification/vault#noop) - take no action (continue running the task)\n- [`\"restart\"`](https://developer.hashicorp.com/nomad/docs/job-specification/vault#restart) - restart the task\n- [`\"signal\"`](https://developer.hashicorp.com/nomad/docs/job-specification/vault#signal) - send a configurable signal to the task"),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("restart")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"change_signal": {
			Description:  lang.Markdown("Specifies the signal to send to the task as a string like `\"SIGUSR1\"` or `\"SIGINT\"`. This option is required if the `change_mode` is `signal`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: mark as enterprise only
		"cluster": {
			Description:  lang.Markdown("Specifies the Vault cluster to use. The Nomad client will retrieve a Vault token from the cluster configured in the agent configuration with the same [`vault.name`](https://developer.hashicorp.com/nomad/docs/configuration/vault#name). In Nomad Community Edition, this field is ignored."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"env": {
			Description:  lang.Markdown("Specifies if the `VAULT_TOKEN` and `VAULT_NAMESPACE` environment variables should be set when starting the task."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		// TODO: add warning
		"disable_file": {
			Description:  lang.Markdown("Specifies if the Vault token should be written to `secrets/vault_token`."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		// TODO: mark as enterprise only
		"namespace": {
			Description:  lang.Markdown("Specifies the Vault Namespace to use for the task. The Nomad client will retrieve a Vault token that is scoped to this particular namespace."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"role": {
			Description:  lang.Markdown("Specifies the Vault role used when retrieving a token from Vault using JWT and workload identity. If not specified the client's [`create_from_role`](https://developer.hashicorp.com/nomad/docs/configuration/vault#create_from_role) value is used."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
