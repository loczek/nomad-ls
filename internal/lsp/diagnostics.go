package lsp

import (
	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/loczek/nomad-ls/internal/schema"
)

func CollectDiagnostics(body hcl.Body) *hcl.Diagnostics {
	var diags hcl.Diagnostics

	diags = diags.Extend(CollectDiagnosticsDFS(body, &diags, &schema.RootBodySchema))

	return &diags
}

func CollectDiagnosticsDFS(body hcl.Body, diags *hcl.Diagnostics, langSchema *hclschema.BodySchema) hcl.Diagnostics {
	if langSchema == nil {
		return make(hcl.Diagnostics, 0)
	}

	var bodyContent *hcl.BodyContent
	var allDiags hcl.Diagnostics

	// Use PartialContent for schemas that allow any attribute (like `variables` or `meta`)
	// to avoid false errors for user-defined attributes
	if langSchema.AnyAttribute != nil {
		bodyContent, _, allDiags = body.PartialContent(langSchema.ToHCLSchema())
	} else {
		bodyContent, allDiags = body.Content(langSchema.ToHCLSchema())
	}

	blocksByType := bodyContent.Blocks.ByType()

	for k, v := range blocksByType {
		for _, b := range v {
			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				allDiags = allDiags.Extend(CollectDiagnosticsDFS(b.Body, diags, langSchema.Blocks[k].Body))
			} else if langSchema.Blocks[k] != nil && langSchema.Blocks[k].DependentBody != nil {
				if bodyContent.Attributes["driver"] != nil {
					driver, _ := bodyContent.Attributes["driver"].Expr.Value(&hcl.EvalContext{})

					allDiags = allDiags.Extend(CollectDiagnosticsDFS(b.Body, diags, langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(driver.AsString())]))
				}
			}
		}
	}

	return allDiags
}
