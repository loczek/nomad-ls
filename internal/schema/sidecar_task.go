package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var SidecarTaskSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"name": {
			Description:  lang.Markdown("Name of the task. Defaults to including the name of the service the proxy or gateway is providing."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("connect-[proxy|gateway]-<service>")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"driver": {
			Description:  lang.Markdown("Driver used for the sidecar task."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("docker")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"user": {
			Description: lang.Markdown("Determines which user is used to run the task, defaults to the same user the Nomad client is being run as."),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"config": {
			Description: lang.Markdown("Configuration provided to the driver for initialization. Keys and values support [runtime variable interpolation](/nomad/docs/reference/runtime-variable-interpolation)."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"env": {
			Description: lang.Markdown("Map of environment variables used by the driver."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"meta": {
			Description: lang.Markdown("Arbitrary metadata associated with this task that's opaque to Nomad."),
			Constraint:  &schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"kill_timeout": {
			Description:  lang.Markdown("Time between signalling a task that will be killed and killing it."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"shutdown_delay": {
			Description:  lang.Markdown("Delay between deregistering the task from Consul and sending it a signal to shutdown."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("5s")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"kill_signal": {
			Description:  lang.Markdown("Kill signal to use for the task, defaults to SIGINT."),
			DefaultValue: &schema.DefaultValue{Value: cty.StringVal("SIGINT")},
			Constraint:   &schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"resources": {
			Description: lang.Markdown("Resources needed by the sidecar task."),
			Body:        ResourcesSchema,
		},
		"logs": {
			Description: lang.Markdown("Specifies logging configuration for the `stdout` and `stderr` of the task."),
			Body:        LogsSchema,
		},
		"volume_mount": {
			Description: lang.Markdown("Specifies where a group volume should be mounted."),
			Body:        VolumeMountSchema,
		},
	},
}
