package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var LimitsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"https_handshake_timeout": {
			Description:  lang.Markdown("Configures the limit for how long the HTTPS server in both client and server agents will wait for a client to complete a TLS handshake. This should be kept conservative as it limits how many connections an unauthenticated attacker can open if [`tls.http = true`](https://developer.hashicorp.com/nomad/docs/configuration/tls) is being used (strongly recommended in production). Default value is `5s`. `0` disables HTTP handshake timeouts."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"http_max_conns_per_client": {
			Description:  lang.Markdown("Configures a limit of how many concurrent TCP connections a single client IP address is allowed to open to the agent's HTTP server. This affects the HTTP servers in both client and server agents. Default value is `100`. `0` disables HTTP connection limits."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(100)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"rpc_handshake_timeout": {
			Description:  lang.Markdown("Configures the limit for how long servers will wait after a client TCP connection is established before they complete the connection handshake. When TLS is used, the same timeout applies to the TLS handshake separately from the initial protocol negotiation. All Nomad clients should perform this immediately on establishing a new connection. This should be kept conservative as it limits how many connections an unauthenticated attacker can open if TLS is being using to authenticate clients (strongly recommended in production). When `tls.rpc` is true on servers, this limits how long the connection and associated goroutines will be held open before the client successfully authenticates. Default value is `5s`. `0` disables RPC handshake timeouts."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"rpc_max_conns_per_client": {
			Description:  lang.Markdown("Configures a limit of how many concurrent TCP connections a single source IP address is allowed to open to a single server. Client agents do not accept RPC TCP connections directly and therefore are not affected. It affects both clients connections and other server connections. Nomad clients multiplex many RPC calls over a single TCP connection, except for streaming endpoints such as [log streaming](https://developer.hashicorp.com/nomad/api-docs/client#stream-logs) which require their own connection when routed through servers. A server needs at least 2 TCP connections (1 Raft, 1 RPC) per peer server locally and in any federated region. Servers also need a TCP connection per routed streaming endpoint concurrently in use. Only operators use streaming endpoints; as of 0.10.3 Nomad client code does not. A reasonably low limit significantly reduces the ability of an unauthenticated attacker to consume unbounded resources by holding open many connections. You may need to increase this if WAN federated servers connect via proxies or NAT gateways or similar causing many legitimate connections from a single source IP. Default value is `100` which is designed to support the majority of users. `0` disables RPC connection limits. `26` is the minimum as `20` connections are always reserved for non-streaming connections (Raft and RPC) to ensure streaming RPCs do not prevent normal server operation. This minimum may be lowered in the future when streaming RPCs no longer require their own TCP connection."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(100)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
	},
}
