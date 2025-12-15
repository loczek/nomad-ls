package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// MetaSchema defines the schema for a `meta` block.
// This block allows specifying arbitrary key-value metadata
// that can be used for job organization, templating, or
// passing information to tasks.
//
// Example:
//
//	meta {
//	  environment = "production"
//	  team        = "platform"
//	}
var MetaSchema = &schema.BodySchema{
	Description: lang.Markdown("Specifies a key-value map that annotates with user-defined metadata."),
	AnyAttribute: &schema.AttributeSchema{
		Description: lang.Markdown("A user-defined key-value pair for metadata."),
		Constraint:  &schema.LiteralType{Type: cty.String},
		IsOptional:  true,
	},
}
