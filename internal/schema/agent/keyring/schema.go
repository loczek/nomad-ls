package keyring

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var KeyringSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description:  lang.Markdown("A unique identifier for the keyring block, used to disambiguate when there are multiple blocks of the same type."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"active": {
			Description:  lang.Markdown("Indicates which block to use for encrypting keys. For existing servers, changing which block is active only impacts new keys created by a key rotation. Existing keys are encrypted with the previous active block, so those blocks should not be removed from the configuration until they have been garbage collected and the keys have been removed from the keystore. In Nomad Community Edition, only a single keyring can be active at a time."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
