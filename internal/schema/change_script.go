package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ChangeScriptSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		// TODO: add examples
		"command": {
			Description:  lang.Markdown("Specifies the full path to a script or executable that is to be executed on template change. The command must return exit code 0 to be considered successful. Path is relative to the driver, e.g., if running with a container driver the path must be existing in the container. This option is required if `change_mode` is `script`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"args": {
			Description: lang.Markdown("List of arguments that are passed to the script that is to be executed on template change."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"timeout": {
			Description:  lang.Markdown("Timeout for script execution specified using a label suffix like `\"30s\"` or `\"1h\"`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"fail_on_error": {
			Description:  lang.Markdown("If `true`, Nomad will kill the task if the script execution fails. If `false`, script failure will be logged but the task will continue uninterrupted."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
