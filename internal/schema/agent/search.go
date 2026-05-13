package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var SearchSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"fuzzy_enabled": {
			Description:  lang.Markdown("Specifies whether the [fuzzy search API](https://developer.hashicorp.com/nomad/api-docs/search#fuzzy-searching	) is enabled. If not enabled, requests to the fuzzy search API endpoint will return an error response."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"limit_query": {
			Description:  lang.Markdown("Specifies the maximum number of Nomad objects to search through per context type in the Nomad server before truncating results. Setting this parameter to a high value may degrade Nomad server performance."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(20)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"limit_results": {
			Description:  lang.Markdown("Specifies the maximum number of matching results to accumulate per context type in the API response before truncating results. Setting this parameter to a high value may cause excessively large API response sizes."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(100)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"min_term_length": {
			Description:  lang.Markdown("Specifies the minimum size of the search term allowed for matching with the fuzzy search API. Setting this value higher can prevent unnecessary load on the Nomad server from broad queries."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(2)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
	},
}
