package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var SentinelSchema = &schema.BodySchema{
	Description: lang.Markdown("Specifies configuration for Sentinel policies." + schemautils.Divider + schemautils.EnterpriseOnly),
	Attributes: map[string]*schema.AttributeSchema{
		"additional_enabled_modules": {
			Description:  lang.Markdown("Specifies a list of additional standard imports (modules) to allow in policies. Nomad currently enables all of Sentinel's standard imports except the \"http\" import, which has performance and security implications. Setting this field to [\"http\"] enables the \"http\" module in addition to the standard imports. In the future, if any new Sentinel imports are not automatically enabled by nomad, you can enable them in this field. Refer to Using the http import in Sentinel policies for recommendations on safe use of this import."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"import": {
			Description: lang.Markdown("Specifies a plugin that should be made available for importing by Sentinel policies. The name of the import matches the name that can be imported."),
			Body:        ImportSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
	},
}
var ImportSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"path": {
			Description:  lang.Markdown("Specifies the path to the import plugin. Must be executable by Nomad."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"args": {
			Description:  lang.Markdown("Specifies arguments to pass to the plugin when starting it."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}
