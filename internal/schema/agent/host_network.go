package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var HostNetworkSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cidr": {
			Description:  lang.Markdown("Specifies a cidr block of addresses to match against. If an address is found on the node that is contained by this cidr block, the host network will be registered with it."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"interface": {
			Description:  lang.Markdown("Filters searching of addresses to a specific interface."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"reserved_ports": {
			Description:  lang.Markdown("Specifies a comma-separated list of ports to reserve on all addresses associated with this network. Ranges can be specified by using a hyphen separating the two inclusive ends. [`reserved.reserved_ports`](https://developer.hashicorp.com/nomad/docs/configuration/client#reserved_ports) are also reserved on each host network."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
