package job

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var LogsSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"max_files": {
			Description: lang.Markdown("Specifies the maximum number of rotated files Nomad will retain for `stdout` and `stderr`. Each stream is tracked individually, so specifying a value of 2 will create 4 files - 2 for stdout and 2 for stderr"),
			DefaultValue: schema.DefaultValue{
				Value: cty.NumberIntVal(10),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
		"max_file_size": {
			Description: lang.Markdown("Specifies the maximum size of each rotated file in `MB`. If the amount of disk resource requested for the task is less than the total amount of disk space needed to retain the rotated set of files, Nomad will return a validation error when a job is submitted."),
			DefaultValue: schema.DefaultValue{
				Value: cty.NumberIntVal(10),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Number},
				schema.AnyExpression{OfType: cty.Number},
			},
			IsOptional: true,
		},
		"disabled": {
			Description: lang.Markdown("Specifies that log collection should be enabled for this task. If set to `true`, the task driver will attach stdout/stderr of the task to `/dev/null` (or `NUL` on Windows). You should only disable log collection if your application has some other way of emitting logs, such as writing to a remote syslog server. Note that the `nomad alloc logs` command and related APIs will return errors (404 \"not found\") if logging is disabled."),
			DefaultValue: schema.DefaultValue{
				Value: cty.BoolVal(false),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Bool},
				schema.AnyExpression{OfType: cty.Bool},
			},
			IsOptional: true,
		},
	},
}
