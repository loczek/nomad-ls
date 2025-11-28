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
			Description:  lang.Markdown("Specifies the Consul cluster to use, when the `provider` is `consul`. The Nomad client will retrieve a Consul token from the cluster configured in the agent configuration with the same [`consul.name`](/nomad/docs/configuration/consul#name). In Nomad Community Edition, this field is ignored."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
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
			Description: lang.PlainText("Specifies how a service instance is weighted in a DNS SRV request based on the service's health status, as described in the Consul [weights](/consul/docs/services/configuration/services-configuration-reference#weights) documentation. Only available where `provider = \"consul\"` The `weight` block supports the following fields:\n\n- `passing`: 1 - The weight of services in passing state.\n- `warning`: 1 - The weight of services in warning state."),
		},
		"connect": {
			Description: lang.PlainText("Configures the [Consul service mesh](/nomad/docs/job-specification/connect) integration. Only available on group services and where `provider = \"consul\"`."),
			Body:        ConnectSchema,
		},
	},
}
