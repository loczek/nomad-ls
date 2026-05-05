package plugin

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var RawExecSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Specifies whether the driver should be enabled or disabled. Defaults to `false`."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		// TODO: add example
		"denied_host_uids": {
			Description: lang.Markdown("Specifies a comma-separated list of host uids to deny. Ranges can be specified by using a hyphen separating the two inclusive ends. If a \"user\" value is specified in task configuration and that user has a user id in the given ranges, the task will error before starting. This will not be checked on Windows clients."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add example
		"denied_host_gids": {
			Description: lang.Markdown("Specifies a comma-separated list of host gids to deny. Ranges can be specified by using a hyphen separating the two inclusive ends. If a \"user\" value is specified in task configuration and that user is part of any groups with gid's in the specified ranges, the task will error before starting. This will not be checked on Windows clients."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		// TODO: add example
		"denied_envvars": {
			Description: lang.Markdown("Passes a list of environment variables that the driver should scrub from all task environments. Supports globbing with \"*\" wildcard accepted as prefix and/or suffix."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
	},
}
