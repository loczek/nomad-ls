package lsp

import (
	"fmt"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/schema"
)

func CollectHoverInfo(body hcl.Body, pos hcl.Pos, schemaMap map[string]*hcl.BodySchema) []string {
	return []string{CollectHoverInfoDFS(body, schemaMap, "root", pos, &schema.RootBodySchema)}
}

func CollectHoverInfoDFS(
	body hcl.Body,
	schemaMap map[string]*hcl.BodySchema,
	schemaKey string,
	pos hcl.Pos,
	langSchema *hclschema.BodySchema,
) string {
	if schemaMap[schemaKey] == nil {
		return ""
	}

	bodyContent, _ := body.Content(schemaMap[schemaKey])
	blocksByType := bodyContent.Blocks.ByType()

	ans := ""

	for k, v := range blocksByType {
		for _, b := range v {
			blockRange := b.Body.(*hclsyntax.Body).SrcRange
			if !blockRange.ContainsPos(pos) {
				blockRange := b.TypeRange
				if blockRange.ContainsPos(pos) {
					return langSchema.Blocks[k].Description.Value
				}
				continue
			}

			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				ans = CollectHoverInfoDFS(b.Body, schemaMap, k, pos, langSchema.Blocks[k].Body)
			} else if langSchema.Blocks[k] != nil && langSchema.Blocks[k].DependentBody != nil {
				if bodyContent.Attributes["driver"] != nil {
					driver, _ := bodyContent.Attributes["driver"].Expr.Value(&hcl.EvalContext{})

					schemaMapDependentKey := fmt.Sprintf("%s:%s", k, driver.AsString())

					ans = CollectHoverInfoDFS(b.Body, schemaMap, schemaMapDependentKey, pos, langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(driver.AsString())])
				}
			}
		}
	}

	for k, v := range bodyContent.Attributes {
		if v.NameRange.ContainsPos(pos) {
			return langSchema.Attributes[k].Description.Value
		}
	}

	return ans
}
