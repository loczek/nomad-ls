package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/schema/agent/keyring"
	"github.com/zclconf/go-cty/cty"
)

var RootSchema = schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"bind_addr": {
			Description:  lang.Markdown("Specifies which address the Nomad agent should bind to for network services, including the HTTP interface as well as the internal gossip protocol and RPC mechanism. This should be specified in IP format, and can be used to easily bind all network services to the same address. It is also possible to bind the individual services to different addresses using the [addresses](https://developer.hashicorp.com/nomad/docs/configuration#addresses) configuration option. Dev mode (`-dev`) defaults to localhost. The value supports [go-sockaddr/template format](https://pkg.go.dev/github.com/hashicorp/go-sockaddr/template)."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("0.0.0.0")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"datacenter": {
			Description:  lang.Markdown("Specifies the data center of the local agent. A datacenter is an abstract grouping of clients within a region. Clients are not required to be in the same datacenter as the servers they are joined with, but do need to be in the same region."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("dc1")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"data_dir": {
			Description: lang.Markdown("Specifies a local directory used to store agent state. Client nodes use this directory by default to store temporary allocation data as well as cluster information. Server nodes use this directory to store cluster state, including the replicated log and snapshot data. This must be specified as an absolute path. Nomad will create the directory on the host, if it does not exist when the agent process starts."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"disable_anonymous_signature": {
			Description:  lang.Markdown("Specifies if Nomad should provide an anonymous signature for de-duplication with the update check."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"disable_update_check": {
			Description:  lang.Markdown("Specifies if Nomad should not check for updates and security bulletins. This defaults to `true` in Nomad Enterprise."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"enable_debug": {
			Description:  lang.Markdown("Specifies if the debugging HTTP endpoints should be enabled. These endpoints can be used with profiling tools to dump diagnostic information about Nomad's internals."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"enable_syslog": {
			Description:  lang.Markdown("Specifies if the agent should log to syslog. This option only works on Unix based systems. The log level inherits from the Nomad agent log set in `log_level`"),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"leave_on_interrupt": {
			Description:  lang.Markdown("Specifies if the agent should leave when receiving the interrupt signal. By default, any stop signal to an agent (interrupt or terminate) will cause the agent to exit after ensuring its internal state is committed to disk as needed. If this value is set to true on a server agent, the server will notify other servers of their intention to leave the peer set. You should only set this value to true on server agents if the terminated server will never join the cluster again. If this value is set to true on a client agent and the client is configured with [`drain_on_shutdown`](https://developer.hashicorp.com/nomad/docs/configuration/client#drain_on_shutdown), the client will drain its workloads before shutting down."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"leave_on_terminate": {
			Description:  lang.Markdown("Specifies if the agent should leave when receiving the terminate signal. By default, any stop signal to an agent (interrupt or terminate) will cause the agent to exit after ensuring its internal state is committed to disk as needed. If this value is set to true on a server agent, the server will notify other servers of their intention to leave the peer set. You should only set this value to true on server agents if the terminated server will never join the cluster again. If this value is set to true on a client agent and the client is configured with [`drain_on_shutdown`](https://developer.hashicorp.com/nomad/docs/configuration/client#drain_on_shutdown), the client will drain its workloads before shutting down."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"log_level": {
			Description: lang.Markdown("Specifies the verbosity of logs the Nomad agent will output. Valid log levels include `WARN`, `INFO`, `DEBUG`, or `TRACE` in increasing order of verbosity."),
			Constraint: schema.OneOf{
				schema.LiteralValue{Value: cty.StringVal("WARN")},
				schema.LiteralValue{Value: cty.StringVal("INFO")},
				schema.LiteralValue{Value: cty.StringVal("DEBUG")},
				schema.LiteralValue{Value: cty.StringVal("TRACE")},
			},
			IsOptional: true,
		},
		"log_include_location": {
			Description:  lang.Markdown("Include file and line information in each log line."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"log_json": {
			Description:  lang.Markdown("Output logs in a JSON format."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"log_file": {
			Description:  lang.Markdown("Specifies the path for logging. If the path does not includes a filename, the filename defaults to `nomad.log`. This setting can be combined with `log_rotate_bytes` and `log_rotate_duration` for a fine-grained log rotation control."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"log_rotate_bytes": {
			Description:  lang.Markdown("Specifies the number of bytes that should be written to a log before it needs to be rotated. Unless specified, there is no limit to the number of bytes that can be written to a log file."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"log_rotate_duration": {
			Description:  lang.Markdown("Specifies the maximum duration a log should be written to before it needs to be rotated. Must be a duration value such as 30s."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("24h")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"log_rotate_max_files": {
			Description:  lang.Markdown("Specifies the maximum number of older log file archives to keep, not including the log file currently being written. If set to 0 no files are ever deleted. Note that the total number of log files, for each of `stderr` and `stdout`, will be 1 greater than the `log_rotate_max_files` value."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"name": {
			Description:  lang.Markdown("Specifies the name of the local node. This value is used to identify individual agents. When specified on a server, the name must be unique within the region."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("[hostname]")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"plugin_dir": {
			Description:  lang.Markdown("Specifies the directory to use for looking up plugins. When this parameter is empty, Nomad will generate the path using the [top-level `data_dir`](https://developer.hashicorp.com/nomad/docs/configuration#data_dir) suffixed with `plugins`, like `\"/opt/nomad/plugins\"`. This must be an absolute path. If you are using plugins that execute as unprivileged users, such as `exec2`, you should set the `plugin_dir` outside the `data_dir` and allow it to be executable."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add warning
		"region": {
			Description:  lang.Markdown("Specifies the region the Nomad agent is a member of. A region typically maps to a geographic region, for example us, with potentially multiple zones, which map to datacenters such as us-west and us-east."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("global")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"syslog_facility": {
			Description:  lang.Markdown("Specifies the syslog facility to write to. This has no effect unless enable_syslog is true."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("LOCAL0")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"acl": {
			Description: ACLSchema.Description,
			Body:        ACLSchema,
		},
		"addresses": {
			Description: lang.Markdown("Specifies the bind address for individual network services. Any values configured in this block take precedence over the default [bind_addr](https://developer.hashicorp.com/nomad/docs/configuration#bind_addr). These values should be specified in IP format without a port (ex. \"0.0.0.0\"). To set the port, refer to the [`ports`](https://developer.hashicorp.com/nomad/docs/configuration#ports) field. The values support [go-sockaddr/template format](https://pkg.go.dev/github.com/hashicorp/go-sockaddr/template)."),
			Body:        AddressesSchema,
		},
		"advertise": {
			Description: lang.Markdown("Specifies the advertise address for individual network services. This can be used to advertise a different address to the peers of a server or a client node to support more complex network configurations such as NAT. This configuration is optional, and defaults to the bind address of the specific network service if it is not provided. Any values configured in this block take precedence over the default [bind_addr](https://developer.hashicorp.com/nomad/docs/configuration#bind_addr).\n\nIf the bind address is `0.0.0.0` then the IP address of the default private network interface advertised. The `advertise` values may include an alternate port, but otherwise default to the port used by the bind address. The values support [go-sockaddr/template format](https://pkg.go.dev/github.com/hashicorp/go-sockaddr/template)."),
			Body:        AdvertiseSchema,
		},
		"audit": {
			Description: AuditSchema.Description,
			Body:        AuditSchema,
		},
		// TODO: add description (missing from docs)
		"autopilot": {
			Body: AutopilotSchema,
		},
		"client": {
			Description: lang.Markdown("Specifies configuration which is specific to the Nomad client."),
			Body:        ClientSchema,
		},
		"consul": {
			Description: lang.Markdown("Specifies configuration for connecting to Consul."),
			Body:        ConsulClSchema,
		},
		"eventlog": {
			Description: lang.Markdown("This is a nested object that configures the behavior with with Windows Event Log"),
			Body:        EventlogSchema,
		},
		"http_api_response_headers": {
			Body: &schema.BodySchema{
				Description: lang.Markdown("Specifies user-defined headers to add to the HTTP API responses."),
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
					IsOptional: true,
				},
			},
		},
		"keyring": {
			Labels: []*schema.LabelSchema{
				{
					Name:        "name",
					IsDepKey:    true,
					Completable: true,
				},
			},
			Body: keyring.KeyringSchema,
			DependentBody: map[schema.SchemaKey]*schema.BodySchema{
				labelKey("awskms"):        keyring.AWSSchema,
				labelKey("azurekeyvault"): keyring.AzureSchema,
				labelKey("gcpckms"):       keyring.GCPSchema,
				labelKey("transit"):       keyring.VaultSchema,
			},
		},
		"limits": {
			Description: lang.Markdown("This is a nested object that configures limits that are enforced by the agent"),
			Body:        LimitsSchema,
		},
		// TODO: add this later
		"plugin": {
			// Description: lang.Markdown("This is a nested object that configures limits that are enforced by the agent"),
			// Body:        LimitsSchema,
		},
		"ports": {
			Description: lang.Markdown("Specifies the network ports used for different services required by the Nomad agent."),
			Body:        PortsSchema,
		},
		// TODO: add description (missing from docs)
		"reporting": {
			Description: lang.Markdown("reporting docs"),
			Body:        ReportingSchema,
		},
		"rpc": {
			Description: lang.Markdown("Specifies configuration which is specific to RPC. We strongly recommend that you do not configure RPC values. Use the default values, which have been production-proven on clusters of thousands of nodes. You should only configure the `rpc` block if you have a specific reason to believe it will improve your particular use case."),
			Body:        RPCSchema,
		},
		"sentinel": {
			Description: lang.Markdown("Specifies configuration for Sentinel policies."),
			Body:        SentinelSchema,
		},
		"server": {
			Description: lang.Markdown("Specifies configuration which is specific to the Nomad server."),
			Body:        ServerSchema,
		},
		// TODO: add description (missing from docs)
		"telemetry": {
			Description: lang.Markdown("telemetry docs"),
			Body:        TelemetrySchema,
		},
		"tls": {
			Description: lang.Markdown("Specifies configuration for TLS."),
			Body:        TLSSchema,
		},
		"vault": {
			Description: lang.Markdown("Specifies configuration for connecting to Vault."),
			// Body:        vault,
		},
	},
}

func labelKey(value string) schema.SchemaKey {
	return schema.NewSchemaKey(schema.DependencyKeys{
		Labels: []schema.LabelDependent{
			{
				Index: 0,
				Value: value,
			},
		},
	})
}
