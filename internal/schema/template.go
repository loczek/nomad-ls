package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var TemplateSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		// TODO: add examples
		"change_mode": {
			Description:  lang.PlainText("Specifies the behavior Nomad should take if the rendered template changes. Nomad will always write the new contents of the template to the specified destination. The following possible values describe Nomad's action after writing the template to disk."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("restart")},
			Constraint:   &schema.LiteralType{Type: cty.String},
		},
		"change_signal": {
			Description:  lang.PlainText("Specifies the signal to send to the task as a string like `\"SIGUSR1\"` or `\"SIGINT\"`. This option is required if the `change_mode` is `signal`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
		},
		"data": {
			Description:  lang.Markdown("Specifies the raw template to execute. One of `source` or `data` must be specified, but not both. This is useful for smaller templates, but we recommend using `source` for larger templates."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"destination": {
			Description:  lang.Markdown("Specifies the location where the resulting template should be rendered, relative to the [task working directory](/nomad/docs/reference/runtime-environment-settings#task-directories). Only drivers without filesystem isolation (ex. `raw_exec`) or that build a chroot in the task working directory (ex. `exec`) can render templates outside of the `NOMAD_ALLOC_DIR`, `NOMAD_TASK_DIR`, or `NOMAD_SECRETS_DIR`. For more details on how `destination` interacts with task drivers, see the [Filesystem internals](/nomad/docs/concepts/filesystem#templates-artifacts-and-dispatch-payloads) documentation. Note that where possible, the `NOMAD_SECRETS_DIR` is mounted `noexec`, so rendered templates can't be used as self-executing scripts."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"env": {
			Description:  lang.Markdown("Specifies the template should be read back in as environment variables for the task ([example](/nomad/docs/job-specification/template#environment-variables)). To update the environment on changes, you must set `change_mode` to `restart`. Setting `env` when the `change_mode` is `signal` will return a validation error. Setting `env` when the `change_mode` is `noop` is permitted but will not update the environment variables in the task."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		// TODO: add examples
		"error_on_missing_key": {
			Description:  lang.Markdown("Specifies how the template behaves when attempting to index a map key that does not exist in the map."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"left_delimiter": {
			Description:  lang.Markdown("Specifies the left delimiter to use in the template. The default is \"{{\" for some templates, it may be easier to use a different delimiter that does not conflict with the output file itself."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("{{")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"perms": {
			Description:  lang.Markdown("Specifies the rendered template's permissions. File permissions are given as octal of the Unix file permissions `rwxrwxrwx`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("644")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add warning from docs
		"uid": {
			Description: lang.Markdown("Specifies the rendered template owner's user ID. If negative or not specified (`nil`) the ID of the Nomad agent user will be used."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"gid": {
			Description: lang.Markdown("Specifies the rendered template owner's group ID. If negative or not specified (`nil`) the ID of the Nomad agent group will be used."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"once": {
			Description:  lang.Markdown("Specifies that the client will wait for the template to be rendered, and then no longer watch the dependencies specified in the template. This is useful for templates that do not need to be updated. Once mode implicitly disables wait/quiescence timers for this template, and ignores change mode/signal/script settings."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"right_delimiter": {
			Description:  lang.Markdown("Specifies the right delimiter to use in the template. The default is \"}}\" for some templates, it may be easier to use a different delimiter that does not conflict with the output file itself."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("}}")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"source": {
			Description:  lang.Markdown("Specifies the path to the template to be rendered. One of `source` or `data` must be specified, but not both. This source can be fetched using an [`artifact`](/nomad/docs/job-specification/artifact) resource. The template must exist in the [task working directory](/nomad/docs/reference/runtime-environment-settings#task-directories) prior to starting the task; it is not possible to reference a template whose source is inside a Docker container, for example."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"splay": {
			Description:  lang.Markdown("Specifies a random amount of time to wait between 0 ms and the given splay value before invoking the change mode. This is specified using a label suffix like \"30s\" or \"1h\", and is often used to prevent a thundering herd problem where all task instances restart at the same time."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"vault_grace": {
			Description:  lang.Markdown("deprecated"),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("15s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsDeprecated: true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"change_script": {
			Description: lang.PlainText("Configures the script triggered on template change. This option is required if the `change_mode` is `script`."),
			Body:        ChangeScriptSchema,
		},
		"wait": {
			Description: lang.PlainText("Defines the minimum and maximum amount of time to wait for the Consul cluster to reach a consistent state before rendering a template. This is useful to enable in systems where network connectivity to Consul is degraded, because it will reduce the number of times a template is rendered. This setting can be overridden by the [`client.template.wait_bounds`](/nomad/docs/configuration/client#wait_bounds). If the template configuration has a `min` lower than `client.template.wait_bounds.min` or a `max` greater than `client.template.wait_bounds.max`, the client's bounds will be enforced, and the template `wait` will be adjusted before being sent to the template engine."),
			Body:        WaitSchema,
		},
	},
}

var WaitSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"min": {
			Description:  lang.Markdown("min wait time"),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"max": {
			Description:  lang.Markdown("max wait time"),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("10s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
