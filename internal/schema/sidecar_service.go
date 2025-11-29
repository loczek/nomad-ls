package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var SidecarServiceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"disable_default_tcp_check": {
			Description:  lang.Markdown("disable the default TCP health check."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"meta": {
			Description: lang.Markdown("Specifies arbitrary KV metadata pairs."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		// TODO: docs don't have default value?
		"port": {
			Description:  lang.Markdown("Port label for sidecar service."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tags": {
			Description: lang.Markdown("Custom Consul service tags for the sidecar service."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"proxy": {
			Description: lang.Markdown("This is used to configure the sidecar proxy service."),
			Body:        ProxySchema,
		},
	},
}

var ProxySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"config": {
			Description: lang.Markdown("Proxy configuration that is opaque to Nomad and passed directly to Consul. See [Consul service mesh documentation](/consul/docs/connect/proxies/envoy#dynamic-configuration) for details. Keys and values support [runtime variable interpolation](/nomad/docs/reference/runtime-variable-interpolation)."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"meta": {
			Description: lang.Markdown("Specifies arbitrary KV metadata pairs."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		// TODO: docs don't have default value?
		"port": {
			Description:  lang.Markdown("Port label for sidecar service."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"local_service_address": {
			Description:  lang.Markdown("The address the local service binds to. Useful to customize in clusters with mixed Connect and non-Connect services."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("127.0.0.1")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"local_service_port": {
			Description: lang.Markdown("The port the local service binds to. Usually the same as the parent service's port, it is useful to customize in clusters with mixed Connect and non-Connect services."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		// TODO: this can also be set to `true` or the `ExposeSchema`
		"expose": {
			Description: lang.Markdown("Used to configure expose path configuration for Envoy. See Consul's [Expose Paths Configuration Reference](/consul/docs/connect/proxies/proxy-config-reference#expose-paths-configuration-reference) for more information."),
			Body:        ExposeSchema,
		},
		"transparent_proxy": {
			Description: lang.Markdown("Used to enable [transparent proxy](/consul/docs/k8s/connect/transparent-proxy) mode, which allows the proxy to use Consul service intentions to automatically configure upstreams, and configures iptables rules to force traffic from the allocation to flow through the proxy."),
			Body:        TransparentProxySchema,
		},
		"upstreams": {
			Description: lang.Markdown("Used to configure details of each upstream service that this sidecar proxy communicates with."),
			Body:        UpstreamsSchema,
		},
	},
}
var ExposeSchema = &schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"path": {
			Description: lang.Markdown("A list of Envoy expose path configurations to expose through Envoy."),
			Body:        PathSchema,
		},
	},
}

var PathSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"path": {
			Description: lang.Markdown("The HTTP or gRPC path to expose. The path must be prefixed with a slash."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"protocol": {
			Description: lang.Markdown("Sets the protocol of the listener. Must be `http` or `http2`. For gRPC use `http2`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"local_path_port": {
			Description: lang.Markdown("The port the service is listening to for connections to the configured `path`. Typically this will be the same as the `service.port` value, but could be different if for example the exposed path is intended to resolve to another task in the task group."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"listener_port": {
			Description: lang.Markdown("The name of the port to use for the exposed listener. The port should be configured to [map inside](/nomad/docs/job-specification/network#to) the task's network namespace."),
			Body:        PortSchema,
		},
	},
}

var TransparentProxySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"exclude_inbound_ports": {
			Description: lang.Markdown("A list of inbound ports to exclude from the inbound traffic redirection. This allows traffic on these ports to bypass the Envoy proxy. These ports can be specified as either [network port labels](/nomad/docs/job-specification/network#port-parameters) or as numeric ports. Nomad will automatically add the following to this list:\n\n- The [`local_path_port`](/nomad/docs/job-specification/expose#local_path_port) of any [`expose`](/nomad/docs/job-specification/expose) block.\n- The port of any service check with [`expose=true`](/nomad/docs/job-specification/check#expose) set.\n- The port of any `network.port` with a [`static`](/nomad/docs/job-specification/network#static) value."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsRequired:  true,
		},
		"exclude_outbound_cidrs": {
			Description: lang.Markdown("A list of CIDR subnets that should be excluded from outbound traffic redirection. This allows traffic to these subnets to bypass the Envoy proxy. Note this is independent of `exclude_outbound_ports`; CIDR subnets listed here are excluded regardless of the port."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
		"exclude_outbound_ports": {
			Description: lang.Markdown("A list of port numbers that should be excluded from outbound traffic redirection. This allows traffic to these subnets to bypass the Envoy proxy. Note this is independent of `exclude_outbound_cidrs`; ports listed here are excluded regardless of the CIDR."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.Number)},
			IsOptional:  true,
		},
		"exclude_uids": {
			Description: lang.Markdown("A list of Unix user IDs (UIDs) that should be excluded from outbound traffic redirection. When unset, only the Envoy proxy's user will be allowed to bypass the iptables rule."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"no_dns": {
			Description:  lang.Markdown("By default, Consul will be set as the nameserver for the workload and IP tables rules will redirect DNS queries to Consul. If you want only external DNS, set `no_dns=true`. You will need to add your own CIDR and port exclusions for your DNS nameserver. You cannot set [`network.dns`](/nomad/docs/job-specification/network#dns-parameters) if `no_dns=false`."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"outbound_port": {
			Description:  lang.Markdown("The port that Envoy will bind on inside the network namespace. The iptables rules created by `consul-cni` will force traffic to flow to this port. You should only set this value if you have specifically set the [`outbound_listener_port`](/consul/docs/connect/proxies/proxy-config-reference#outbound_listener_port) in your Consul proxy configuration. You can change the default value for a given node via [client metadata](/nomad/docs/job-specification/transparent_proxy#client-metadata) (see below)."),
			DefaultValue: &schema.DefaultValue{Value: cty.NumberIntVal(15001)},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"uid": {
			Description:  lang.Markdown("The Unix user ID (UID) used by the Envoy proxy. You should only set this value if you have a custom build of the Envoy container image which uses a different UID. You can change the default value for a given node via [client metadata](/nomad/docs/job-specification/transparent_proxy#client-metadata) (see below). Note that your workload's task cannot use the same UID as the Envoy sidecar proxy."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("101")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"listener_port": {
			Description: lang.Markdown("The name of the port to use for the exposed listener. The port should be configured to [map inside](/nomad/docs/job-specification/network#to) the task's network namespace."),
			Body:        PortSchema,
		},
	},
}

var UpstreamsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"config": {
			Description: lang.Markdown("Upstream configuration that is opaque to Nomad and passed directly to Consul. See [Consul service mesh documentation](/consul/docs/connect/proxies/proxy-config-reference#expose-paths-configuration-reference) for details. Keys and values support [runtime variable interpolation](/nomad/docs/reference/runtime-variable-interpolation)."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"destination_name": {
			Description: lang.Markdown("Name of the upstream service."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"destination_namespace": {
			Description: lang.Markdown("Name of the upstream Consul namespace."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"destination_partition": {
			Description:  lang.Markdown("Name of the Cluster admin partition containing the upstream service."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
		},
		"destination_peer": {
			Description:  lang.Markdown("Name of the peer cluster containing the upstream service."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
		},
		"destination_type": {
			Description:  lang.Markdown("The type of discovery query the proxy should use for finding service mesh instances."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("service")},
			Constraint:   &schema.LiteralType{Type: cty.String},
		},
		"local_bind_port": {
			Description: lang.Markdown("The port the proxy will receive connections for the upstream on."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsRequired:  true,
		},
		"datacenter": {
			Description:  lang.Markdown("The Consul datacenter in which to issue the discovery query. Defaults to the empty string, which Consul interprets as the local Consul datacenter."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"local_bind_address": {
			Description:  lang.Markdown("The address the proxy will receive connections for the upstream on."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"local_bind_socket_mode": {
			Description:  lang.Markdown("Unix octal that configures file permissions for the socket."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"local_bind_socket_path": {
			Description:  lang.Markdown("The path at which to bind a Unix domain socket listener."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"mesh_gateway": {
			Description: lang.Markdown("Configures the mesh gateway behavior for connecting to this upstream."),
			Body:        MeshGatewaySchema,
		},
	},
}

var MeshGatewaySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"mode": {
			Description:  lang.Markdown("The mode of operation in which to use [Connect Mesh Gateways](/consul/docs/connect/gateways/mesh-gateway/service-to-service-traffic-datacenters#mesh-gateways). If left unset, the mode will default to the mode as determined by the Consul [service-defaults](/consul/docs/connect/config-entries/service-defaults#meshgateway) configuration for the service. Can be configured with the following modes:\n\n - [`local`](/nomad/docs/job-specification/upstreams#local) - In this mode the Connect proxy makes its outbound connection to a gateway running in the same datacenter. That gateway is then responsible for ensuring the data gets forwarded along to gateways in the destination datacenter.\n - [`remote`](/nomad/docs/job-specification/upstreams#remote) - In this mode the Connect proxy makes its outbound connection to a gateway running in the destination datacenter. That gateway will then forward the data to the final destination service.\n - [`none`](/nomad/docs/job-specification/upstreams#none) - In this mode, no gateway is used and a Connect proxy makes its outbound connections directly to the destination services."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
