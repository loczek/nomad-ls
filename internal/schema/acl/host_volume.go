package schemaACL

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var HostVolumeDenyCapability = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("Do not allow a user to mount a volume in any way."),
}
var HostVolumeMountReadOnlyCapability = schema.LiteralValue{
	Value:       cty.StringVal("mount-readonly"),
	Description: lang.Markdown("Only allow the user to mount the volume as `readonly`"),
}
var HostVolumeMountReadWriteCapability = schema.LiteralValue{
	Value:       cty.StringVal("mount-readwrite"),
	Description: lang.Markdown("Allow the user to mount the volume as `readonly` or `readwrite` if the `host_volume` configuration allows it."),
}

var HostVolumeSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"policy": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				ReadPolicy,
				WritePolicy,
				DenyPolicy,
			},
			IsRequired: true,
		},
		"capabilities": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				HostVolumeDenyCapability,
				HostVolumeMountReadOnlyCapability,
				HostVolumeMountReadWriteCapability,
			},
			IsRequired: true,
		},
	},
}
