package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ServerJoinSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		// TODO: add examples
		"retry_join": {
			Description:  lang.Markdown("Specifies a list of Nomad server addresses and [Cloud Auto-join](https://developer.hashicorp.com/nomad/docs/configuration/server_join#cloud-auto-join) configurations that are joined as cluster members. This is similar to [`start_join`](https://developer.hashicorp.com/nomad/docs/configuration/server_join#start_join), but join attempts are retried up to [retry_max](https://developer.hashicorp.com/nomad/docs/configuration/server_join#retry_max) times. Further, `retry_join` is available to both Nomad servers and clients, while `start_join` is only defined for Nomad servers. This is useful for cases where we know the address will become available eventually. Use `retry_join` with an array as a replacement for `start_join`, do not use both options."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"retry_interval": {
			Description:  lang.Markdown("Specifies the time to wait between retry join attempts."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"retry_max": {
			Description:  lang.Markdown("Specifies the maximum number of join attempts to be made before exiting with a return code of 1. By default, this is set to 0 which is interpreted as infinite retries."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"start_join": {
			Description:  lang.Markdown("Specifies a list of server addresses to join on startup. If Nomad is unable to join with any of the specified addresses, agent startup will fail. See the [server address format](https://developer.hashicorp.com/nomad/docs/configuration/server_join#server-address-format) section for more information on the format of the string. This field is defined only for Nomad servers and will result in a configuration parse error if included in a client configuration."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}
