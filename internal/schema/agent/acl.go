package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ACLSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Specifies if ACL enforcement is enabled. All other ACL configuration options depend on this value. All agents should have the same value for this parameter. For example the Nomad command line will send requests for client endpoints such as `alloc exec` directly to Nomad clients whenever they are accessible. In this scenario, the client will enforce ACLs, so both servers and clients should have ACLs enabled."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"token_ttl": {
			Description:  lang.Markdown("Specifies the maximum time-to-live (TTL) for cached ACL tokens. This does not affect servers, since they do not cache tokens. Setting this value lower reduces how stale a token can be, but increases the request load against servers. If a client cannot reach a server, for example because of an outage, the TTL will be ignored and the cached value used."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"policy_ttl": {
			Description:  lang.Markdown("Specifies the maximum time-to-live (TTL) for cached ACL policies. This does not affect servers, since they do not cache policies. Setting this value lower reduces how stale a policy can be, but increases the request load against servers. If a client cannot reach a server, for example because of an outage, the TTL will be ignored and the cached value used."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"role_ttl": {
			Description:  lang.Markdown("Specifies the maximum time-to-live (TTL) for cached ACL roles. This does not affect servers, since they do not cache roles. Setting this value lower reduces how stale a role can be, but increases the request load against servers. If a client cannot reach a server, for example because of an outage, the TTL will be ignored and the cached value used."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"replication_token": {
			Description:  lang.Markdown("Specifies the Secret ID of the ACL token to use for replicating policies and tokens. This is used by servers in non-authoritative region to mirror the policies and tokens into the local region from the [`authoritative_region`](https://developer.hashicorp.com/nomad/docs/configuration/server#authoritative_region). Setting `replication_token` requires that ACLs have been bootstrapped in the authoritative region. Refer to [Configure for multiple regions](https://developer.hashicorp.com/nomad/docs/secure/acl/bootstrap#configure-for-multiple-regions) in the ACLs tutorial."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"token_min_expiration_ttl": {
			Description:  lang.Markdown("Specifies the lowest acceptable TTL value for an ACL token when setting expiration. This is used by the Nomad servers to validate ACL tokens and ACL authentication methods."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("1m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"token_max_expiration_ttl": {
			Description:  lang.Markdown("Specifies the highest acceptable TTL value for an ACL token when setting expiration. This is used by the Nomad servers to validate ACL tokens and ACL authentication methods."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("1m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
