package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/loczek/nomad-ls/internal/schema/drivers"
)

const (
	variablesLabel = "variables"
	variableLabel  = "variable"
	localsLabel    = "locals"
	vaultLabel     = "vault"
	taskLabel      = "task"
	secretLabel    = "secret"

	inputVariablesAccessor = "var"
	localsAccessor         = "local"
)

var SchemaMapBetter map[string]*hcl.BodySchema = map[string]*hcl.BodySchema{
	"root": RootBodySchema.Copy().ToHCLSchema(),

	"variable": VariableSchema.Copy().ToHCLSchema(),

	"affinity":         AffinitySchema.Copy().ToHCLSchema(),
	"artifact":         ArtifactSchema.Copy().ToHCLSchema(),
	"change_script":    ChangeScriptSchema.Copy().ToHCLSchema(),
	"check":            CheckSchema.Copy().ToHCLSchema(),
	"check_restart":    CheckRestartSchema.Copy().ToHCLSchema(),
	"cni":              CniSchema.Copy().ToHCLSchema(),
	"connect":          ConnectSchema.Copy().ToHCLSchema(),
	"constraint":       ConstraintSchema.Copy().ToHCLSchema(),
	"consul":           ConsulSchema.Copy().ToHCLSchema(),
	"dispatch_payload": DispatchPayloadSchema.Copy().ToHCLSchema(),
	"dns":              DnsSchema.Copy().ToHCLSchema(),
	"env":              EnvSchema.Copy().ToHCLSchema(),
	"ephemeral_disk":   EphemeralDiskSchema.Copy().ToHCLSchema(),
	"expose":           ExposeSchema.Copy().ToHCLSchema(),
	"group":            GroupSchema.Copy().ToHCLSchema(),
	"identity":         IdentitySchema.Copy().ToHCLSchema(),
	"job":              JobSchemaBetter.Copy().ToHCLSchema(),
	"lifecycle":        LifecycleSchema.Copy().ToHCLSchema(),
	"listener_port":    PortSchema.Copy().ToHCLSchema(),
	"logs":             LogsSchema.Copy().ToHCLSchema(),
	"mesh_gateway":     MeshGatewaySchema.Copy().ToHCLSchema(),
	"meta":             MetaSchema.Copy().ToHCLSchema(),
	"migrate":          MigrateSchema.Copy().ToHCLSchema(),
	"network":          NetworkSchema.Copy().ToHCLSchema(),
	"parameterized":    ParameterizedSchema.Copy().ToHCLSchema(),
	"path":             PathSchema.Copy().ToHCLSchema(),
	"periodic":         PeriodicSchema.Copy().ToHCLSchema(),
	"port":             PortSchema.Copy().ToHCLSchema(),
	"proxy":            ProxySchema.Copy().ToHCLSchema(),
	"reschedule":       RescheduleSchema.Copy().ToHCLSchema(),
	"resources":        ResourcesSchema.Copy().ToHCLSchema(),
	"restart":          RestartSchema.Copy().ToHCLSchema(),
	"service":          ServiceSchema.Copy().ToHCLSchema(),
	"sidecar_service":  SidecarServiceSchema.Copy().ToHCLSchema(),
	"sidecar_task":     SidecarTaskSchema.Copy().ToHCLSchema(),
	"spread":           SpreadSchema.Copy().ToHCLSchema(),
	"target":           TargetSchema.Copy().ToHCLSchema(),
	"task":             TaskSchema.Copy().ToHCLSchema(),
	"template":         TemplateSchema.Copy().ToHCLSchema(),
	"update":           UpdateSchema.Copy().ToHCLSchema(),
	"upstreams":        UpstreamsSchema.Copy().ToHCLSchema(),
	"volume":           VolumeSchema.Copy().ToHCLSchema(),
	"volume_mount":     VolumeMountSchema.Copy().ToHCLSchema(),
	"wait":             WaitSchema.Copy().ToHCLSchema(),

	"transparent_proxy": TransparentProxySchema.Copy().ToHCLSchema(),

	"config:docker":   drivers.DockerDriverSchema.Copy().ToHCLSchema(),
	"config:raw_exec": drivers.RawExecDriverSchema.Copy().ToHCLSchema(),
	"config:exec":     drivers.ExecDriverSchema.Copy().ToHCLSchema(),
	"config:qemu":     drivers.QemuDriverSchema.Copy().ToHCLSchema(),
	"config:java":     drivers.JavaDriverSchema.Copy().ToHCLSchema(),
}

var RootBodySchema = schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"variable": {
			Description: lang.Markdown("## h2\nvariable docs"),
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: VariableSchema,
		},
		"job": {
			Description: lang.Markdown("## h2\njob docs"),
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: JobSchemaBetter,
		},
	},
}
