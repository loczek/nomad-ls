package plugin

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var JavaSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"default_pid_mode": {
			Description:  lang.Markdown("Defaults to `\"private\"`. Set to `\"private\"` to enable PID namespace isolation for tasks by default, or `\"host\"` to disable isolation."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("private")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"default_ipc_mode": {
			Description:  lang.Markdown("Defaults to `\"private\"`. Set to `\"private\"` to enable IPC namespace isolation for tasks by default, or `\"host\"` to disable isolation."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("private")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"allow_caps": {
			Description: lang.Markdown("A list of allowed Linux capabilities. Defaults to\n\n```\n[\"audit_write\", \"chown\", \"dac_override\", \"fowner\", \"fsetid\", \"kill\", \"mknod\", \"net_bind_service\", \"setfcap\", \"setgid\", \"setpcap\", \"setuid\", \"sys_chroot\"]\n\nwhich is modeled after the capabilities allowed by [docker by default](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities) (without [`NET_RAW`](https://developer.hashicorp.com/nomad/docs/upgrade/upgrade-specific#nomad-1-1-0-rc1-1-0-5-0-12-12)). Allows the operator to control which capabilities can be obtained by tasks using [`cap_add`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/exec#cap_add) and [`cap_drop`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/exec#cap_drop) options. Supports the value `\"all\"` as a shortcut for allow-listing all capabilities supported by the operating system."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("audit_write"),
				cty.StringVal("chown"),
				cty.StringVal("dac_override"),
				cty.StringVal("fowner"),
				cty.StringVal("fsetid"),
				cty.StringVal("kill"),
				cty.StringVal("mknod"),
				cty.StringVal("net_bind_service"),
				cty.StringVal("setfcap"),
				cty.StringVal("setgid"),
				cty.StringVal("setpcap"),
				cty.StringVal("setuid"),
				cty.StringVal("sys_chroot"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
	},
}
