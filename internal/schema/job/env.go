package job

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var EnvSchema = &schema.BodySchema{
	Description: lang.Markdown("`The `env` block configures a list of environment variables to populate the task's environment before starting."),
	AnyAttribute: &schema.AttributeSchema{
		Description: lang.Markdown("The parameters for the `env` block can be any key-value. The keys and values are both of type `string`, but they can be specified as other types. They will automatically be converted to strings. Invalid characters such as dashes (`-`) will be converted to underscores."),
		Constraint: schema.OneOf{
			schema.LiteralType{Type: cty.String},
			schema.AnyExpression{OfType: cty.String},
		},
		IsOptional: true,
	},
}
