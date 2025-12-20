package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var UISchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"description": {
			Description:  lang.Markdown("The markdown-enabled description of the job. We support [GitHub Flavored Markdown](https://github.github.com/gfm/)."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsDeprecated: true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"link": {
			Description: lang.PlainText("A link that should show up in the header of the job index page in the Web UI. A job can have any number of links, and they must contain both a string `label` and `url`."),
			Body:        LinkSchema,
		},
	},
}

var LinkSchema = &schema.BodySchema{
	AnyAttribute: &schema.AttributeSchema{
		Description: lang.Markdown("A user-defined key-value pair for metadata."),
		Constraint:  schema.LiteralType{Type: cty.String},
		IsOptional:  true,
	},
}
