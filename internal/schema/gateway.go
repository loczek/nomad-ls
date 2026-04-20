package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var GatewaySchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"proxy": {
			Description: lang.Markdown("Configuration of the Envoy proxy that will be injected into the task group."),
			Body:        GatewayProxySchema,
		},
		"ingress": {
			Description: lang.Markdown("Configuration Entry of type `ingress-gateway` that will be associated with the service."),
			Body:        IngressSchema,
		},
		"terminating": {
			Description: lang.Markdown("Configuration Entry of type `terminating-gateway` that will be associated with the service."),
			Body:        TerminatingSchema,
		},
		"mesh": {
			Description: lang.Markdown("Indicates a mesh gateway will be associated with the service."),
			Body:        MeshSchema,
		},
	},
}

var GatewayProxySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"connect_timeout": {
			Description:  lang.Markdown("The amount of time to allow when making upstream connections before timing out. Defaults to 5 seconds. If the upstream service has the configuration option `[connect_timeout_ms](https://developer.hashicorp.com/consul/docs/connect/config-entries/service-resolver#connecttimeout)` set for the `service-resolver`, that timeout value will take precedence over this gateway proxy option."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"envoy_gateway_bind_tagged_addresses": {
			Description:  lang.Markdown("Indicates that the gateway services tagged addresses should be bound to listeners in addition to the default listener address."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		// TODO: fix constraints and add example
		"envoy_gateway_bind_addresses": {
			Description: lang.Markdown("A map of additional addresses to be bound. The keys to this map are the same of the listeners to be created and the values are a map with two keys - address and port, that combined make the address to bind the listener to. These are bound in addition to the default address. If `bridge` networking is in use, this map is automatically populated with additional listeners enabling the Envoy proxy to work from inside the network namespace."),
			Constraint:  schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"envoy_gateway_no_default_bind": {
			Description:  lang.Markdown("Prevents binding to the default address of the gateway service. This should be used with one of the other options to configure the gateway's bind addresses. If `bridge` networking is in use, this value will default to `true` since the Envoy proxy does not need to bind to the service address from inside the network namespace."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"envoy_dns_discovery_type": {
			Description: lang.Markdown("Determintes how Envoy will resolve hostnames. Defaults to `LOGICAL_DNS`. Must be one of `STRICT_DNS` or `LOGICAL_DNS`. Details for each type are available in the [Envoy Documentation](https://www.envoyproxy.io/docs/envoy/v1.16.1/intro/arch_overview/upstream/service_discovery). This option applies to terminating gateways that route to services addressed by a hostname."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: fix constraints
		"config": {
			Description: lang.Markdown("Escape hatch for [Advanced Configuration](https://developer.hashicorp.com/consul/docs/connect/proxies/envoy#advanced-configuration) of Envoy. Keys and values support [runtime variable interpolation](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation)."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
	},
}

var AddressSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"address": {
			Description: lang.Markdown("The address to bind to when combined with `port`."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"port": {
			Description: lang.Markdown("The port to listen to."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsRequired:  true,
		},
	},
}

var IngressSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"listener": {
			Description: lang.Markdown("One or more listeners that the ingress gateway should setup, uniquely identified by their port number."),
			Body:        ListenerSchema,
		},
		"tls": {
			Description: lang.Markdown("TLS configuration for this gateway."),
			Body:        TLSSchema,
		},
	},
}

var ListenerSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"port": {
			Description: lang.Markdown("The port that the listener should receive traffic on."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsRequired:  true,
		},
		// TODO: add warning from docs
		"protocol": {
			Description:  lang.Markdown("The protocol associated with the listener. One of `tcp`, `http`, `http2`, or `grpc`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("tcp")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"service": {
			Description: lang.Markdown("One or more services to be exposed via this listener. For `tcp` listeners, only a single service is allowed."),
			Body:        ListenerServiceSchema,
			MinItems:    1,
		},
	},
}

var TLSSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Set this configuration to enable TLS for every listener on the gateway. If TLS is enabled, then each `host` defined in the host field will be added as a DNSSAN to the gateway's x509 certificate."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"cipher_suites": {
			Description: lang.Markdown("Set the default list of TLS cipher suites for the gateway's listeners. Refer to [`CipherSuites`](https://developer.hashicorp.com/consul/docs/connect/config-entries/ingress-gateway#ciphersuites) in the Consul documentation for the supported cipher suites."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"tls_max_version": {
			Description: lang.Markdown("Set the default maximum TLS version supported by the gateway. Refer to [`TLSMaxVersion`](https://developer.hashicorp.com/consul/docs/connect/config-entries/ingress-gateway#tlsmaxversion) in the Consul documentation for supported versions."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"tls_min_version": {
			Description: lang.Markdown("Set the default minimum TLS version supported by the gateway. Refer to [`TLSMinVersion`](https://developer.hashicorp.com/consul/docs/connect/config-entries/ingress-gateway#tlsminversion) in the Consul documentation for supported versions."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"sds": {
			Description: lang.Markdown("Defines a set of parameters that configures the listener to load TLS certificates from an external Secret Discovery Service ([SDS](https://developer.hashicorp.com/consul/docs/connect/config-entries/ingress-gateway#listeners-services-tls-sds))."),
			Body:        SDSSchema,
		},
	},
}

var SDSSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cluster_name": {
			Description: lang.Markdown("The SDS cluster name to connect to to retrieve certificates."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"cert_resource": {
			Description: lang.Markdown("The SDS resource name to request when fetching the certificate from the SDS service."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
	},
}

var ListenerServiceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("The name of the service that should be exposed through this listener. This can be either a service registered in the catalog, or a service defined by other config entries, or a service that is going to be configured by Nomad. If the wildcard specifier `*` is provided, then ALL services will be exposed through this listener. This is not supported for a listener with protocol `tcp`."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"hosts": {
			Description: lang.Markdown("A list of hosts that specify what requests will match this service. This cannot be used with a `tcp` listener, and cannot be specified alongside a wildcard (`*`) service name. If not specified, the default domain `<service-name>.ingress.*` will be used to match services. Requests _must_ send the correct host to be routed to the defined service.\n\nThe wildcard specifier `*` can be used by itself to match all traffic coming to the ingress gateway, if TLS is not enabled. This allows a user to route all traffic to a single service without specifying a host, allowing simpler tests and demos. Otherwise, the wildcard specifier can be used as part of the host to match multiple hosts, but only in the leftmost DNS label. This ensures that all defined hosts are valid DNS records. For example, `*.example.com` is valid while `example.*` and `*-suffix.example.com` are not."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"max_concurrent_requests": {
			Description: lang.Markdown("Specifies the maximum number of concurrent HTTP/2 traffic requests that are allowed at a single point in time. If unset, will default to the Envoy proxy's default."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"max_connections": {
			Description: lang.Markdown("Specifies the maximum number of HTTP/1.1 connections a service instance is allowed to establish against the upstream. If unset, will default to the Envoy proxy's default."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"max_pending_requests": {
			Description: lang.Markdown("Specifies the maximum number of requests that are allowed to queue while waiting to establish a connection. If unset, will default to the Envoy proxy's default."),
			Constraint:  schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"request_headers": {
			Description: lang.Markdown("A set of HTTP-specific header modification rules that will be applied to requests routed to this service. This cannot be used with a tcp listener."),
			Body:        HeaderModifierSchema,
		},
		"response_headers": {
			Description: lang.Markdown("A set of HTTP-specific header modification rules that will be applied to responses from this service. This cannot be used with a tcp listener."),
			Body:        HeaderModifierSchema,
		},
		"tls": {
			Description: lang.Markdown("TLS configuration for this service."),
			Body:        TLSSchema,
		},
	},
}

var HeaderModifierSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"add": {
			Description: lang.Markdown("A set of key-value pairs to add to the headers, where header names are keys and header values are the values. Header names are not case-sensitive. If header values with the same name already exist, the value is appended and Consul applies both headers."),
			Constraint:  schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"set": {
			Description: lang.Markdown("A set of key-value pairs to add to the response header or to replace existing header values with. Use header names as the keys. Header names are not case-sensitive. If header values with the same names already exist, Consul replaces the header values."),
			Constraint:  schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"remove": {
			Description: lang.Markdown("Defines a list of headers to remove. Consul removes only headers containing exact matches. Header names are not case-sensitive."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
	},
}

var TerminatingSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"service": {
			Description: lang.Markdown("One or more services to be linked with the gateway. The gateway will proxy traffic to these services. These linked services must be registered with Consul for the gateway to discover their addresses. They must also be registered in the same Consul datacenter as the terminating gateway."),
			Body:        LinkedServiceSchema,
		},
	},
}

var LinkedServiceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description: lang.Markdown("The name of the service to link with the gateway. If the wildcard specifier `*` is provided, then ALL services within the Consul namespace wil lbe linked with the gateway."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"ca_file": {
			Description: lang.Markdown("A file path to a PEM-encoded certificate authority. The file must be accessible by the gateway task. The certificate authority is used to verify the authenticity of the service linked with the gateway. It can be provided along with a `cert_file` and `key_file` for mutual TLS authentication, or on its own for one-way TLS authentication. If none is provided the gateway **will not** encrypt traffic to the destination."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"cert_file": {
			Description: lang.Markdown("A file path to a PEM-encoded certificate. The file must be accessible by the gateway task. The certificate is provided to servers to verify the gateway's authenticity. It must be provided if a `key_file` is provided."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"key_file": {
			Description: lang.Markdown("A file path to a PEM-encoded private key. The file must be accessible by the gateway task. The key is used with the certificate to verify the gateway's authenticity. It must be provided if a `cert_file` is provided."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"sni": {
			Description: lang.Markdown("An optional hostname or domain name to specify during the TLS handshake."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}

// this schema is empty by design
var MeshSchema = &schema.BodySchema{}
