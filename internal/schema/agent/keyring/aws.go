package keyring

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var AWSSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"region": {
			Description:  lang.Markdown("The AWS region where the encryption key lives. If not provided, may be populated from the `AWS_REGION` or `AWS_DEFAULT_REGION` environment variables, from your `~/.aws/config` file, or from instance metadata."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("us-east-1")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"access_key": {
			Description: lang.Markdown("The AWS access key ID to use. Alternately specify via the `AWS_ACCESS_KEY_ID` environment variable or as part of the AWS profile from the AWS CLI or instance profile."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"session_token": {
			Description:  lang.Markdown("Specifies the AWS session token. Alternately specify via the environment variable `AWS_SESSION_TOKEN`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"secret_key": {
			Description: lang.Markdown("The AWS secret access key to use. Alternately specify via the `AWS_SECRET_ACCESS_KEY` environment variable or as part of the AWS profile from the AWS CLI or instance profile."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"kms_key_id": {
			Description: lang.Markdown("The AWS KMS key ID or ARN to use for encryption and decryption. You can alternately use an alias in the format `alias/key-alias-name`."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"endpoint": {
			Description:  lang.Markdown("The KMS API endpoint for AWS KMS requests. Alternately specify via the `AWS_KMS_ENDPOINT` environment variable. This is useful, for example, when connecting to KMS over a [VPC Endpoint](https://docs.aws.amazon.com/kms/latest/developerguide/kms-vpc-endpoint.html). If not set, Nomad uses the default API endpoint for your region."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
