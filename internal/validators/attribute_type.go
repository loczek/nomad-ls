package custom_validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl-lang/validator"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

var _ validator.Validator = (*TypeAttribute)(nil)

type TypeAttribute struct{}

func (v TypeAttribute) Visit(ctx context.Context, node hclsyntax.Node, nodeSchema schema.Schema) (context.Context, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	attr, ok := node.(*hclsyntax.Attribute)
	if !ok {
		return ctx, diags
	}

	if nodeSchema == nil {
		return ctx, diags
	}

	attrSchema := nodeSchema.(*schema.AttributeSchema)
	val, _ := attr.Expr.Value(&hcl.EvalContext{})

	if c, ok := attrSchema.Constraint.(schema.LiteralType); ok {
		rng := attr.Expr.Range()

		if c.Type != cty.DynamicPseudoType && val.Type() != c.Type {
			diags = append(diags, &hcl.Diagnostic{
				Detail:  fmt.Sprintf("Type \"%s\" is not assignable to type \"%s\"", val.Type().FriendlyName(), c.Type.FriendlyName()),
				Subject: &rng,
			})
		}
	}

	return ctx, diags
}
