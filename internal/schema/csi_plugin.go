package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var CsiPluginSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"id": {
			Description: lang.Markdown("This is the ID for the plugin. Some plugins will require both controller and node plugin types (see below); you need to use the same ID for both so that Nomad knows they belong to the same plugin."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"type": {
			Description: lang.Markdown("One of `node`, `controller`, or `monolith`. Each plugin supports one or more types. Each Nomad client node where you want to mount a volume will need a `node` plugin instance. Some plugins will also require one or more `controller` plugin instances to communicate with the storage provider's APIs. Some plugins can serve as both `controller` and `node` at the same time, and these are called `monolith` plugins. Refer to your CSI plugin's documentation."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
		"mount_dir": {
			Description: lang.Markdown("The directory path inside the container where the plugin will expect a Unix domain socket for bidirectional communication with Nomad. This field is typically not required. Refer to your CSI plugin's documentation for details."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"stage_publish_base_dir": {
			Description: lang.Markdown("The base directory path inside the container where the plugin will be instructed to stage and publish volumes. This field is typically not required. Refer to your CSI plugin's documentation for details. This can not be a subdirectory of `mount_dir`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"health_timeout": {
			Description:  lang.Markdown("The duration that the plugin supervisor will wait before restarting an unhealthy CSI plugin. Must be a duration value such as `30s` or `2m`. Defaults to `30s` if not set."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
