package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var AutopilotSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cleanup_dead_servers": {
			Description:  lang.Markdown("Specifies automatic removal of dead server nodes periodically and whenever a new server is added to the cluster."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"last_contact_threshold": {
			Description:  lang.Markdown("Specifies the maximum amount of time a server can go without contact from the leader before being considered unhealthy. Must be a duration value such as `10s`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("200ms")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"max_trailing_logs": {
			Description:  lang.Markdown("Specifies the maximum number of log entries that a server can trail the leader by before being considered unhealthy."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(250)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"server_stabilization_time": {
			Description:  lang.Markdown("Specifies the minimum amount of time a server must be stable in the 'healthy' state before being added to the cluster. Only takes effect if all servers are running Raft protocol version 3 or higher. Must be a duration value such as `30s`."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("10s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"enable_redundancy_zones": {
			Description:  lang.Markdown("Controls whether Autopilot separates servers into zones for redundancy, in conjunction with the redundancy_zone parameter. Only one server in each zone can be a voting member at one time." + schemautils.Divider + schemautils.EnterpriseOnly),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"disable_upgrade_migration": {
			Description:  lang.Markdown("Disables Autopilot's upgrade migration strategy in Nomad Enterprise of waiting until enough newer-versioned servers have been added to the cluster before promoting any of them to voters." + schemautils.Divider + schemautils.EnterpriseOnly),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"enable_custom_upgrades": {
			Description:  lang.Markdown("Specifies whether to enable using custom upgrade versions when performing migrations, in conjunction with the upgrade_version parameter." + schemautils.Divider + schemautils.EnterpriseOnly),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
