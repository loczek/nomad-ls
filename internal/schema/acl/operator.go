package schemaACL

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var OperatorDenyCapability = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("forbids any Operator API operation. Deny takes precedence over all other capabilities associated with a token."),
}
var OperatorKeyringRotateCapability = schema.LiteralValue{
	Value:       cty.StringVal("keyring-rotate"),
	Description: lang.Markdown("allows rotating the root keyring."),
}
var OperatorKeyringDeleteCapability = schema.LiteralValue{
	Value:       cty.StringVal("keyring-delete"),
	Description: lang.Markdown("allows deleting keys from the root keyring."),
}
var OperatorKeyringReadCapability = schema.LiteralValue{
	Value:       cty.StringVal("keyring-read"),
	Description: lang.Markdown("allows reading key metadata from the root keyring."),
}
var OperatorLicenseReadCapability = schema.LiteralValue{
	Value:       cty.StringVal("license-read"),
	Description: lang.Markdown("allows reading the license of a Nomad Enterprise server."),
}
var OperatorSnapshotSaveCapability = schema.LiteralValue{
	Value:       cty.StringVal("snapshot-save"),
	Description: lang.Markdown("allows saving Raft snapshots."),
}

var OperatorSchema = &schema.BodySchema{
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
				OperatorDenyCapability,
				OperatorKeyringRotateCapability,
				OperatorKeyringDeleteCapability,
				OperatorKeyringReadCapability,
				OperatorLicenseReadCapability,
				OperatorSnapshotSaveCapability,
			},
			IsRequired: true,
		},
	},
}
