package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var ConsulClSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"address": {
			Description:  lang.Markdown("Specifies the address to the local Consul agent, given in the format `host:port`. Supports Unix sockets with the format: `unix:///tmp/consul/consul.sock`. Will default to the `CONSUL_HTTP_ADDR` environment variable if set. The value supports [go-sockaddr/template format](https://pkg.go.dev/github.com/hashicorp/go-sockaddr/template)."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("127.0.0.1:8500")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"auth": {
			Description:  lang.Markdown("Specifies the HTTP Basic Authentication information to use for access to the Consul Agent, given in the format `username:password`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"auto_advertise": {
			Description:  lang.Markdown("Specifies if Nomad should advertise its services in Consul. The services are named according to `server_service_name` and `client_service_name`. Nomad servers and clients advertise their respective services, each tagged appropriately with either `http` or `rpc` tag. Nomad servers also advertise a `serf` tagged service."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"ca_file": {
			Description:  lang.Markdown("Specifies an optional path to the CA certificate used for Consul communication. This defaults to the system bundle if unspecified. Will default to the `CONSUL_CACERT` environment variable if set."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"cert_file": {
			Description:  lang.Markdown("Specifies the path to the certificate used for Consul communication. If this is set then you need to also set `key_file`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"checks_use_advertise": {
			Description:  lang.Markdown("Specifies if Consul health checks should bind to the advertise address. By default, this is the first [HTTP address](https://developer.hashicorp.com/nomad/docs/configuration#http). If no [HTTP address](https://developer.hashicorp.com/nomad/docs/configuration#http) is specified, it will fall back to the [bind_addr](https://developer.hashicorp.com/nomad/docs/configuration#bind_addr)."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"key_file": {
			Description:  lang.Markdown("Specifies the path to the private key used for Consul communication. If this is set then you need to also set `cert_file`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"name": {
			Description:  lang.Markdown("Specifies a name for the cluster so it can be referred to by job submitters in the job specification's [`consul.cluster`](https://developer.hashicorp.com/nomad/docs/job-specification/consul#cluster) or [`service.cluster`](https://developer.hashicorp.com/nomad/docs/job-specification/service#cluster) fields. In Nomad Community Edition, only the `\"default\"` cluster will be used, so this field should be omitted." + schemautils.EnterpriseOnly),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"namespace": {
			Description:  lang.Markdown("Specifies the [Consul namespace](https://developer.hashicorp.com/consul/docs/enterprise/namespaces) used by the Consul integration. If non-empty, this namespace will be used on all Consul API calls and for Consul service mesh configurations, unless overridden by the job's [`consul.namespace`](https://developer.hashicorp.com/nomad/docs/job-specification/consul#namespace) field. In Nomad Community Edition, only the \"default\" namespace is used, so you should omit this field." + schemautils.EnterpriseOnly),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"ssl": {
			Description:  lang.Markdown("Specifies if the transport scheme should use HTTPS to communicate with the Consul agent. Will default to the `CONSUL_HTTP_SSL` environment variable if set."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"tags": {
			Description:  lang.Markdown("Specifies optional Consul tags to be registered with the Nomad server and client services."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"timeout": {
			Description:  lang.Markdown("Specifies a time limit for requests made against Consul. This is specified using a label suffix like \"10s\"."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"token": {
			Description:  lang.Markdown("Specifies the token used to provide a per-request ACL token. This option overrides the Consul Agent's default token. If the token is not set here or on the Consul agent, it defaults to Consul's anonymous policy, which may or may not allow writes. Defaults to the `CONSUL_HTTP_TOKEN` environment variable if set. Nomad cannot refresh this token; if the token is deleted, Nomad is not able to communicate with Consul. Nomad also looks for the environment variable `CONSUL_HTTP_TOKEN_*name*`, where name is the `consul.name` parameter. This allows Nomad Enterprise users to specify multiple tokens for multiple clusters via environment variables. In Nomad Enterprise, if the Consul agent running alongside Nomad is in a Consul Enterprise admin partition, you must create the Consul token provided to the Nomad client in the same partition."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"verify_ssl": {
			Description:  lang.Markdown("Specifies if SSL peer verification should be used when communicating to the Consul API client over HTTPS. Will default to the `CONSUL_HTTP_SSL_VERIFY` environment variable if set."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},

		// CLIENT ONLY
		"client_auto_join": {
			Description:  lang.Markdown("Specifies if the Nomad clients should automatically discover Nomad servers in the same region by searching for the Consul service name defined in the `server_service_name` option. The search occurs if the client is not registered with any Nomad servers or it is unable to heartbeat to the leader of the region, in which case it may be partitioned and searches for other Nomad servers."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"client_service_name": {
			Description:  lang.Markdown("Specifies the name of the service in Consul for the Nomad clients."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("nomad-client")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"client_http_check_name": {
			Description:  lang.Markdown("Specifies the HTTP health check name in Consul for the Nomad clients."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("Nomad Client HTTP Check")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"client_failures_before_critical": {
			Description:  lang.Markdown("Specifies the number of consecutive failures before the Nomad client Consul health check is critical."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"client_failures_before_warning": {
			Description:  lang.Markdown("Specifies the number of consecutive failures before the Nomad client Consul health check shows a warning."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"grpc_address": {
			Description:  lang.Markdown("Specifies the address to the local Consul agent for `gRPC` requests, given in the format `host:port`. Note that Consul does not enable the [`grpc`](https://developer.hashicorp.com/consul/docs/agent/config/config-files#grpc_port) or [`grpc_tls`](https://developer.hashicorp.com/consul/docs/agent/config/config-files#grpc_tls_port) listeners by default."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("127.0.0.1:8502")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		// TODO: add warning
		"grpc_ca_file": {
			Description:  lang.Markdown("Specifies an optional path to the GRPC CA certificate used for communication between Connect sidecar proxies and Consul agents. Will default to the `CONSUL_GRPC_CACERT` environment variable if set."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"share_ssl": {
			Description:  lang.Markdown("Specifies whether the Nomad client should share its Consul SSL configuration with Connect Native applications. Includes values of `ca_file`, `cert_file`, `key_file`, `ssl`, and `verify_ssl`. Does not include the values for the ACL `token` or `auth`. This option should be disabled in environments where Consul ACLs are not enabled."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"service_auth_method": {
			Description:  lang.Markdown("Specifies the name of the Consul [authentication method](https://developer.hashicorp.com/consul/docs/security/acl/auth-methods/jwt) that will be used to login with a Nomad JWT for services."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("nomad-workloads")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"task_auth_method": {
			Description:  lang.Markdown("Specifies the name of the Consul [authentication method](https://developer.hashicorp.com/consul/docs/security/acl/auth-methods/jwt) that will be used to login with a Nomad JWT for tasks."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("nomad-workloads")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},

		// SERVER ONLY
		"server_service_name": {
			Description:  lang.Markdown("Specifies the name of the service in Consul for the Nomad servers."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("nomad")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"server_http_check_name": {
			Description:  lang.Markdown("Specifies the HTTP health check name in Consul for the Nomad servers."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("Nomad Server HTTP Check")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"server_serf_check_name": {
			Description:  lang.Markdown("Specifies the Serf health check name in Consul for the Nomad servers."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("Nomad Server Serf Check")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"server_rpc_check_name": {
			Description:  lang.Markdown("Specifies the RPC health check name in Consul for the Nomad servers."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("Nomad Server RPC Check")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"server_auto_join": {
			Description:  lang.Markdown("Specifies if the Nomad servers should automatically discover and join other Nomad servers by searching for the Consul service name defined in the `server_service_name` option. This search only happens if the server does not have a leader."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"server_failures_before_critical": {
			Description:  lang.Markdown("Specifies the number of consecutive failures before the Nomad server Consul health check is critical."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"server_failures_before_warning": {
			Description:  lang.Markdown("Specifies the number of consecutive failures before the Nomad server Consul health check shows a warning."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		// SERVER ONLY
		"service_identity": {
			Description: lang.Markdown("Specifies a default Workload Identity to use when obtaining Service Identity tokens from Consul to register services. Refer to [Workload Identity](https://developer.hashicorp.com/nomad/docs/configuration/consul#workload-identity) for a recommended configuration."),
			Body:        ServiceIdentitySchema,
		},
		"task_identity": {
			Description: lang.Markdown("Specifies a default Workload Identity to use when obtaining Consul tokens from Consul to support [`template`](https://developer.hashicorp.com/nomad/docs/job-specification/template) blocks. Refer to [Workload Identity](https://developer.hashicorp.com/nomad/docs/configuration/consul#workload-identity) for a recommended configuration."),
			Body:        TaskIdentitySchema,
		},
	},
}

var ServiceIdentitySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"aud": {
			Description:  lang.Markdown("List of valid recipients for this workload identity. This value must match the [`BoundAudiences`](https://developer.hashicorp.com/consul/docs/security/acl/auth-methods/jwt#boundaudiences) configuration in the Consul JWT auth method. It is recommended to provide one, and only one, audience to minimize where the identity may be used."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"ttl": {
			Description:  lang.Markdown("Specifies for how long the workload identity should be considered as valid before expiring."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
var TaskIdentitySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"aud": {
			Description:  lang.Markdown("List of valid recipients for this workload identity. This value must match the [`BoundAudiences`](https://developer.hashicorp.com/consul/docs/security/acl/auth-methods/jwt#boundaudiences) configuration in the Consul JWT auth method. It is recommended to provide one, and only one, audience to minimize where the identity may be used."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"env": {
			Description:  lang.Markdown("If true the workload identity will be available in the task's `NOMAD_TOKEN_consul_default` (or `NOMAD_TOKEN_consul_<name>` depending on the [`name`](https://developer.hashicorp.com/nomad/docs/configuration/consul#name) field) environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"file": {
			Description:  lang.Markdown("If true the workload identity will be available in the task's filesystem via the path `secrets/nomad_consul_default.jwt` (or `secrets/nomad_consul_<name>.jwt` depending on the [`name`](https://developer.hashicorp.com/nomad/docs/configuration/consul#name) field). If the [`task.user`](https://developer.hashicorp.com/nomad/docs/job-specification/task#user) parameter is set, the token file will only be readable by that user. Otherwise the file is readable by everyone but is protected by parent directory permissions."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"ttl": {
			Description:  lang.Markdown("Specifies for how long the workload identity should be considered as valid before expiring."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
