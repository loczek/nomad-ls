package volume

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var MountOptionsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"fs_type": &schema.AttributeSchema{
			Description: lang.Markdown("File system type (ex. `\"ext4\"`)"),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"mount_flags": {
			Description:  lang.Markdown("The flags passed to `mount` (ex. `[\"ro\", \"noatime\"]`)"),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}
