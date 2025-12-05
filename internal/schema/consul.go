package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ConsulSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cluster": {
			Description: lang.Markdown("Specifies the Consul cluster to use. The Nomad client will retrieve a Consul token from the cluster configured in the agent configuration with the same [`consul.name`](https://developer.hashicorp.com/nomad/docs/configuration/consul#name). In Nomad Community Edition, this field is ignored."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("default"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"namespace": {
			Description: lang.Markdown("The Consul namespace in which group and task-level services within the group will be registered. Use of `template` to access Consul KV will read from the specified Consul namespace. Specifying `namespace` takes precedence over the [`-consul-namespace`](https://developer.hashicorp.com/nomad/commands/job/run#consul-namespace) command line argument in `job run`. In Nomad Community Edition, this field is ignored."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"partition": {
			Description: lang.Markdown("When this field is set, a constraint will be added to the group or task to ensure that the allocation is placed on a Nomad client that has a Consul Enterprise agent in the specified Consul [admin partition](https://developer.hashicorp.com/consul/docs/enterprise/admin-partitions). Using this field requires the following:"),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
	},
}
