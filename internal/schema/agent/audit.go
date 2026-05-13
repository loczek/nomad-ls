package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var AuditSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Specifies if audit logging should be enabled. When enabled, audit logging will occur for every request, unless it is filtered by a filter."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"sink": {
			Description: SinkSchema.Description,
			Body:        SinkSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
		"filter": {
			Description: FilterSchema.Description,
			Body:        FilterSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
	},
}

var SinkSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description:  lang.Markdown("Specifies the type of sink to create. Currently only `\"file\"` type is supported."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("file")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"delivery_guarantee": {
			Description:  lang.Markdown("Specifies the delivery guarantee that will be made for each audit log entry. Available options are `\"enforced\"` and `\"best-effort\"`. `\"enforced\"` will halt request execution if the audit log event fails to be written to its sink. `\"best-effort\"` will not halt request execution, meaning a request could potentially be un-audited."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("enforced")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"format": {
			Description:  lang.Markdown("Specifies the output format to be sent to a sink. Currently only `\"json\"` format is supported."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("json")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"mode": {
			Description:  lang.Markdown("Specifies the permissions mode for the audit log files using octal notation."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("0600")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"path": {
			Description:  lang.Markdown("Specifies the path and file name to use for the audit log. By default Nomad will use its configured [`data_dir`](https://developer.hashicorp.com/nomad/docs/configuration#data_dir) for a combined path of `/data_dir/audit/audit.log`. If `rotate_bytes` or `rotate_duration` are set file rotation will occur. In this case the filename will be post-fixed with a timestamp `\"filename-{timestamp}.log\"`"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("[data_dir]/audit/audit.log")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"rotate_bytes": {
			Description:  lang.Markdown("Specifies the number of bytes that should be written to an audit log before it needs to be rotated. Unless specified, there is no limit to the number of bytes that can be written to a log file."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"rotate_duration": {
			Description:  lang.Markdown("Specifies the maximum duration a audit log should be written to before it needs to be rotated. Must be a duration value such as 30s."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("24h")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"rotate_max_files": {
			Description:  lang.Markdown("Specifies the maximum number of older audit log file archives to keep. If 0, no files are ever deleted."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
	},
}

var FilterSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description:  lang.Markdown("Specifies the type of filter to create. Currently only HTTPEvent is supported."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("HTTPEvent")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"endpoints": {
			Description:  lang.Markdown("Specifies the list of endpoints to apply the filter to."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"stages": {
			Description:  lang.Markdown("Specifies the list of stages (`\"OperationReceived\"`, `\"OperationComplete\"`, `\"*\"`) to apply the filter to for a matching endpoint."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"operations": {
			Description:  lang.Markdown("Specifies the list of operations to apply the filter to for a matching endpoint. For HTTPEvent types this corresponds to an HTTP verb (GET, PUT, POST, DELETE...)."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
	},
}
