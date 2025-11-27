package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var NetworkSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"mbits": {
			Description: lang.Markdown("Specifies the bandwidth required in MBits."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(10),
			},
			Constraint:   &schema.LiteralType{Type: cty.Number},
			IsDeprecated: true,
		},
		// TODO: update docs
		"mode": {
			Description: lang.Markdown(""),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("host"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
		"hostname": {
			Description: lang.Markdown("The hostname assigned to the network namespace. This is currently only supported using the [Docker driver](/nomad/docs/job-declare/task-driver/docker) and when the [mode](/nomad/docs/job-specification/network#mode) is set to [`bridge`](/nomad/docs/job-specification/network#bridge). This parameter supports [interpolation](/nomad/docs/reference/runtime-variable-interpolation)."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsDeprecated: true,
		},
		"image_pull_timeout": {
			Description: lang.Markdown("A time duration that controls how long Nomad will wait before cancelling an in-progress pull of the Docker image as specified in `image`. Defaults to `\"5m\"`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("5m"),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"port": {
			Description: lang.Markdown("Specifies a TCP/UDP port allocation and can be used to specify both dynamic ports and reserved ports."),
			Body:        PortSchema,
			Labels: []*schema.LabelSchema{
				{
					Name: "name",
				},
			},
		},
		"dns": {
			Description: lang.Markdown("Sets the DNS configuration for the allocations. By default all task drivers will inherit DNS configuration from the client host. DNS configuration is only supported on Linux clients at this time. Note that if you are using a `mode=\"cni/*`, these values will override any DNS configuration the CNI plugins return."),
			Body:        DnsSchema,
		},
		"cni": {
			Description: lang.Markdown("Sets the custom CNI arguments for a network configuration per allocation, for use with `mode=\"cni/*`."),
			Body:        CniSchema,
		},
	},
}

var PortSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"static": {
			Description: lang.Markdown("Specifies the static TCP/UDP port to allocate. If omitted, a dynamic port is chosen. We do not recommend using static ports, except for system or specialized jobs like load balancers."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"to": {
			Description: lang.Markdown("Applicable when using \"bridge\" mode to configure port to map to inside the task's network namespace. Omitting this field or setting it to `-1` sets the mapped port equal to the dynamic port allocated by the scheduler. The `NOMAD_PORT_<label>` environment variable will contain the `to` value."),
			Constraint:  &schema.LiteralType{Type: cty.Number},
			IsOptional:  true,
		},
		"host_network": {
			Description: lang.Markdown("Designates the host network name to use when allocating the port. When port mapping the host port will only forward traffic to the matched host network address."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"ignore_collision": {
			Description: lang.Markdown("A time duration that controls how long Nomad will wait before cancelling an in-progress pull of the Docker image as specified in `image`. Defaults to `\"5m\"`."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: &schema.LiteralType{Type: cty.Bool},
			IsOptional: true,
		},
	},
}

var DnsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"server": {
			Description: lang.Markdown("Sets the DNS nameservers the allocation uses for name resolution."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
		},
		"searches": {
			Description: lang.Markdown("Sets the search list for hostname lookup"),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
		},
		"options": {
			Description: lang.Markdown("Sets internal resolver variables."),
			Constraint:  &schema.LiteralType{Type: cty.List(cty.String)},
		},
	},
}

var CniSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"args": {
			Description: lang.Markdown("Sets CNI arguments for network configuration. These get turned into `CNI_ARGS` per the [CNI spec](https://www.cni.dev/docs/spec/#parameters)."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
		},
	},
}
