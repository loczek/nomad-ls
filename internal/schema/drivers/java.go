package drivers

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var JavaDriverSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"class": {
			Description: lang.Markdown("The name of the class to run. If `jar_path` is specified and the manifest specifies a main class, this is optional. If shipping classes rather than a Jar, please specify the class to run and the `class_path`."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"class_path": {
			Description: lang.Markdown("The `class_path` specifies the class path used by Java to lookup classes and Jars."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"jar_path": {
			Description: lang.Markdown("The path to the downloaded Jar. In most cases this will just be the name of the Jar. However, if the supplied artifact is an archive that contains the Jar in a subfolder, the path will need to be the relative path (`subdir/from_archive/my.jar`)."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"args": {
			Description: lang.Markdown("A list of arguments to the Jar's main method. References to environment variables or any [interpretable Nomad variables](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation) will be interpreted before launching the task."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"jvm_options": {
			Description: lang.Markdown("A list of JVM options to be passed while invoking java. These options are passed without being validated in any way by Nomad."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		// TODO: add warning from docs
		"pid_mode": {
			Description: lang.Markdown("Set to `private` to enable PID namespace isolation for this task, or `host` to disable isolation. If left unset, the behavior is determined from the [`default_pid_mode`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#default_pid_mode) in plugin configuration."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add warning from docs
		"ipc_mode": {
			Description: lang.Markdown("Set to `private` to enable IPC namespace isolation for this task, or `host` to disable isolation. If left unset, the behavior is determined from the [`default_ipc_mode`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#default_ipc_mode) in plugin configuration."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"cap_add": {
			Description: lang.Markdown("A list of Linux capabilities to enable for the task. Effective capabilities (computed from `cap_add` and `cap_drop`) must be a subset of the allowed capabilities configured with [`allow_caps`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#allow_caps). Note that `all` is not permitted here if the `allow_caps` field in the driver configuration doesn't also allow all capabilities."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		// TODO: add example
		"cap_drop": {
			Description: lang.Markdown("A list of Linux capabilities to disable for the task. Effective capabilities (computed from `cap_add` and `cap_drop`) must be a subset of the allowed capabilities configured with [`allow_caps`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#allow_caps)."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"work_dir": {
			Description: lang.Markdown("Sets a custom working directory for the task. This path must be absolute and within the task's [chroot](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#chroot) or in a [host volume](https://developer.hashicorp.com/nomad/docs/configuration/client#host_volume-block) mounted with a [`volume_mount`](https://developer.hashicorp.com/nomad/docs/job-specification/volume_mount) block. This will also change the working directory when using `nomad alloc exec`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
