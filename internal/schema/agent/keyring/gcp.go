package keyring

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var GCPSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"credentials": {
			Description: lang.Markdown("The path to the credentials JSON file to use. Alternately specify via the `GOOGLE_CREDENTIALS` or `GOOGLE_APPLICATION_CREDENTIALS` environment variable or set automatically if running under Google Compute Engine."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"project": {
			Description: lang.Markdown("The GCP project ID to use. Alternately specify via the `GOOGLE_PROJECT` environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"region": {
			Description: lang.Markdown("The GCP region/location where the key ring lives. Alternately specify via the `GOOGLE_REGION` environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"key_ring": {
			Description: lang.Markdown("The GCP KMS key ring to use."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"crypto_key": {
			Description: lang.Markdown("The GCP KMS crypto key to use for encryption and decryption."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
	},
}
