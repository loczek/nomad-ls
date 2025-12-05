package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ConnectSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"native": {
			Description:  lang.Markdown("This is used to configure the service as supporting [Consul service mesh native](https://developer.hashicorp.com/consul/docs/automate/native) applications."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"sidecar_service": {
			Description: lang.PlainText("This is used to configure the sidecar service created by Nomad for Consul service mesh."),
			Body:        SidecarServiceSchema,
		},
		"sidecar_task": {
			Description: lang.PlainText("This modifies the task configuration of the Envoy proxy created as a sidecar or gateway."),
			Body:        SidecarTaskSchema,
		},
		"gateway": {
			Description: lang.PlainText("This is used to configure the gateway service created by Nomad for Consul service mesh."),
		},
	},
}
