package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ServiceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"provider": {
			Description:  lang.Markdown("Specifies the service registration provider to use for service registrations. Valid options are either `consul` or `nomad`. All services within a single task group must utilise the same provider value."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("consul")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: mark as enterprise only
		"cluster": {
			Description:  lang.Markdown("Specifies the Consul cluster to use, when the `provider` is `consul`. The Nomad client will retrieve a Consul token from the cluster configured in the agent configuration with the same [`consul.name`](https://developer.hashicorp.com/nomad/docs/configuration/consul#name). In Nomad Community Edition, this field is ignored."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"kind": {
			Description: lang.Markdown("Configures the [Consul Service Kind](https://developer.hashicorp.com/consul/api-docs/agent/service#kind) to pass to Consul during service registration. Only available when `provider = \"consul\"`, and is ignored if a Consul service mesh Gateway is defined."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"name": {
			Description:  lang.Markdown("Specifies the name this service will be advertised as in Consul. If not supplied, this will default to the name of the job, task group, and task concatenated together with a dash, like `\"docs-example-server\"`. Each service must have a unique name within the cluster. Names must adhere to [RFC-1123 §2.1](https://tools.ietf.org/html/rfc1123#section-2) and are limited to alphanumeric and hyphen characters (i.e. `[a-z0-9\\-]`), and be less than 64 characters in length.\n\nIn addition to the standard [Nomad interpolation](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation), the following keys are also available:\n\n - [`${JOB}`](https://developer.hashicorp.com/nomad/docs/job-specification/service#job) - the name of the job\n - [`${TASKGROUP}`](https://developer.hashicorp.com/nomad/docs/job-specification/service#taskgroup) - the name of the task group\n - [`${TASK}`](https://developer.hashicorp.com/nomad/docs/job-specification/service#task) - the name of the task\n - [`${BASE}`](https://developer.hashicorp.com/nomad/docs/job-specification/service#base) - shorthand for `${JOB}-${TASKGROUP}-${TASK}`\n\nValidation of the name occurs in two parts. When the job is registered, an initial validation pass checks that the service name adheres to RFC-1123 §2.1 and the length limit, excluding any variables requiring interpolation. Once the client receives the service and all interpretable values are available, the service name will be interpolated and revalidated. This can cause certain service names to pass validation at submit time but fail at runtime."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("<job>-<taskgroup>-<task>")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add examples
		"port": {
			Description: lang.Markdown("Specifies the port to advertise for this service. The value of `port` depends on which [`address_mode`](https://developer.hashicorp.com/nomad/docs/job-specification/service#address_mode) is being used:\n\n- [`alloc`](https://developer.hashicorp.com/nomad/docs/job-specification/service#alloc) - Advertise the mapped `to` value of the labeled port and the allocation address. If a `to` value is not set, the port falls back to using the allocated host port. The `port` field may be a numeric port or a port label specified in the same group's network block.\n- [`alloc_ipv6`](https://developer.hashicorp.com/nomad/docs/job-specification/service#alloc_ipv6) - Same as `alloc` but use the IPv6 address in case of dual-stack or IPv6-only.\n- [`driver`](https://developer.hashicorp.com/nomad/docs/job-specification/service#driver) - Advertise the port determined by the driver (e.g. Docker). The `port` may be a numeric port or a port label specified in the driver's `ports` field.\n- [`host`](https://developer.hashicorp.com/nomad/docs/job-specification/service#host) - Advertise the host port for this service. `port` must match a port _label_ specified in the [`network`](https://developer.hashicorp.com/nomad/docs/job-specification/network) block."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"tags": {
			Description: lang.Markdown("Specifies the list of tags to associate with this service. If this is not supplied, no tags will be assigned to the service when it is registered."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"canary_tags": {
			Description: lang.Markdown("Specifies the list of tags to associate with this service when the service is part of an allocation that is currently a canary. Once the canary is promoted, the registered tags will be updated to those specified in the `tags` parameter. If this is not supplied, the registered tags will be equal to that of the `tags` parameter."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"enable_tag_override": {
			Description:  lang.Markdown("Enables users of Consul's Catalog API to make changes to the tags of a service without having those changes be overwritten by Consul's anti-entropy mechanism. See Consul [documentation](https://developer.hashicorp.com/consul/docs/concepts/anti-entropy#enable-tag-override) for more information. Only available where `provider = \"consul\"`."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"address": {
			Description: lang.Markdown("Specifies a custom address to advertise in Consul or Nomad service registration. If set, `address_mode` must be in `auto` mode. Useful with interpolation - for example to advertise the public IP address of an AWS EC2 instance set this to `${attr.unique.platform.aws.public-ipv4}`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"tagged_addresses": {
			Description: lang.Markdown("Specifies custom [tagged addresses](https://developer.hashicorp.com/consul/docs/discovery/services#tagged-addresses) to advertise in the Consul service registration. Only available where `provider = \"consul\"`."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"address_mode": {
			Description:  lang.Markdown("Specifies which address (host, alloc, alloc_ipv6 or driver-specific) this service should advertise. See [below for examples.](https://developer.hashicorp.com/nomad/docs/job-specification/service#using-driver-address-mode) Valid options are:\n\n- [`alloc`](https://developer.hashicorp.com/nomad/docs/job-specification/service#alloc-1) - For allocations which create a network namespace, this address mode uses the IP address inside the namespace. Can only be used with \"bridge\" and \"cni\" [networking modes](https://developer.hashicorp.com/nomad/docs/job-specification/network#mode). A numeric port may be specified for situations where no port mapping is necessary. This mode can only be set for services which are defined in a \"group\" block.\n- [`alloc_ipv6`](https://developer.hashicorp.com/nomad/docs/job-specification/service#alloc_ipv6-1) - Same as `alloc` but use the IPv6 address in case of dual-stack or IPv6-only.\n- [`auto`](https://developer.hashicorp.com/nomad/docs/job-specification/service#auto) - Allows the driver to determine whether the host or driver address should be used. Defaults to `host` and only implemented by Docker. If you use a Docker network plugin such as weave, Docker will automatically use its address.\n- [`driver`](https://developer.hashicorp.com/nomad/docs/job-specification/service#driver-1) - Use the IP specified by the driver, and the port specified in a port map. A numeric port may be specified since port maps aren't required by all network plugins. Useful for advertising SDN and overlay network addresses. Task will fail if driver network cannot be determined. Only implemented for Docker. This mode can only be set for services which are defined in a \"task\" block.\n- [`host`](https://developer.hashicorp.com/nomad/docs/job-specification/service#host-1) - Use the host IP and port."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("auto")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"task": {
			Description:  lang.Markdown("Specifies the name of the Nomad task associated with this service definition. Only available on group services. Must be set if this service definition represents a Consul service mesh native service and there is more than one task in the task group."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: check if constraints are correct
		"meta": {
			Description: lang.Markdown("Specifies a key-value map that annotates the Consul service with user-defined metadata. Only available where `provider = \"consul\"`."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.Map(cty.String))},
			IsOptional:  true,
		},
		// TODO: check if constraints are correct
		"canary_meta": {
			Description: lang.Markdown("Specifies a key-value map that annotates the Consul service with user-defined metadata when the service is part of an allocation that is currently a canary. Once the canary is promoted, the registered meta will be updated to those specified in the `meta` parameter. If this is not supplied, the registered meta will be set to that of the `meta` parameter. Only available where `provider = \"consul\"`."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.Map(cty.String))},
			IsOptional:  true,
		},
		"on_update": {
			Description:  lang.Markdown("Specifies how checks should be evaluated when determining deployment health (including a job's initial deployment). This allows job submitters to define certain checks as readiness checks, progressing a deployment even if the Service's checks are not yet healthy. Checks inherit the Service's value by default. The check status is not altered in Consul and is only used to determine the check's health during an update.\n\n- [`require_healthy`](https://developer.hashicorp.com/nomad/docs/job-specification/service#require_healthy) - In order for Nomad to consider the check healthy during an update it must report as healthy.\n[`ignore_warnings`](https://developer.hashicorp.com/nomad/docs/job-specification/service#ignore_warnings) - If a Service Check reports as warning, Nomad will treat the check as healthy. The Check will still be in a warning state in Consul.\n- [`ignore`](https://developer.hashicorp.com/nomad/docs/job-specification/service#ignore) - Any status will be treated as healthy.\n\n**Caveat:** `on_update` is only compatible with certain [`check_restart`](https://developer.hashicorp.com/nomad/docs/job-specification/check_restart) configurations. `on_update = \"ignore_warnings\"` requires that `check_restart.ignore_warnings = true`. `check_restart` can however specify `ignore_warnings = true` with `on_update = \"require_healthy\"`. If `on_update` is set to `ignore`, `check_restart` must be omitted entirely."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("require_healthy")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"check": {
			Description: lang.PlainText("Specifies a health check associated with the service. This can be specified multiple times to define multiple checks for the service. At this time, a check using the Nomad provider supports `tcp` and `http` checks. The Consul integration supports the `grpc`, `http`, `script`1, and `tcp` checks."),
			Body:        CheckSchema,
		},
		// TODO: add body
		"weights": {
			Description: lang.PlainText("Specifies how a service instance is weighted in a DNS SRV request based on the service's health status, as described in the Consul [weights](https://developer.hashicorp.com/consul/docs/services/configuration/services-configuration-reference#weights) documentation. Only available where `provider = \"consul\"` The `weight` block supports the following fields:\n\n- `passing`: 1 - The weight of services in passing state.\n- `warning`: 1 - The weight of services in warning state."),
		},
		"connect": {
			Description: lang.PlainText("Configures the [Consul service mesh](https://developer.hashicorp.com/nomad/docs/job-specification/connect) integration. Only available on group services and where `provider = \"consul\"`."),
			Body:        ConnectSchema,
		},
		"identity": {Description: lang.PlainText("Specifies a Workload Identity to use when obtaining Service Identity tokens from Consul to register the service. Only available where `provider = \"consul\"`. Typically this can be omitted so that Nomad will fall back to the server's [`consul.service_identity`](https://developer.hashicorp.com/nomad/docs/configuration/consul#service_identity) block."),
			Body: IdentitySchema,
		},
	},
}
