package schemaACL

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var NamespaceSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"policy": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				ReadPolicy,
				WritePolicy,
				DenyPolicy,
				ScalePolicy,
			},
			IsRequired: true,
		},
		"capabilities": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				DenyCapability,
				ListJobsCapability,
				ParseJobCapability,
				ReadJobCapability,
				SubmitJobCapability,
				DispatchJobCapability,
				ReadLogsCapability,
				ReadFsCapability,
				AllocExecCapability,
				AllocNodeExecCapability,
				AllocLifecycleCapability,
				CsiRegisterPluginCapability,
				CsiWriteVolumeCapability,
				CsiReadVolumeCapability,
				CsiListVolumeCapability,
				CsiMountVolumeCapability,
				HostVolumeCreateCapability,
				HostVolumeDeleteCapability,
				HostVolumeReadCapability,
				HostVolumeRegisterCapability,
				HostVolumeWriteCapability,
				ListScalingPoliciesCapability,
				ReadScalingPolicyCapability,
				ReadJobScalingCapability,
				ScaleJobCapability,
				SentinelOverrideCapability,
				SubmitRecommendationCapability,
			},
			IsRequired: true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"variables": {
			Body:     VariableSchema,
			MaxItems: 1,
		},
	},
}

var VariableSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"policy": {
			Description: lang.Markdown("The Nomad region that the limit applies to."),
			Constraint: schema.OneOf{
				ReadPolicy,
				WritePolicy,
				DenyPolicy,
				ScalePolicy,
			},
			IsRequired: true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"path": {
			Body: VariablePathSchema,
		},
	},
}

var VariablePathWritePolicy = schema.LiteralValue{
	Value:       cty.StringVal("write"),
	Description: lang.Markdown("Create or update Variables at this path. Includes the \"list\" capability but not the \"read\" or \"destroy\" capabilities."),
}
var VariablePathReadPolicy = schema.LiteralValue{
	Value:       cty.StringVal("read"),
	Description: lang.Markdown("Read the decrypted contents of Variables at this path. Also includes the \"list\" capability"),
}
var VariablePathListPolicy = schema.LiteralValue{
	Value:       cty.StringVal("list"),
	Description: lang.Markdown("List the metadata but not contents of Variables at this path."),
}
var VariablePathDestroyPolicy = schema.LiteralValue{
	Value:       cty.StringVal("destroy"),
	Description: lang.Markdown("Delete Variables at this path."),
}
var VariablePathDenyPolicy = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("No permissions at this path. Deny takes precedence over other capabilities."),
}

var VariablePathSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"capabilities": {
			Constraint: schema.OneOf{
				VariablePathWritePolicy,
				VariablePathReadPolicy,
				VariablePathListPolicy,
				VariablePathDestroyPolicy,
				VariablePathDenyPolicy,
			},
			IsRequired: true,
		},
	},
}
