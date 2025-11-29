package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var SecretSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"provider": {
			Description:  lang.Markdown("Specifies the underlying implementation to use in order for Nomad to interact with a specific secret store."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"path": {
			Description:  lang.Markdown("Specifies the location of the secret within the given secret store."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"env": {
			Description: lang.Markdown("Specifies any environment variables to pass to a custom plugin provider. Only available for custom providers."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"config": {
			Description: lang.Markdown("Specifies any custom attributes used by built-in providers order to fetch the secret. Only available for built-in `vault` and `nomad` providers."),
			Type:        schema.BlockTypeMap,
		},
	},
}
