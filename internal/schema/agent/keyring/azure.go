package keyring

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var AzureSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"tenant_id": {
			Description: lang.Markdown("The tenant id for the Azure Active Directory organization. Alternately specify via the `AZURE_TENANT_ID` environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"client_id": {
			Description: lang.Markdown("The client id for credentials to query the Azure APIs. Alternately specify via the `AZURE_CLIENT_ID` environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"client_secret": {
			Description: lang.Markdown("The client secret for credentials to query the Azure APIs. Alternately specify via the `AZURE_CLIENT_SECRET` environment variable."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"environment": {
			Description:  lang.Markdown("The Azure Cloud environment API endpoints to use. Alternately specify via the `AZURE_ENVIRONMENT` environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("AZUREPUBLICCLOUD")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"vault_name": {
			Description: lang.Markdown("The Key Vault vault to use the encryption keys for encryption and decryption."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"key_name": {
			Description: lang.Markdown("The Key Vault key to use for encryption and decryption."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"resource": {
			Description:  lang.Markdown("The Key Vault resource's DNS Suffix to connect to. Alternately specify via the `AZURE_AD_RESOURCE` environment variable. Needs to be changed to connect to Azure's Managed HSM KeyVault instance type."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("vault.azure.net")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
