package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var AdvertiseSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"http": {
			Description: lang.Markdown("The address to advertise for the HTTP interface. This should be reachable by all the nodes from which end users are going to use the Nomad CLI tools."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"rpc": {
			Description: lang.Markdown("The address used to advertise to Nomad clients for connecting to Nomad servers for RPC. This allows Nomad clients to connect to Nomad servers from behind a NAT gateway. This address much be reachable by all Nomad client nodes. When set, the Nomad servers will use the `advertise.serf` address for RPC connections amongst themselves. Setting this value on a Nomad client has no effect."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"serf": {
			Description: lang.Markdown("The address advertised for the gossip layer. This address must be reachable from all server nodes. It is not required that clients can reach this address. Nomad servers will communicate to each other over RPC using the advertised Serf IP and advertised RPC Port."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
