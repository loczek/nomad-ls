package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/schema/drivers"
	"github.com/zclconf/go-cty/cty"
)

var TaskSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"driver": {
			Description: lang.Markdown("Specifies the task driver that should be used to run the task. See the [driver documentation](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver) for what is available. Examples include `docker`, `qemu`, `java` and `exec`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsRequired: true,
			IsDepKey:   true,
		},
		"kill_timeout": {
			Description: lang.Markdown("Specifies the duration to wait for an application to gracefully quit before force-killing. Nomad first sends a [`kill_signal`](https://developer.hashicorp.com/nomad/docs/job-specification/task#kill_signal). If the task does not exit before the configured timeout, `SIGKILL` is sent to the task. Note that the value set here is capped at the value set for [`max_kill_timeout`](https://developer.hashicorp.com/nomad/docs/configuration/client#max_kill_timeout) on the agent running the task, which has a default value of 30 seconds."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("5s"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"kill_signal": {
			Description: lang.Markdown("Specifies a configurable kill signal for a task, where the default is SIGINT (or SIGTERM for `docker`, or CTRL_BREAK_EVENT for `raw_exec` on Windows). Note that this is only supported for drivers sending signals (currently `docker`, `exec`, `raw_exec`, and `java` drivers)."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("5s"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"leader": {
			Description: lang.Markdown("Specifies whether the task is the leader task of the task group. If set to `true`, when the leader task completes, all other tasks within the task group will be gracefully shutdown. The shutdown process starts by applying the `shutdown_delay` if configured. It then stops the the leader task first, if any, followed by non-sidecar and non-poststop tasks, and finally sidecar tasks. Once this process completes, post-stop tasks are triggered. See the [lifecycle](https://developer.hashicorp.com/nomad/docs/job-specification/lifecycle) documentation for a complete description of task lifecycle management."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},

			Constraint: &schema.LiteralType{Type: cty.Bool},
		},
		"shutdown_delay": {
			Description: lang.Markdown("Specifies the duration to wait when killing a task between removing its service registrations from Consul or Nomad, and sending it a shutdown signal. Ideally services would fail health checks once they receive a shutdown signal. Alternatively, `shutdown_delay` may be set to give in flight requests time to complete before shutting down. This `shutdown_delay` only applies to services defined at the task level by the [`service`](https://developer.hashicorp.com/nomad/docs/job-specification/task#service) block. In addition, task groups have their own [`shutdown_delay`](https://developer.hashicorp.com/nomad/docs/job-specification/group#shutdown_delay) which waits between de-registering group services and stopping tasks."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("0s"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"user": {
			Description: lang.Markdown("Specifies the user that will run the task. Defaults to `nobody` for the [`exec`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/exec) and [`java`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/java) drivers. [Docker](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/docker) images specify their own default users. Clients can restrict [which drivers](https://developer.hashicorp.com/nomad/docs/configuration/client#user-checked_drivers) are allowed to run tasks as [certain users](https://developer.hashicorp.com/nomad/docs/configuration/client#user-denylist). On UNIX-like systems, setting `user` also affects the environment variables `HOME`, `USER`, and `LOGNAME` available to the task. On Windows, when Nomad is running as a [system service](https://developer.hashicorp.com/nomad/docs/job-specification/service) for the [`raw_exec`](https://developer.hashicorp.com/nomad/docs/job-declare/task-driver/raw_exec) driver, you may specify a less-privileged service user. For example, `NT AUTHORITY\\LocalService`, `NT AUTHORITY\\NetworkService`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"kind": {
			Description: lang.Markdown("Used internally to manage tasks according to the value of this field. Initial use case is for Consul service mesh."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"config": {
			Description: lang.PlainText("Specifies the driver configuration, which is passed directly to the driver to start the task. The details of configurations are specific to each driver, so please see specific driver documentation for more information."),
			DependentBody: map[schema.SchemaKey]*schema.BodySchema{
				"docker":   drivers.DockerDriverSchema,
				"exec":     drivers.ExecDriverSchema,
				"raw_exec": drivers.RawExecDriverSchema,
				"java":     drivers.JavaDriverSchema,
				"qemu":     drivers.QemuDriverSchema,
			},
		},
		"artifact": {
			Description: lang.PlainText("Defines an artifact to download before running the task. This may be specified multiple times to download multiple artifacts."),
			Body:        ArtifactSchema,
		},
		"consul": {
			Description: lang.PlainText(" Consul configuration options specific to the task."),
			Body:        ConsulSchema,
		},
		"constraint": {
			Description: lang.PlainText("Specifies user-defined constraints on the task. This can be provided multiple times to define additional constraints."),
			Body:        ConstraintSchema,
		},
		"csi_plugin": {
			Description: lang.PlainText("Specifies user-defined constraints on the task. This can be provided multiple times to define additional constraints."),
			Body:        CsiPluginSchema,
		},
		"affinity": {
			Description: lang.PlainText("This can be provided multiple times to define preferred placement criteria."),
			Body:        AffinitySchema,
		},
		"dispatch_payload": {
			Description: lang.PlainText("Configures the task to have access to dispatch payloads."),
			Body:        DispatchPayloadSchema,
		},
		"env": {
			Description: lang.PlainText("Specifies environment variables that will be passed to the running process."),
			Type:        schema.BlockTypeMap,
		},
		"identity": {
			Description: lang.PlainText("Expose [Workload Identity](https://developer.hashicorp.com/nomad/docs/concepts/workload-identity) to the task."),
			Body:        IdentitySchema,
		},
		"lifecycle": {
			Description: lang.PlainText("Specifies when a task is run within the lifecycle of a task group. Added in Nomad v0.11."),
			Body:        LifecycleSchema,
		},
		"logs": {
			Description: lang.PlainText("Specifies logging configuration for the stdout and stderr of the task."),
			Body:        LogsSchema,
		},
		"meta": {
			Description: lang.PlainText("Specifies a key-value map that annotates with user-defined metadata."),
			Body:        MetaSchema,
		},
		"resources": {
			Description: lang.PlainText("Specifies the minimum resource requirements such as RAM, CPU and devices."),
			Body:        ResourcesSchema,
		},
		"restart": {
			Description: lang.PlainText("Specifies the restart policy for all tasks in this group. If omitted, a default policy exists for each job type, which can be found in the [restart block documentation](https://developer.hashicorp.com/nomad/docs/job-specification/restart)."),
			Body:        RestartSchema,
		},
		"service": {
			Description: lang.Markdown("Specifies integrations with Nomad or [Consul](https://www.consul.io/) for service discovery. Nomad automatically registers when a task is started and de-registers it when the task dies."),
			Body:        ServiceSchema,
		},
		// TODO: add docs
		"scaling": {
			Description: lang.Markdown("scaling docs"),
			Body:        ScalingSchema,
		},
		// TODO: add docs
		"secret": {
			Description: lang.PlainText("secret docs"),
			Body:        SecretSchema,
		},
		"schedule": {
			Description: lang.PlainText("schedule docs"),
			Body:        ScheduleSchema,
		},
		"template": {
			Description: lang.PlainText("Specifies the set of templates to render for the task. Templates can be used to inject both static and dynamic configuration with data populated from environment variables, Consul and Vault."),
			Body:        TemplateSchema,
		},
		"vault": {
			Description: lang.PlainText("Specifies the set of Vault policies required by the task. This overrides any `vault` block set at the `group` or `job` level."),
			Body:        VaultSchema,
		},
		"volume_mount": {
			Description: lang.PlainText("Specifies where a group volume should be mounted."),
			Body:        VolumeMountSchema,
		},
	},
}
