package schemaACL

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var SentinelDenyCapability = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("forbids any Sentinel API operation. Deny takes precedence over all other capabilities associated with a token."),
}
var SentinelReadCapability = schema.LiteralValue{
	Value:       cty.StringVal("sentinel-read"),
	Description: lang.Markdown("allows reading or listing Sentinel policies."),
}
var SentinelSubmitCapability = schema.LiteralValue{
	Value:       cty.StringVal("sentinel-submit"),
	Description: lang.Markdown("allows submitting new Sentinel policies or updating policies."),
}
var SentinelDeleteCapability = schema.LiteralValue{
	Value:       cty.StringVal("sentinel-delete"),
	Description: lang.Markdown("allows deleting Sentinel policies."),
}

var SentinelSchema = &schema.BodySchema{
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
				SentinelDenyCapability,
				SentinelReadCapability,
				SentinelSubmitCapability,
				SentinelDeleteCapability,
			},
			IsRequired: true,
		},
	},
}
