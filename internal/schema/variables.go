package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// VariablesSchema defines the schema for a `variables` block.
// This block provides a less precise but more concise way to declare
// multiple input variables at once. Unlike the `variable` block,
// it does not support type constraints, defaults, descriptions, or validation.
// Variables declared this way are implicitly strings unless the value
// can be parsed as another type.
//
// Example:
//
//	variables {
//	  foo       = "value"
//	  my_secret = "foo"
//	}
var VariablesSchema = &schema.BodySchema{
	Description: lang.Markdown("A less precise block for declaring multiple input variables at once. Each attribute defines a variable name and its default value. Variables declared in a `variables` block do not support type constraints, descriptions, or validation rules. For more control over variable behavior, use individual `variable` blocks instead."),
	AnyAttribute: &schema.AttributeSchema{
		Description: lang.Markdown("Defines a variable with the given name and default value."),
		Constraint:  &schema.LiteralType{Type: cty.DynamicPseudoType},
		IsOptional:  true,
	},
}
