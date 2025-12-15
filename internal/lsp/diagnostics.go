package lsp

import (
	"fmt"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/loczek/nomad-ls/internal/schema"
)

func CollectDiagnostics(body hcl.Body, schemaMap map[string]*hcl.BodySchema) *hcl.Diagnostics {
	var diags hcl.Diagnostics

	diags = diags.Extend(CollectDiagnosticsDFS(body, &diags, schemaMap, schema.SchemaMapBetter["root"], &schema.RootBodySchema))

	return &diags
}

func CollectDiagnosticsDFS(body hcl.Body, diags *hcl.Diagnostics, schemaMap map[string]*hcl.BodySchema, currSchema *hcl.BodySchema, langSchema *hclschema.BodySchema) hcl.Diagnostics {
	if currSchema == nil {
		return make(hcl.Diagnostics, 0)
	}

	var bodyContent *hcl.BodyContent
	var allDiags hcl.Diagnostics

	// Use PartialContent for schemas that allow any attribute (like `variables` or `meta`)
	// to avoid false errors for user-defined attributes
	if langSchema != nil && langSchema.AnyAttribute != nil {
		bodyContent, _, allDiags = body.PartialContent(currSchema)
	} else {
		bodyContent, allDiags = body.Content(currSchema)
	}

	blocksByType := bodyContent.Blocks.ByType()

	for k, v := range blocksByType {
		for _, b := range v {
			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				allDiags = allDiags.Extend(CollectDiagnosticsDFS(b.Body, diags, schemaMap, schemaMap[k], langSchema.Blocks[k].Body))
			} else if langSchema.Blocks[k] != nil && langSchema.Blocks[k].DependentBody != nil {
				if bodyContent.Attributes["driver"] != nil {
					driver, _ := bodyContent.Attributes["driver"].Expr.Value(&hcl.EvalContext{})

					schemaMapDependentKey := fmt.Sprintf("%s:%s", k, driver.AsString())

					allDiags = allDiags.Extend(CollectDiagnosticsDFS(b.Body, diags, schemaMap, schemaMap[schemaMapDependentKey], langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(driver.AsString())]))
				}
			}
		}
	}

	return allDiags
}
