package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var AddressesSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"http": {
			Description: lang.Markdown("The address the HTTP server is bound to. This is the most common bind address to change. The `http` field accepts multiple values, separated by spaces, to bind to multiple addresses."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"rpc": {
			Description: lang.Markdown("The address to bind the internal RPC interfaces to. Should be exposed only to other cluster members if possible."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"serf": {
			Description: lang.Markdown("The address used to bind the gossip layer to. Both a TCP and UDP listener will be exposed on this address. Should be exposed only to other cluster members if possible."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
