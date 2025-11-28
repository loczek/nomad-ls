package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var VolumeMountSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"volume": {
			Description:  lang.Markdown("Specifies the group volume that the mount is going to access."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
		},
		"destination": {
			Description:  lang.Markdown("Specifies where the volume should be mounted inside the task's allocation."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"read_only": {
			Description:  lang.Markdown("When a group volume is writeable, you may specify that it is `read_only` on a per mount level using the `read_only` option here."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		// TODO: add warning from docs
		"propagation_mode": {
			Description:  lang.Markdown("Specifies the mount propagation mode for nested volumes. Possible values are:\n\n- [`private`](/nomad/docs/job-specification/volume_mount#private) - the task is not allowed to access nested mounts.\n- [`host-to-task`](/nomad/docs/job-specification/volume_mount#host-to-task) - allows new mounts that have been created outside of the task to be visible inside the task.\n- [`bidirectional`](/nomad/docs/job-specification/volume_mount#bidirectional) - allows the task to both access new mounts from the host and also create new mounts. This mode requires `ReadWrite` permission."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("private")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"selinux_label": {
			Description:  lang.Markdown("Specifies the SELinux label for the mount. This is only supported on Linux hosts and when supported by the task driver. Refer to the task driver documentation for more information. Possible values are:\n\n- [`Z`](/nomad/docs/job-specification/volume_mount#z) - Specifies that the volume content is private and unshared between containers.\n- [`z`](/nomad/docs/job-specification/volume_mount#z-1) - Specifies that the volume content is shared among containers."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
