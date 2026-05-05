package keyring

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var VaultSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"key_name": {
			Description: lang.Markdown("The transit key to use for encryption and decryption."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"key_id_prefix": {
			Description:  lang.Markdown("An optional string to add to the key id of values wrapped by this transit keyring. This can help disambiguate between two transit keyring."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"mount_path": {
			Description: lang.Markdown("The mount path to the transit secret engine."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"disable_renewal": {
			Description:  lang.Markdown("Disables the automatic renewal of the token in case the lifecycle of the token is managed with some other mechanism outside of Vault, such as Vault Agent."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("false")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: this can't be required
		"address": {
			Description: lang.Markdown("The full address to the Vault cluster. Alternately specify via the VAULT_ADDR environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: this can't be required
		"token": {
			Description: lang.Markdown("The Vault token to use. Alternately specify via the VAULT_TOKEN environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"namespace": {
			Description:  lang.Markdown("The namespace path to the transit secret engine. Alternately specify via the VAULT_NAMESPACE environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_ca_cert": {
			Description:  lang.Markdown("Specifies the path to the CA certificate file used for communication with the Vault server. Alternately specify via the VAULT_CACERT environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_client_cert": {
			Description:  lang.Markdown("Specifies the path to the client certificate for communication with the Vault server. Alternately specify via the VAULT_CLIENT_CERT environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_client_key": {
			Description:  lang.Markdown("Specifies the path to the private key for communication with the Vault server. Alternately specify via the VAULT_CLIENT_KEY environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_server_name": {
			Description:  lang.Markdown("Name to use as the SNI host when connecting to the Vault server via TLS. Alternately specify via the VAULT_TLS_SERVER_NAME environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_skip_verify": {
			Description:  lang.Markdown("Disable verification of TLS certificates. Using this option is highly discouraged and decreases the security of data transmissions to and from the Vault server. Alternately specify via the VAULT_SKIP_VERIFY environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("false")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
