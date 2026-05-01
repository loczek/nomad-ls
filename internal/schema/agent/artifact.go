package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ArtifactSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"http_read_timeout": {
			Description:  lang.Markdown("Specifies the maximum duration in which an HTTP download request must complete before it is canceled. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"http_max_size": {
			Description:  lang.Markdown("Specifies the maximum size allowed for artifacts downloaded via HTTP. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("100GB")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"gcs_timeout": {
			Description:  lang.Markdown("Specifies the maximum duration in which a Google Cloud Storate operation must complete before it is canceled. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"git_timeout": {
			Description:  lang.Markdown("Specifies the maximum duration in which a Git operation must complete before it is canceled. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"hg_timeout": {
			Description:  lang.Markdown("Specifies the maximum duration in which a Mercurial operation must complete before it is canceled. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"s3_timeout": {
			Description:  lang.Markdown("Specifies the maximum duration in which an S3 operation must complete before it is canceled. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"decompression_size_limit": {
			Description:  lang.Markdown("Specifies the maximum amount of data that will be decompressed before triggering an error and cancelling the operation. Set to \"0\" to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("100GB")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"decompression_file_count_limit": {
			Description:  lang.Markdown("Specifies the maximum number of files that will be decompressed before triggering an error and cancelling the operation. Set to `0` to not enforce a limit."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(4096)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"disable_artifact_inspection": {
			Description:  lang.Markdown("Specifies whether to disable artifact inspection for sandbox escapes. If the platform supports filesystem isolation, and it is not disabled, artifact inspection will not be performed regardless of this value."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"disable_filesystem_isolation": {
			Description:  lang.Markdown("Specifies whether filesystem isolation should be disabled for artifact downloads. Applies only to systems where filesystem isolation via [landlock](https://docs.kernel.org/userspace-api/landlock.html) is possible (Linux kernel 5.13+)."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"filesystem_isolation_extra_paths": {
			Description:  lang.Markdown("Allow extra paths in the filesystem isolation. Paths are specified in the form `[kind]:[mode]:[path]` where `kind` must be either `f` or `d` (file or directory) and mode must be zero or more of `r`, `w`, `c`, `x` (read, write, create, execute) e.g. `f:r:/dev/urandom` would enable reading the /dev/urandom file, `d:rx:/opt/bin` would enable reading and executing from the /opt/bin directory"),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"set_environment_variables": {
			Description:  lang.Markdown("Specifies a comma separated list of environment variables that should be inherited by the artifact sandbox from the Nomad client's environment. By default a minimal environment is set including a `PATH` appropriate for the operating system."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
