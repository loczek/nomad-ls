package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var EphemeralDiskSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"migrate": {
			Description: lang.Markdown("This specifies that the Nomad client should make a best-effort attempt to migrate the data from the previous allocation, even if the previous allocation was on another client. Enabling `migrate` automatically enables `sticky` as well. During data migration, the task will block starting until the data migration has completed.\n\nSuccessful migration requires that the clients can reach each other directly over the Nomad HTTP port. Any failure of the transfer will result in data loss, so this feature is only suitable for data that can be recreated at the destination (for example, cache data). Migration is atomic and any partially migrated data will be removed from the destination if an error is encountered. Note that data migration will not take place if a client garbage collects a failed allocation or if the allocation has been intentionally stopped via `nomad alloc stop`, because the original allocation has already been removed."),
			DefaultValue: schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
		"size": {
			Description: lang.Markdown("Specifies the size of the ephemeral disk in MB. The current Nomad ephemeral storage implementation does not enforce this limit; however, it is used during job placement."),
			DefaultValue: schema.DefaultValue{
				Value: cty.NumberIntVal(300),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
		"sticky": {
			Description: lang.Markdown("Specifies that Nomad should make a best-effort attempt to place the updated allocation on the same machine. This will move the `local/` and `alloc/data` directories to the new allocation."),
			DefaultValue: schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
	},
}
