package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var CheckSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"address_mode": {
			Description:  lang.Markdown("Specifies which address (host, alloc, alloc IPv6 or driver-specific) this service should use to make checks, similar to `address_mode` on `service`. When the service uses the Consul provider, Consul needs access to the address for any HTTP or TCP checks. Unlike `port`, this setting is **not** inherited from the `service`. If the service `address` is set, the service `address_mode` is `\"auto\"`, and the check `address_mode` is not set, Nomad uses the service `address` value for the check address.\n\nValid options are:\n\n* `alloc` - For allocations which create a network namespace, this address mode uses the IP address inside the namespace. Use only with \"bridge\" and \"cni\" networking modes. You may specify a numeric port for situations where no port mapping is necessary. Set this mode only for checks which are defined in a \"group\" service block.\n* `alloc_ipv6` - Same as `alloc` but use the IPv6 address in case of dual-stack or IPv6-only.\n* `driver` - Use the IP specified by the driver, and the port specified in a port map. You may specify a numeric port since port maps aren't required by all network plugins. Useful for checking SDN and overlay network addresses. Task fails if driver network cannot be determined. Only implemented for Docker. Set this mode only for checks which are defined in a \"task\" service block. Refer to Using driver address mode for an example of use.\n* `host` - Use the host IP and the exposed port.\n\nNote there is no `\"auto\"` mode for checks, unlike services."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("host")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"args": {
			Description:  lang.Markdown("Specifies additional arguments to the `command`. This only applies to script-based health checks."),
			DefaultValue: &schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"command": {
			Description: lang.Markdown("Specifies the command to run for performing the health check. The script must exit: 0 for passing, 1 for warning, or any other value for a failing health check. This is required for script-based health checks. Only supported in the Consul service provider.\n\n**Caveat:** The command must be the path to the command on disk, and no shell exists by default. That means operators like `||` or `&&` are not available. Additionally, all arguments must be supplied via the `args` parameter. To achieve the behavior of shell operators, specify the command as a shell, like `/bin/bash` and then use `args` to run the check."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"grpc_service": {
			Description: lang.Markdown("What service, if any, to specify in the gRPC health check. gRPC health checks require Consul 1.0.5 or later."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"grpc_use_tls": {
			Description:  lang.Markdown("Use TLS to perform a gRPC health check. May be used with `tls_skip_verify` to use TLS but skip certificate verification. May be used with `tls_server_name` to specify the ServerName to use for SNI and validation of the certificate presented by the server being checked."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"initial_status": {
			Description: lang.Markdown("Specifies the starting status of the service. Valid options are `passing`, `warning`, and `critical`. Omitting this field (or submitting an empty string) will result in the Consul default behavior, which is `critical`. Only supported in the Consul service provider. In the Nomad service provider, the initial status of a check is `pending` until Nomad produces an initial check status result."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"success_before_passing": {
			Description:  lang.Markdown("The number of consecutive successful checks required before Consul will transition the service status to `passing`. Only supported by the Consul service provider and not applicable for health checks of type \"script\"."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"failures_before_critical": {
			Description:  lang.Markdown("The number of consecutive failing checks required before Consul will transition the service status to `critical`. Only supported by the Consul service provider and not applicable for health checks of type \"script\"."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"failures_before_warning": {
			Description:  lang.Markdown("The number of consecutive failing checks required before Consul will transition the service status to `warning`. Only supported by the Consul service provider and not applicable for health checks of type \"script\"."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"interval": {
			Description: lang.Markdown("Specifies the frequency of the health checks that Consul or Nomad service provider will perform. This is specified using a label suffix like \"30s\" or \"1h\". This must be greater than or equal to \"1s\"."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"method": {
			Description:  lang.Markdown("Specifies the HTTP method to use for HTTP checks. Must be a valid HTTP method."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("GET")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"body": {
			Description:  lang.Markdown("Specifies the HTTP body to use for HTTP checks."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"name": {
			Description: lang.Markdown("Specifies the name of the health check. If the name is not specified Nomad generates one based on the service name."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"path": {
			Description: lang.Markdown("Specifies the path of the HTTP endpoint which will be queried to observe the health of a service. Nomad will automatically add the IP of the service and the port, so this is just the relative URL to the health check endpoint. This is required for HTTP-based health checks."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"expose": {
			Description:  lang.Markdown("Specifies whether an [Expose Path](/nomad/docs/job-specification/expose#path-parameters) should be automatically generated for this check. Only compatible with Connect-enabled task-group services using the default Connect proxy. If set, check [`type`](/nomad/docs/job-specification/check#type) must be `http` or `grpc`, and check `name` must be set. Only supported in the Consul service provider."),
			DefaultValue: &schema.DefaultValue{Value: cty.False},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"port": {
			Description: lang.Markdown("Specifies the label of the port on which the check will be performed. Note this is the _label_ of the port and not the port number unless `address_mode = driver`. The port label must match one defined in the [`network`](/nomad/docs/job-specification/network) block. If a port value was declared on the `service`, this will inherit from that value if not supplied. If supplied, this value takes precedence over the `service.port` value. This is useful for services which operate on multiple ports. `grpc`, `http`, and `tcp` checks require a port while `script` checks do not. Checks will use the host IP and ports by default. Numeric ports may be used if `address_mode=\"driver\"` is set on the check."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"protocol": {
			Description:  lang.Markdown("Specifies the protocol for the HTTP-based health checks. Valid options are `http` and `https`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("http")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"task": {
			Description:  lang.Markdown("Specifies the task associated with this check. Scripts are executed within the task's environment, and `check_restart` blocks will apply to the specified task. Inherits the [`service.task`](/nomad/docs/job-specification/service#task-1) value if not set. Must be unset or equivalent to `service.task` in task-level services."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"timeout": {
			Description: lang.Markdown("Specifies how long to wait for a health check query to succeed. This is specified using a label suffix like \"30s\" or \"1h\". This must be greater than or equal to \"1s\".\n\n**Caveat:** Script checks use the task driver to execute in the task's environment. For task drivers with namespace isolation such as `docker` or `exec`, setting up the context for the script check may take an unexpectedly long amount of time (a full second or two), especially on busy hosts. The timeout configuration must allow for both this setup and the execution of the script. Operators should use long timeouts (5 or more seconds) for script checks, and monitor telemetry for `client.allocrunner.taskrunner.tasklet_timeout`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"type": {
			Description: lang.Markdown("This indicates the check types supported by Nomad. For Consul service checks, valid options are `grpc`, `http`, `script`, and `tcp`. For Nomad service checks, valid options are `http` and `tcp`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"tls_server_name": {
			Description:  lang.Markdown("Indicates the ServerName to use for SNI and validation of the certificate presented by the server being checked, when performing TLS enabled checks (`https` and `grpc` with `grpc_use_tls`). If left unspecified, the ServerName will be inferred from the address of the server being checked unless the address is an IP address. There are two common cases where this is beneficial:\n\n* When the check address contains an IP, `tls_server_name` can be specified for SNI. Note: setting `tls_server_name` will also override the hostname used to verify the certificate presented by the server being checked.\n* When the check address contains a hostname which isn't be present in the SAN (Subject Alternative Name) field of the certificate presented by the server being checked. Note: setting `tls_server_name` will also override the hostname used for SNI.\n\nThis field is only supported in the Consul service provider."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_skip_verify": {
			Description:  lang.Markdown("Skip verification of certificates for `https` and `grpc` with `grpc_use_tls` checks."),
			DefaultValue: &schema.DefaultValue{Value: cty.False},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"on_update": {
			Description:  lang.Markdown("Specifies how checks should be evaluated when determining deployment health (including a job's initial deployment). This allows job submitters to define certain checks as readiness checks, progressing a deployment even if the Service's checks are not yet healthy. Checks inherit the Service's value by default. The check status is not altered in the service provider and is only used to determine the check's health during an update.\n\n* `require_healthy` - In order for Nomad to consider the check healthy during an update it must report as healthy.\n* `ignore_warnings` - If a Service Check reports as warning, Nomad will treat the check as healthy. The Check will still be in a warning state in Consul.\n* `ignore` - Any status will be treated as healthy.\n\n**Caveat:** `on_update` is only compatible with certain [`check_restart`](/nomad/docs/job-specification/check_restart) configurations. `on_update = \"ignore_warnings\"` requires that `check_restart.ignore_warnings = true`. `check_restart` can however specify `ignore_warnings = true` with `on_update = \"require_healthy\"`. If `on_update` is set to `ignore`, `check_restart` must be omitted entirely."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("require_healthy")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"check_restart": {
			Description: lang.PlainText("See check_restart block."),
			Body:        CheckRestartSchema,
		},
	},
}
