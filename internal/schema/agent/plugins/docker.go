package plugin

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var DockerSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"endpoint": {
			Description:  lang.Markdown("If using a non-standard socket, HTTP or another location, or if TLS is being used, docker.endpoint must be set. If unset, Nomad will attempt to instantiate a Docker client using the `DOCKER_HOST` environment variable and then fall back to the default listen address for the given operating system. Defaults to `unix:///var/run/docker.sock` on Unix platforms and `npipe:////./pipe/docker_engine` for Windows."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("unix:///var/run/docker.sock")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"allow_privileged": {
			Description:  lang.Markdown("Defaults to `false`. Changing this to true will allow containers to use privileged mode, which gives the containers full access to the host's devices. Note that you must set a similar setting on the Docker daemon for this to work."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"pull_activity_timeout": {
			Description:  lang.Markdown("Defaults to `2m`. If Nomad receives no communication from the Docker engine during an image pull within this timeframe, Nomad will time out the request that initiated the pull command. (Minimum of `1m`)"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("2m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"pids_limit": {
			Description:  lang.Markdown("Defaults to unlimited (`0`). An integer value that specifies the pid limit for all the Docker containers running on that Nomad client. You can override this limit by setting [`pids_limit`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/docker#pids_limit) in your task config. If this value is greater than `0`, your task `pids_limit` must be less than or equal to the value defined here."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		// TODO: add warning
		"allow_caps": {
			Description: lang.Markdown("A list of allowed Linux capabilities. Defaults to\n\n```\n[\"audit_write\", \"chown\", \"dac_override\", \"fowner\", \"fsetid\", \"kill\", \"mknod\", \"net_bind_service\", \"setfcap\", \"setgid\", \"setpcap\", \"setuid\", \"sys_chroot\"]\n```\n\nwhich is the same list of capabilities allowed by [docker by default](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities) (without [`NET_RAW`](https://developer.hashicorp.com/nomad/docs/upgrade/upgrade-specific#nomad-1-1-0-rc1-1-0-5-0-12-12)). Allows the operator to control which capabilities can be obtained by tasks using [`cap_add`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/docker#cap_add) and [`cap_drop`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/docker#cap_drop) options. Supports the value `\"all\"` as a shortcut for allow-listing all capabilities supported by the operating system. Note that due to a limitation in Docker, tasks running as non-root users cannot expand the capabilities set beyond the default. They can only have their capabilities reduced."),
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
		"allow_runtimes": {
			Description: lang.Markdown("A list of the allowed docker runtimes a task may use."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("runc"),
				cty.StringVal("nvidia"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"extra_labels": {
			Description:  lang.Markdown("Extra labels to add to Docker containers. Available options are `job_name`, `job_id`, `task_group_name`, `task_name`, `namespace`, `node_name`, `node_id`. Globs are supported (e.g. `task*`)"),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"disable_log_collection": {
			Description:  lang.Markdown("Setting this to true will disable Nomad logs collection of Docker tasks. If you don't rely on nomad log capabilities and exclusively use host based log aggregation, you may consider this option to disable nomad log collection overhead."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"infra_image": {
			Description:  lang.Markdown("This is the Docker image to use when creating the parent container necessary when sharing network namespaces between tasks. Defaults to `registry.k8s.io/pause-<goarch>:3.3`. The image will only be pulled from the container registry if its tag is `latest` or the image doesn't yet exist locally."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("registry.k8s.io/pause-<goarch>:3.3")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"infra_image_pull_timeout": {
			Description:  lang.Markdown("A time duration that controls how long Nomad will wait before cancelling an in-progress pull of the Docker image as specified in `infra_image`. Defaults to `\"5m\"`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"image_pull_timeout": {
			Description:  lang.Markdown("A default time duration that controls how long Nomad waits before cancelling an in-progress pull of the Docker image as specified in `image` across all tasks. Defaults to `\"5m\"`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"windows_allow_insecure_container_admin": {
			Description:  lang.Markdown("Indicates that on windows, docker checks the `task.user` field or, if unset, the container image manifest after pulling the container, to see if it's running as `ContainerAdmin`. If so, exits with an error unless the task config has `privileged=true`. Defaults to `false`."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"auth": {
			Body: AuthSchema,
		},
		"tls": {
			Body: TLSSchema,
		},
		"logging": {
			Body: LoggingSchema,
		},
		"gc": {
			Body: GCSchema,
		},
		"volumes": {
			Body: VolumesSchema,
		},
	},
}

var AuthSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"config": {
			Description: lang.Markdown("Allows an operator to specify a JSON file which is in the dockercfg format containing authentication information for a private registry, from either (in order) auths, credsStore or credHelpers."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"helper": {
			Description: lang.Markdown("Allows an operator to specify a [credsStore](https://docs.docker.com/engine/reference/commandline/login/#credential-helper-protocol) like script on `$PATH` to lookup authentication information from external sources. The script's name must begin with `docker-credential-` and this option should include only the basename of the script, not the path.\n\nIf you set an auth helper, it will be tried for all images, including public images. If you mix private and public images, you will need to include [`auth_soft_fail=true`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/docker#auth_soft_fail) in every job using a public image."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}

var TLSSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cert": {
			Description: lang.Markdown("Path to the server's certificate file (`.pem`). Specify this along with `key` and `ca` to use a TLS client to connect to the docker daemon. `endpoint` must also be specified or this setting will be ignored."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"key": {
			Description: lang.Markdown("Path to the client's private key (`.pem`). Specify this along with `cert` and `ca` to use a TLS client to connect to the docker daemon. `endpoint` must also be specified or this setting will be ignored."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"ca": {
			Description: lang.Markdown("Path to the server's CA file (`.pem`). Specify this along with `cert` and `key` to use a TLS client to connect to the docker daemon. `endpoint` must also be specified or this setting will be ignored."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}

// TODO: this is a duplicate schema
var LoggingSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description:  lang.Markdown("Defaults to `\"json-file\"`. Specifies the logging driver docker should use for all containers Nomad starts. Note that for older versions of Docker, only `json-file` file or `journald` will allow Nomad to read the driver's logs via the Docker API, and this will prevent commands such as `nomad alloc logs` from functioning."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("json-file")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"config": {
			Description: lang.Markdown("Defaults to `{ max-file = \"2\", max-size = \"2m\" }`. This option can also be used to pass further [configuration](https://docs.docker.com/config/containers/logging/configure/) to the logging driver."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Description: lang.Markdown("Logging driver configuration option."),
					Constraint:  schema.LiteralType{Type: cty.String},
					IsOptional:  true,
				},
			},
		},
	},
}

var GCSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"image": {
			Description:  lang.Markdown("Defaults to `true`. Changing this to `false` will prevent Nomad from removing images from stopped tasks."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"image_delay": {
			Description:  lang.Markdown("A time duration, as [defined here](https://golang.org/pkg/time/#ParseDuration), that defaults to `3m`. The delay controls how long Nomad will wait between an image being unused and deleting it. If a task is received that uses the same image within the delay, the image will be reused. If an image is referenced by more than one tag, `image_delay` may not work correctly."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("3m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"container": {
			Description:  lang.Markdown("Defaults to `true`. This option can be used to disable Nomad from removing a container when the task exits. Under a name conflict, Nomad may still remove the dead container."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"dangling_containers": {
			Body: DanglingContainersSchema,
		},
	},
}

var DanglingContainersSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Defaults to `true`. Enables dangling container handling."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"dry_run": {
			Description:  lang.Markdown("Defaults to `false`. Only log dangling containers without cleaning them up."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"period": {
			Description:  lang.Markdown("Defaults to `\"5m\"`. A time duration that controls interval between Nomad scans for dangling containers."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"creation_grace": {
			Description:  lang.Markdown("Defaults to `\"5m\"`. Grace period after a container is created during which the GC ignores it. Only used to prevent the GC from removing newly created containers before they are registered with the GC. Should not need adjusting higher but may be adjusted lower to GC more aggressively."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}

var VolumesSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Defaults to `false`. Allows tasks to bind host paths (`volumes`) inside their container and use volume drivers (`volume_driver`). Binding relative paths is always allowed and will be resolved relative to the allocation's directory."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"selinuxlabel": {
			Description: lang.Markdown("Allows the operator to set a SELinux label to the allocation and task local bind-mounts to containers. If used with `docker.volumes.enabled` set to false, the labels will still be applied to the standard binds in the container."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
