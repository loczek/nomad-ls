package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ArtifactSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"destination": {
			Description: lang.Markdown("Specifies the directory path to download the artifact, relative to the root of the task's working directory. If omitted, the default value is to place the artifact in local/. The destination is treated as a directory unless mode is set to file. Source files will be downloaded into that directory path. For more details on how the destination interacts with task drivers, see the Filesystem internals documentation."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("local/"),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		"mode": {
			Description: lang.Markdown("One of `any`, `file`, or `dir`. If set to `file` the destination must be a file, not a directory. By default the `destination` will be `local/<filename>`."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal("any"),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
		// TODO: change the default value type?
		"options": {
			Description: lang.Markdown("Specifies configuration parameters to fetch the artifact. The key-value pairs map directly to parameters appended to the supplied `source` URL. Please see the [`go-getter` documentation](https://github.com/hashicorp/go-getter) for a complete list of options and examples."),
			DefaultValue: schema.DefaultValue{
				Value: cty.MapValEmpty(cty.String),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Map(cty.String)},
				schema.AnyExpression{OfType: cty.Map(cty.String)},
			},
			IsOptional: true,
		},
		"headers": {
			Description: lang.Markdown("Specifies HTTP headers to set when fetching the artifact using `http` or `https` protocol. Please see the [`go-getter` headers documentation](https://github.com/hashicorp/go-getter#headers) for more information."),
			DefaultValue: schema.DefaultValue{
				Value: cty.MapValEmpty(cty.String),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.Map(cty.String)},
				schema.AnyExpression{OfType: cty.Map(cty.String)},
			},
			IsOptional: true,
		},
		"source": {
			Description: lang.Markdown("Specifies the URL of the artifact to download. See [`go-getter`](https://github.com/hashicorp/go-getter) for details."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsRequired: true,
		},
		"chown": {
			Description: lang.Markdown("Specifies whether Nomad should recursively `chown` the downloaded artifact to be owned by the [`task.user`](https://developer.hashicorp.com/nomad/docs/job-specification/task#user) uid and gid."),
			DefaultValue: schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: schema.OneOf{
				schema.LiteralType{Type: cty.String},
				schema.AnyExpression{OfType: cty.String},
			},
			IsOptional: true,
		},
	},
}
