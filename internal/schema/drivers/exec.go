package drivers

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ExecDriverSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"args": {
			Description: lang.Markdown("A list of arguments to the `command`. References to environment variables or any [interpretable Nomad variables](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation) will be interpreted before launching the task."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"command": {
			Description: lang.Markdown("The command to execute. Must be provided. If executing a binary that exists on the host, the path must be absolute and within the task's [chroot](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#chroot) or in a [host volume](https://developer.hashicorp.com/nomad/docs/configuration/client#host_volume-block) mounted with a [`volume_mount`](https://developer.hashicorp.com/nomad/docs/job-specification/volume_mount) block. The driver will make the binary executable and will search, in order:\n\n- The `local` directory with the task directory.\n- The task directory.\n- Any mounts, in the order listed in the job specification.\n- The `usr/local/bin`, `usr/bin` and `bin` directories inside the task directory.\n\nIf executing a binary that is downloaded from an [`artifact`](https://developer.hashicorp.com/nomad/docs/job-specification/artifact), the path can be relative from the allocation's root directory."),
			Constraint:  &schema.LiteralType{Type: cty.String},
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
