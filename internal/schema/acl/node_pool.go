package schemaACL

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var NodePoolDenyCapability = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("forbids node pools to be read or modified. Deny takes precedence when multiple policies are associated with a token."),
}
var NodePoolDeleteCapability = schema.LiteralValue{
	Value:       cty.StringVal("delete"),
	Description: lang.Markdown("allows node pools to be deleted."),
}
var NodePoolReadCapability = schema.LiteralValue{
	Value:       cty.StringVal("read"),
	Description: lang.Markdown("allows node pools to be listed and read."),
}
var NodePoolWriteCapability = schema.LiteralValue{
	Value:       cty.StringVal("write"),
	Description: lang.Markdown("allows node pools to be created and updated."),
}

var NodePoolSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"policy": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				ReadPolicy,
				WritePolicy,
				DenyPolicy,
			},
			IsRequired: true,
		},
		"capabilities": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				NodePoolDenyCapability,
				NodePoolDeleteCapability,
				NodePoolReadCapability,
				NodePoolWriteCapability,
			},
			IsRequired: true,
		},
	},
}
