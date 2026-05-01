package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var TemplateSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"function_denylist": {
			Description: lang.Markdown("Specifies a list of template rendering functions that should be disallowed in job specs. By default the `plugin`, `executeTemplate` and `writeToFile` functions are disallowed as they allow unrestricted root access to the host or allow recursive execution."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("plugin"),
				cty.StringVal("executeTemplate"),
				cty.StringVal("writeToFile"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"disable_file_sandbox": {
			Description:  lang.Markdown("Allows templates access to arbitrary files on the client host via the `file` function. By default, templates can access files only within the [task working directory](https://developer.hashicorp.com/nomad/docs/reference/runtime-environment-settings#task-directories)."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"max_stale": {
			Description:  lang.Markdown("This is the maximum interval to allow \"stale\" data. If `max_stale` is set to `0`, only the Consul leader will respond to queries, and requests that reach a follower will forward to the leader. In large clusters with many requests, this is not as scalable. This option allows any follower to respond to a query, so long as the last-replicated data is within this bound. Higher values result in less cluster load, but are more likely to have outdated data. This default of 10 years (`87600h`) matches the default Consul configuration."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("87600h")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"block_query_wait": {
			Description:  lang.Markdown("This is amount of time in seconds to wait for the results of a blocking query. Many endpoints in Consul support a feature known as \"blocking queries\". A blocking query is used to wait for a potential change using long polling."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"wait": {
			Description: lang.Markdown("Defines the minimum and maximum amount of time to wait before attempting to re-render a template. Consul Template re-renders templates whenever rendered variables from Consul, Nomad, or Vault change. However in order to minimize how often tasks are restarted or reloaded, Nomad will configure Consul Template with a backoff timer that will tick on an interval equal to the specified `min` value. Consul Template will always wait at least the as long as the `min` value specified. If the underlying data has not changed between two tick intervals, Consul Template will re-render. If the underlying data has changed, Consul Template will delay re-rendering until the underlying data stabilizes for at least one tick interval, or the configured `max` duration has elapsed. Once the `max` duration has elapsed, Consul Template will re-render the template with the data available at the time. This is useful to enable in systems where Consul is in a degraded state, or the referenced data values are changing rapidly, because it will reduce the number of times a template is rendered. Setting both `min` and `max` to `0` disables the feature. This configuration is also exposed in the task template block to allow overrides per task."),
			Body:        WaitSchema,
		},
		"wait_bounds": {
			Description: lang.Markdown("Defines client level lower and upper bounds for per-template `wait` configuration. If the individual template configuration has a `min` lower than `wait_bounds.min` or a `max` greater than the `wait_bounds.max`, the bounds will be enforced, and the template `wait` will be adjusted before being sent to `consul-template`."),
			Body:        WaitBoundsSchema,
		},
		"consul_retry": {
			Description: lang.Markdown("This controls the retry behavior when an error is returned from Consul. The template runner will not exit in the face of failure. Instead, it uses exponential back-off and retry functions to wait for the Consul cluster to become available, as is customary in distributed systems."),
			Body:        NomadVaultConsulRetry,
		},
		"vault_retry": {
			Description: lang.Markdown("This controls the retry behavior when an error is returned from Vault. Consul Template is highly fault tolerant, meaning it does not exit in the face of failure. Instead, it uses exponential back-off and retry functions to wait for the cluster to become available, as is customary in distributed systems."),
			Body:        NomadVaultConsulRetry,
		},
		"nomad_retry": {
			Description: lang.Markdown("This controls the retry behavior when an error is returned from Nomad. Consul Template is highly fault tolerant, meaning it does not exit in the face of failure. Instead, it uses exponential back-off and retry functions to wait for the cluster to become available, as is customary in distributed systems."),
			Body:        NomadVaultConsulRetry,
		},
	},
}

var WaitSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"min": {
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"max": {
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("4m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}

var WaitBoundsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"min": {
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"max": {
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("10s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}

var NomadVaultConsulRetry = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"attempts": {
			Description:  lang.Markdown("This specifies the number of attempts to make before giving up. Each attempt adds the exponential backoff sleep time. Setting this to zero will implement an unlimited number of retries."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(12)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"backoff": {
			Description:  lang.Markdown("This is the base amount of time to sleep between retry attempts. Each retry sleeps for an exponent of 2 longer than this base. For 5 retries, the sleep times would be: 250ms, 500ms, 1s, 2s, then 4s."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("250ms")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"max_backoff ": {
			Description:  lang.Markdown("This is the maximum amount of time to sleep between retry attempts. When max_backoff is set to zero, there is no upper limit to the exponential sleep between retry attempts. If max_backoff is set to 10s and backoff is set to 1s, sleep times would be: 1s, 2s, 4s, 8s, 10s, 10s, ..."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("1m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
