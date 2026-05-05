package plugin

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var QEMUSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"image_paths": {
			Description:  lang.Markdown("Specifies the host paths the QEMU driver is allowed to load images from."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"args_allowlist": {
			Description:  lang.Markdown("Specifies the command line flags that the [`args`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/qemu#args) option is permitted to pass to QEMU. If unset, a job submitter can pass any command line flag into QEMU, including flags that provide the VM with access to host devices such as USB drives. Refer to the [QEMU documentation](https://www.qemu.org/docs/master/system/invocation.html) for the available flags."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"emulators_allowlist": {
			Description:  lang.Markdown("The list of emulator architectures allowed on this client. For example, `[\"x86_64\", \"aarch64\"]`. If unset, all emulators are allowed."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}
