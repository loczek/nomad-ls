package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var HostVolumeSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"path": {
			Description:  lang.Markdown("Specifies the path on the host that should be used as the source when this volume is mounted into a task. The path must exist on client startup."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"read_only": {
			Description:  lang.Markdown("Specifies whether the volume should only ever be allowed to be mounted `read_only`, or if it should be writeable."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
