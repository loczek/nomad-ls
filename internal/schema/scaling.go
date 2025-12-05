package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ScalingSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"min": {
			Description: lang.Markdown("The minimum acceptable count for the task group. This should be honored by the external autoscaler. It will also be honored by Nomad during job updates and scaling operations. Defaults to the specified task group [`count`](https://developer.hashicorp.com/nomad/docs/job-specification/group#count)."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"max": {
			Description: lang.Markdown("The maximum acceptable count for the task group. This should be honored by the external autoscaler. It will also be honored by Nomad during job updates and scaling operations."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"enabled": {
			Description:  lang.Markdown("Whether the scaling policy is enabled. This is intended to allow temporarily disabling an autoscaling policy, and should be honored by the external autoscaler."),
			DefaultValue: &schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   &schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"policy": {
			Description: lang.Markdown("The autoscaling policy. This is opaque to Nomad, consumed and parsed only by the external autoscaler. Therefore, its contents are specific to the autoscaler; consult the [Nomad Autoscaler documentation](https://developer.hashicorp.com/nomad/tools/autoscaling/policy) for more details."),
			Type:        schema.BlockTypeMap,
		},
	},
}
