package schemaACL

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ReadPolicy = schema.LiteralValue{
	Value:       cty.StringVal("read"),
	Description: lang.Markdown("allow the resource to be read but not modified"),
}

var WritePolicy = schema.LiteralValue{
	Value:       cty.StringVal("write"),
	Description: lang.Markdown("allow the resource to be read and modified"),
}

var DenyPolicy = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("do not allow the resource to be read or modified. Deny takes precedence when multiple policies are associated with a token."),
}

var ScalePolicy = schema.LiteralValue{
	Value:       cty.StringVal("scale"),
	Description: lang.Markdown("allow the resource to be scaled by the nomad-autoscaler"),
}

var ListPolicy = schema.LiteralValue{
	Value:       cty.StringVal("list"),
	Description: lang.Markdown("allow the resource to be listed, but not inspected in detail"),
}
