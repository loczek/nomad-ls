package drivers

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var DockerDriverSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"image": {
			Description:  lang.Markdown("The Docker image to run. The image may include a tag or custom URL and should include `https://` if required. By default it will be fetched from Docker Hub. If the tag is omitted or equal to `latest` the driver will always try to pull the image. If the image to be pulled exists in a registry that requires authentication credentials must be provided to Nomad."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"image_pull_timeout": {
			Description:  lang.Markdown("A time duration that controls how long Nomad will wait before cancelling an in-progress pull of the Docker image as specified in `image`. Defaults to `\"5m\"`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"args": {
			Description: lang.Markdown("A list of arguments to the optional `command`. If no `command` is specified, the arguments are passed directly to the container. References to environment variables or any [interpretable Nomad variables](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation) will be interpreted before launching the task. For example:"),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		// TODO: update with docs
		"auth": {
			Description: lang.Markdown("A list of arguments to the optional `command`. If no `command` is specified, the arguments are passed directly to the container. References to environment variables or any [interpretable Nomad variables](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation) will be interpreted before launching the task. For example:"),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"auth_soft_fail": {
			Description:  lang.Markdown("Don't fail the task on an auth failure. Attempt to continue without auth. If the Nomad client configuration has an [`auth.helper`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/docker#helper) block, the helper will be tried for all images, including public images. If you mix private and public images, you will need to include `auth_soft_fail=true` in every job using a public image."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
		},
		"command": {
			Description: lang.Markdown("The command to run when starting the container."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"cgroupns": {
			Description: lang.Markdown("Cgroup namespace to use. Set to `host` or `private`. If not specified, the driver uses Docker's default. Refer to Docker's [dockerd reference](https://docs.docker.com/reference/cli/dockerd/) for more information."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"container_exists_attempts": {
			Description: lang.Markdown("A number of attempts to be made to purge a container if during task creation Nomad encounters an existing one in non-running state for the same task. Defaults to `5`."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"dns_search_domains": {
			Description: lang.Markdown("A list of DNS search domains for the container to use. If you are using bridge networking mode with a `network` block in the task group, you must set all DNS options in the `network.dns` block instead."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"dns_options": {
			Description: lang.Markdown("A list of DNS options for the container to use. If you are using bridge networking mode with a `network` block in the task group, you must set all DNS options in the `network.dns` block instead."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"dns_servers": {
			Description: lang.Markdown("A list of DNS servers for the container to use (e.g. [\"8.8.8.8\", \"8.8.4.4\"]). Requires Docker v1.10 or greater. If you are using bridge networking mode with a `network` block in the task group, you must set all DNS options in the `network.dns` block instead."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"entrypoint": {
			Description: lang.Markdown("A string list overriding the image's entrypoint."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"extra_hosts": {
			Description: lang.Markdown("A list of hosts, given as host:IP, to be added to `/etc/hosts`. This option may not work as expected in `bridge` network mode when there is more than one task within the same group. Refer to the [upgrade guide](https://developer.hashicorp.com/nomad/docs/upgrade/upgrade-specific#docker-driver) for more information."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"force_pull": {
			Description:  lang.Markdown("`true` or `false` (default). Always pull most recent image instead of using existing local image. Should be set to `true` if repository tags are mutable. If image's tag is `latest` or omitted, the image will always be pulled regardless of this setting."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"group_add": {
			Description: lang.Markdown("A list of supplementary groups to be applied to the container user."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"healthchecks": {
			Description: lang.Markdown("A configuration block for controlling how the docker driver manages HEALTHCHECK directives built into the container. Set `healthchecks.disable` to disable any built-in healthcheck."),
			Constraint:  &schema.Map{Name: "disable", Elem: schema.LiteralType{Type: cty.Bool}},
			IsOptional:  true,
		},
		"hostname": {
			Description: lang.Markdown("The hostname to assign to the container. When launching more than one of a task (using `count`) with this option set, every container the task starts will have the same hostname."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"init": {
			Description:  lang.Markdown("`true` or `false` (default). Enable init (tini) system when launching your container. When enabled, an init process will be used as the PID1 in the container. Specifying an init process ensures the usual responsibilities of an init system, such as reaping zombie processes, are performed inside the created container.\n\nThe default init process used is the first `docker-init` executable found in the system path of the Docker daemon process. This `docker-init` binary, included in the default installation, is backed by [tini](https://github.com/krallin/tini)."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"interactive": {
			Description:  lang.Markdown("`true` or `false` (default). Keep STDIN open on the container."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"isolation": {
			Description:  lang.Markdown("Specifies [Windows isolation](https://learn.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/hyperv-container) mode: `hyperv` or `process`. Defaults to `hyperv`."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("hyperv")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add proper constraints, defaults and update docs with examples
		"sysctl": {
			Description: lang.Markdown("A key-value map of sysctl configurations to set to the containers on start."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		// TODO: update docs with examples
		"ulimit": {
			Description: lang.Markdown("A key-value map of ulimit configurations to set to the containers on start."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"privileged": {
			Description:  lang.Markdown("`true` or `false` (default). Privileged mode gives the container access to devices on the host. Note that this also requires the nomad agent and docker daemon to be configured to allow privileged containers."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"ipc_mode": {
			Description:  lang.Markdown("The IPC mode to be used for the container. The default is `none` for a private IPC namespace. Other values are `host` for sharing the host IPC namespace or the name or id of an existing container. Note that it is not possible to refer to Docker containers started by Nomad since their names are not known in advance. Note that setting this option also requires the Nomad agent to be configured to allow privileged containers."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("none")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"ipv4_address": {
			Description: lang.Markdown("The IPv4 address to be used for the container when using user defined networks. Requires Docker 1.13 or greater."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"ipv6_address": {
			Description: lang.Markdown("The IPv6 address to be used for the container when using user defined networks. Requires Docker 1.13 or greater."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add example from docs
		"labels": {
			Description: lang.Markdown("The IPv6 address to be used for the container when using user defined networks. Requires Docker 1.13 or greater."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"load": {
			Description: lang.Markdown("Load an image from a `tar` archive file instead of from a remote repository. Equivalent to the `docker load -i <filename>` command. If you're using an `artifact` block to fetch the archive file, you'll need to ensure that Nomad keeps the archive intact after download."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"logging": {
			Description: lang.Markdown("A key-value map of Docker logging options. Defaults to `json-file` with log rotation (`max-file=2` and `max-size=2m`)."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"mac_address": {
			Description: lang.Markdown("The MAC address for the container to use (e.g. \"02:68:b3:29:da:98\")."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"memory_hard_limit": {
			Description: lang.Markdown("The maximum allowable amount of memory used (megabytes) by the container. If set, the [`memory`](https://developer.hashicorp.com/nomad/docs/job-specification/resources#memory) parameter of the task resource configuration becomes a soft limit passed to the docker driver as [`--memory_reservation`](https://docs.docker.com/config/containers/resource_constraints/#limit-a-containers-access-to-memory), and `memory_hard_limit` is passed as the [`--memory`](https://docs.docker.com/config/containers/resource_constraints/#limit-a-containers-access-to-memory) hard limit. When the host is under memory pressure, the behavior of soft limit activation is governed by the [Kernel](https://www.kernel.org/doc/Documentation/cgroup-v1/memory.txt)."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add example
		"network_aliases": {
			Description: lang.Markdown("A list of network-scoped aliases, provide a way for a container to be discovered by an alternate name by any other container within the scope of a particular network. Network-scoped alias is supported only for containers in user defined networks"),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"network_mode": {
			Description:  lang.Markdown("The network mode to be used for the container. In order to support userspace networking plugins in Docker 1.9 this accepts any value. The default is `bridge` for all operating systems but Windows, which defaults to `nat`. Other networking modes may not work without additional configuration on the host (which is outside the scope of Nomad). Valid values pre-docker 1.9 are `default`, `bridge`, `host`, `none`, or `container:name`.\n\nThe default `network_mode` for tasks that use group networking in [`bridge`](https://developer.hashicorp.com/nomad/docs/job-specification/network#bridge) mode will be `container:<name>`, where the name is the container name of the parent container used to share network namespaces between tasks. If you set the group [`network.mode`](https://developer.hashicorp.com/nomad/docs/job-specification/network#mode) to `bridge` you should not set this Docker `network_mode` config, otherwise the container will be unable to reach other containers in the task group. This will also prevent [Connect-enabled](https://developer.hashicorp.com/nomad/docs/job-specification/connect) tasks from reaching the Envoy sidecar proxy. You must also set any DNS options in the `network.dns` block and not in the task configuration.\n\nIf you are in the process of migrating from the default Docker network to group-wide bridge networking, you may encounter issues preventing your containers from reaching networks outside of the bridge interface on systems with firewalld enabled. This behavior is often caused by the CNI plugin not registering the group network as trusted and can be resolved as described in the [network block](https://developer.hashicorp.com/nomad/docs/job-specification/network#bridge-mode) documentation."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("bridge")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"oom_score_adj": {
			Description:  lang.Markdown("A positive integer to indicate the likelihood of the task being OOM killed (valid only for Linux). Defaults to 0."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"pid_mode": {
			Description:  lang.Markdown("`host` or not set (default). Set to `host` to share the PID namespace with the host. Note that this also requires the Nomad agent to be configured to allow privileged containers. See below for more details."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"ports": {
			Description: lang.Markdown("A list of port labels to map into the container"),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"port_map": {
			Description:  lang.Markdown("A key-value map of port labels"),
			Constraint:   &schema.LiteralType{Type: cty.Map(cty.String)},
			IsDeprecated: true,
			IsOptional:   true,
		},
		// TODO: add example
		"security_opt": {
			Description: lang.Markdown("A list of string flags to pass directly to [`--security-opt`](https://docs.docker.com/engine/reference/run/#security-configuration). For example:"),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"shm_size": {
			Description: lang.Markdown("The size (bytes) of /dev/shm for the container."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add example
		"storage_opt": {
			Description: lang.Markdown("A key-value map of storage options set to the containers on start. This overrides the [host dockerd configuration](https://docs.docker.com/engine/reference/commandline/dockerd/#options-per-storage-driver). For example:"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"tty": {
			Description:  lang.Markdown("`true` or `false` (default). Allocate a pseudo-TTY for the container."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"uts_mode": {
			Description:  lang.Markdown("`host` or not set (default). Set to `host` to share the UTS namespace with the host. Note that this also requires the Nomad agent to be configured to allow privileged containers."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"userns_mode": {
			Description:  lang.Markdown("`host` or not set (default). Set to `host` to use the host's user namespace (effectively disabling user namespacing) when user namespace remapping is enabled on the docker daemon. This field has no effect if the docker daemon does not have user namespace remapping enabled."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"volumes": {
			Description: lang.Markdown("A list of `host_path:container_path` strings to bind host paths to container paths. Mounting host paths outside of the [allocation working directory](https://developer.hashicorp.com/nomad/docs/reference/runtime-environment-settings#task-directories) is prevented by default and limits volumes to directories that exist inside the allocation working directory. You can allow mounting host paths outside of the [allocation working directory](https://developer.hashicorp.com/nomad/docs/reference/runtime-environment-settings#task-directories) on individual clients by setting the `docker.volumes.enabled` option to `true` in the [client's configuration](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/docker#client-requirements). We recommend using [`mount`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/docker#mount) if you wish to have more control over volume definitions."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"volume_driver": {
			Description: lang.Markdown("The name of the volume driver used to mount volumes. Must be used along with `volumes`. If `volume_driver` is omitted, then relative paths will be mounted from inside the allocation dir. If a `local` or other driver is used, then they may be named volumes instead. If `docker.volumes.enabled` is false then volume drivers and paths outside the allocation directory are disallowed."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"work_dir": {
			Description: lang.Markdown("The working directory inside the container."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add body from docker docs
		"mount": {
			Description:  lang.Markdown("Specify a [mount](https://docs.docker.com/engine/reference/commandline/service_create/#add-bind-mounts-volumes-or-memory-filesystems) to be mounted into the container. Volume, bind, and tmpfs type mounts are supported. May be specified multiple times."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add mounts here as deprecated
		// "mounts":{}
		// TODO: update body
		"devices": {
			Description: lang.Markdown("A list of [devices](https://docs.docker.com/engine/reference/commandline/run/#add-host-device-to-container-device) to be exposed the container. `host_path` is the only required field. By default, the container will be able to `read`, `write` and `mknod` these devices. Use the optional `cgroup_permissions` field to restrict permissions."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		// TODO: add example
		"cap_add": {
			Description: lang.Markdown("A list of Linux capabilities as strings to pass directly to [`--cap-add`](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities). Effective capabilities (computed from `cap_add` and `cap_drop`) must be a subset of the allowed capabilities configured with the [`allow_caps`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/docker#allow_caps) plugin option key in the client node's configuration. Note that `all` is not permitted here if the `allow_caps` field in the driver configuration doesn't also allow all capabilities."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		// TODO: add example
		"cap_drop": {
			Description: lang.Markdown("A list of Linux capabilities as strings to pass directly to [`--cap-drop`](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities). Effective capabilities (computed from `cap_add` and `cap_drop`) must be a subset of the allowed capabilities configured with the [`allow_caps`](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/docker#allow_caps) plugin option key in the client node's configuration."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"cpu_hard_limit": {
			Description:  lang.Markdown("`true` or `false` (default). Use hard CPU limiting instead of soft limiting. By default this is `false` which means soft limiting is used and containers are able to burst above their CPU limit when there is idle capacity."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"cpu_cfs_period": {
			Description:  lang.Markdown("An integer value that specifies the duration in microseconds of the period during which the CPU usage quota is measured. The default is 100000 (0.1 second) and the maximum allowed value is 1000000 (1 second). See [here](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/sec-cpu#sect-cfs) for more details."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(100000)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"advertise_ipv6_address": {
			Description:  lang.Markdown("`true` or `false` (default). Use the container's IPv6 address (GlobalIPv6Address in Docker) when registering services and checks. See [IPv6 Docker containers](https://developer.hashicorp.com/nomad/docs/job-specification/service#ipv6-docker-containers) for details."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"readonly_rootfs": {
			Description:  lang.Markdown("`true` or `false` (default). Mount the container's filesystem as read only."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"runtime": {
			Description:  lang.Markdown("A string representing a configured runtime to pass to docker. This is equivalent to the `--runtime` argument in the docker CLI For example, to use gVisor:\n\n```hcl\nconfig{\nruntime = \"runsc\"\n}\n```"),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"pids_limit": {
			Description: lang.Markdown("An integer value that specifies the pid limit for the container. Defaults to unlimited."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}

// var HealthchecksSchema = &schema.BodySchema{
// 	Attributes: map[string]*schema.AttributeSchema{
// 		"disable": {
// 			Description:  lang.Markdown("The Docker image to run. The image may include a tag or custom URL and should include `https://` if required. By default it will be fetched from Docker Hub. If the tag is omitted or equal to `latest` the driver will always try to pull the image. If the image to be pulled exists in a registry that requires authentication credentials must be provided to Nomad."),
// 			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
// 			Constraint:   &schema.LiteralType{Type: cty.Bool},
// 		},
// 	},
// }
