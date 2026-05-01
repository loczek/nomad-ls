package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var UsersSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"dynamic_user_min": {
			Description:  lang.Markdown("The lowest UID/GID to allocate for task drivers capable of making use of dynamic workload users."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(80000)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"dynamic_user_max": {
			Description:  lang.Markdown("The highest UID/GID to allocate for task drivers capable of making use of dynamic workload users."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(89999)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
	},
}
