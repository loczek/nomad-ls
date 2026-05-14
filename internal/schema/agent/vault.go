package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var VaultSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description:  lang.Markdown("Specifies a name for the cluster so it can be referred to by job submitters in the job specification's [`vault.cluster`](https://developer.hashicorp.com/nomad/docs/job-specification/vault#cluster) field. In Nomad Community Edition, only the `\"default\"` cluster will be used, so this field should be omitted." + schemautils.Divider + schemautils.EnterpriseOnly),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"enabled": {
			Description:  lang.Markdown("Specifies if the Vault integration should be activated."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"namespace": {
			Description:  lang.Markdown("Specifies the [Vault namespace](https://developer.hashicorp.com/vault/docs/enterprise/namespaces) used by the Vault integration. If non-empty, this namespace will be used on all Vault API calls."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"address": {
			Description:  lang.Markdown("Specifies the address to the Vault server. This must include the protocol, host/ip, and port given in the format `protocol://host:port`. If your Vault installation is behind a load balancer, this should be the address of the load balancer."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("https://vault.service.consul:8200")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"jwt_auth_backend_path": {
			Description:  lang.Markdown("Specifies the mount [path](https://developer.hashicorp.com/vault/docs/commands/auth/enable#path) of the JWT authentication method to be used to login with workload identity JWTs."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("jwt-nomad")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"ca_file": {
			Description:  lang.Markdown("Specifies an optional path to the CA certificate used for Vault communication. If unspecified, this will fallback to the default system CA bundle, which varies by OS and version."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"ca_path": {
			Description:  lang.Markdown("Specifies an optional path to a folder containing CA certificates to be used for Vault communication. If unspecified, this will fallback to the default system CA bundle, which varies by OS and version."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"cert_file": {
			Description:  lang.Markdown("Specifies the path to the certificate used for Vault communication. This must be set if [tls_require_and_verify_client_cert](https://developer.hashicorp.com/vault/docs/configuration/listener/tcp#tls_require_and_verify_client_cert) is enabled in Vault."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"create_from_role": {
			Description:  lang.Markdown("Specifies the role to create tokens from. This field defines the role used to derive task tokens when the job does not define a value for [`vault.role`](https://developer.hashicorp.com/nomad/docs/job-specification/vault#role). If empty, the default Vault cluster role is used."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"key_file": {
			Description:  lang.Markdown("Specifies the path to the private key used for Vault communication. If this is set then you need to also set `cert_file`. This must be set if [tls_require_and_verify_client_cert](https://developer.hashicorp.com/vault/docs/configuration/listener/tcp#tls_require_and_verify_client_cert) is enabled in Vault."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_server_name": {
			Description:  lang.Markdown("Specifies an optional string used to set the SNI host when connecting to Vault via TLS."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_skip_verify": {
			Description:  lang.Markdown("Specifies if SSL peer validation should be enforced."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"default_identity": {
			Description: lang.Markdown("Specifies the default workload identity configuration to use when a task with a `vault` block does not specify an [`identity`](https://developer.hashicorp.com/nomad/docs/job-specification/identity) block named `vault_<name>`, where `<name>` matches the value of this `vault` block [`name`](https://developer.hashicorp.com/nomad/docs/configuration/vault#name) parameter. Setting a default identity causes the value of `allow_unauthenticated` to be ignored."),
			Body:        DefaultIdentitySchema,
		},
	},
}

var DefaultIdentitySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"aud": {
			Description:  lang.Markdown("List of valid recipients for this workload identity. This value must match the [`bound_audiences`](https://developer.hashicorp.com/vault/api-docs/auth/jwt#bound_audiences) configuration in the Vault JWT auth method. It is recommended to provide one, and only one, audience to minimize where the identity may be used."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"env": {
			Description:  lang.Markdown("If true the workload identity will be available in the task's `NOMAD_TOKEN_vault` environment variable."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"file": {
			Description:  lang.Markdown("If true the workload identity will be available in the task's filesystem via the path `secrets/nomad_vault.jwt`. If the [`task.user`](https://developer.hashicorp.com/nomad/docs/job-specification/task#user \"Nomad task Block\") parameter is set, the token file will only be readable by that user. Otherwise the file is readable by everyone but is protected by parent directory permissions."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"ttl": {
			Description:  lang.Markdown("Specifies for how long the workload identity should be considered as valid before expiring."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"extra_claims": {
			Description: lang.Markdown("A set of key-value pairs that will be provided as extra identity claims for workloads. You can use the keys as [user claims in Vault role configurations](https://developer.hashicorp.com/vault/api-docs/auth/jwt#user_claim). The values are interpolated. For example, if you include the extra claim `unique_id = \"${job.region}:${job.namespace}:${job.id}\"`, you could set the user claim field to `/extra_claims/unique_id` to map that identifier to an entity alias. The available attributes for interpolation are:\n\n[`${job.region}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#job-region) - The region where the job is running.\n[`${job.namespace}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#job-namespace) - The job's namespace.\n[`${job.id}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#job-id) - The job's ID.\n[`${job.node_pool}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#job-node_pool) - The node pool where the allocation is running.\n[`${group.name}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#group-name) - The task group name of the task using Vault.\n[`${alloc.id}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#alloc-id) - The allocation's ID.\n[`${task.name}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#task-name) - The name of the task using Vault.\n[`${node.id}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#node-id) - The ID of the node where the allocation is running.\n[`${node.datacenter}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#node-datacenter) - The datacenter of the node where the allocation is running.\n[`${node.pool}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#node-pool) - The node pool of the node where the allocation is running.\n[`${node.class}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#node-class) - The class of the node where the allocation is running.\n[`${vault.cluster}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#vault-cluster) - The Vault cluster name.\n[`${vault.namespace}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#vault-namespace) - The Vault namespace.\n[`${vault.role}`](https://developer.hashicorp.com/nomad/docs/configuration/vault#vault-role) - The Vault role."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
					IsOptional: true,
				},
			},
		},
	},
}
