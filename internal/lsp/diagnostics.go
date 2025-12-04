package lsp

import (
	"fmt"
	"log"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/loczek/nomad-ls/internal/schema"
)

func CollectDiagnistics(body hcl.Body, schemaMap map[string]*hcl.BodySchema) *hcl.Diagnostics {
	var diags hcl.Diagnostics

	diags = diags.Extend(CollectDiagnisticsDFS(body, &diags, schemaMap, schema.SchemaMapBetter["root"], &schema.RootBodySchema))

	return &diags
}

func CollectDiagnisticsDFS(body hcl.Body, diags *hcl.Diagnostics, schemaMap map[string]*hcl.BodySchema, currSchema *hcl.BodySchema, langSchema *hclschema.BodySchema) hcl.Diagnostics {
	if currSchema == nil {
		return make(hcl.Diagnostics, 0)
	}

	bodyContent, allDiags := body.Content(currSchema)
	blocksByType := bodyContent.Blocks.ByType()

	for k, v := range blocksByType {
		for _, b := range v {
			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				allDiags = allDiags.Extend(CollectDiagnisticsDFS(b.Body, diags, schemaMap, schemaMap[k], langSchema.Blocks[k].Body))
			} else if langSchema.Blocks[k] != nil && langSchema.Blocks[k].DependentBody != nil {
				log.Printf("found config!")
				if bodyContent.Attributes["driver"] != nil {
					driver, _ := bodyContent.Attributes["driver"].Expr.Value(&hcl.EvalContext{})

					log.Printf("driver: %s", driver.AsString())

					schemaMapDependentKey := fmt.Sprintf("%s:%s", k, driver.AsString())

					log.Printf("map key: %s", schemaMapDependentKey)

					// langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(bodyContent.Attributes["driver"].Name)]
					allDiags = allDiags.Extend(CollectDiagnisticsDFS(b.Body, diags, schemaMap, schemaMap[schemaMapDependentKey], langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(driver.AsString())]))
				}
			}
		}
	}

	return allDiags
}
