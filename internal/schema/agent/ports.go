package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var PortsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"http": {
			Description:  lang.Markdown("The port used to run the HTTP server."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(4646)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"rpc": {
			Description:  lang.Markdown("The port used for internal RPC communication between agents and servers, and for inter-server traffic for the consensus algorithm (raft)."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(4647)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"serf": {
			Description:  lang.Markdown("The port used for the gossip protocol for cluster membership. Both TCP and UDP should be routable between the server nodes on this port."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(4648)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
	},
}
