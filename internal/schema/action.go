package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ActionSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"command": {
			Description: lang.Markdown("Specifies the command to be executed."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"args": {
			Description: lang.Markdown("Provides a list of arguments to pass to the command."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
		},
	},
}
