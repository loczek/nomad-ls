package drivers

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var QemuDriverSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"image_path": {
			Description: lang.Markdown("The path to the downloaded image. In most cases this will just be the name of the image. However, if the supplied artifact is an archive that contains the image in a subfolder, the path will need to be the relative path (`subdir/from_archive/my.img`)."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.List(cty.String)},
				schema.AnyExpression{OfType: cty.List(cty.String)},
			},
			IsRequired: true,
		},
		"emulator": {
			Description:  lang.Markdown("The QEMU emulator arch used to run the virtual machine. This also supports the full emulator binary name. For example, `qemu-system-x86_64`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("x86_64")},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"machine_type": {
			Description:  lang.Markdown("Used to select the emulated machine by name"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("pc")},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"drive_interface": {
			Description:  lang.Markdown("This option defines on which type of interface the drive is connected. Available types are: `ide`, `scsi`, `sd`, `mtd`, `floppy`, `pflash`, `virtio` and `none`. Default is `ide`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("ide")},
			Constraint: schema.OneOf{
				schema.LiteralValue{Value: cty.StringVal("ide")},
				schema.LiteralValue{Value: cty.StringVal("scsi")},
				schema.LiteralValue{Value: cty.StringVal("sd")},
				schema.LiteralValue{Value: cty.StringVal("mtd")},
				schema.LiteralValue{Value: cty.StringVal("floppy")},
				schema.LiteralValue{Value: cty.StringVal("pflash")},
				schema.LiteralValue{Value: cty.StringVal("virtio")},
				schema.LiteralValue{Value: cty.StringVal("none")},
			},
			IsOptional: true,
		},
		"accelerator": {
			Description:  lang.Markdown("The type of accelerator to use in the invocation. If the host machine has `qemu` installed with KVM support, users can specify `kvm` for the `accelerator`. Default is `tcg`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("tcg")},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"graceful_shutdown": {
			Description:  lang.Markdown("Using the [qemu monitor](https://en.wikibooks.org/wiki/QEMU/Monitor), send an ACPI shutdown signal to virtual machines rather than simply terminating them. This emulates a physical power button press, and gives instances a chance to shut down cleanly. If the VM is still running after `kill_timeout`, it will be forcefully terminated. This feature uses a Unix socket that is placed within the task directory and operating systems may impose a limit on how long these paths can be. This feature is currently not supported on Windows."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
		"guest_agent": {
			Description:  lang.Markdown("Enable support for the [QEMU Guest Agent](https://wiki.qemu.org/Features/GuestAgent) for this virtual machine. This will add the necessary virtual hardware and create a `qa.sock` file in the task's working directory for interacting with the agent. The QEMU Guest Agent must be running in the guest VM. This feature is currently not supported on Windows."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
		// TODO: add example
		"port_map": {
			Description: lang.Markdown("A key-value map of port labels."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Map(cty.String)},
				schema.AnyExpression{OfType: cty.Map(cty.String)},
			},
			IsOptional: true,
		},
		"args": {
			Description: lang.Markdown("A list of strings that is passed to QEMU as command line options."),
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.List(cty.String)},
				schema.AnyExpression{OfType: cty.List(cty.String)},
			},
			IsOptional: true,
		},
	},
}
