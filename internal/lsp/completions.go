package lsp

import (
	"fmt"
	"log"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/zclconf/go-cty/cty"
	"go.lsp.dev/protocol"
)

func CollectCompletions(body hcl.Body, pos hcl.Pos, schemaMap map[string]*hcl.BodySchema) []protocol.CompletionItem {
	var blocks []protocol.CompletionItem

	CollectCompletionsDFS(body, &blocks, schemaMap, "root", pos, &schema.RootBodySchema, 1)

	return blocks
}

func CollectCompletionsDFS(
	body hcl.Body,
	blocks *[]protocol.CompletionItem,
	schemaMap map[string]*hcl.BodySchema,
	schemaKey string,
	pos hcl.Pos,
	langSchema *hclschema.BodySchema,
	depth int,
) {
	if schemaMap[schemaKey] == nil {
		return
	}

	bodyContent, _ := body.Content(schemaMap[schemaKey])
	blocksByType := bodyContent.Blocks.ByType()

	var matchingBlocks uint

	for k, v := range blocksByType {
		for _, b := range v {
			blockRange := b.Body.(*hclsyntax.Body).SrcRange
			if !blockRange.ContainsPos(pos) {
				continue
			}

			matchingBlocks += 1

			if langSchema.Blocks[k] != nil && langSchema.Blocks[k].Body != nil {
				CollectCompletionsDFS(b.Body, blocks, schemaMap, k, pos, langSchema.Blocks[k].Body, depth+1)
			} else if langSchema.Blocks[k] != nil && langSchema.Blocks[k].DependentBody != nil {
				if bodyContent.Attributes["driver"] != nil {
					driver, _ := bodyContent.Attributes["driver"].Expr.Value(&hcl.EvalContext{})

					schemaMapDependentKey := fmt.Sprintf("%s:%s", k, driver.AsString())

					CollectCompletionsDFS(b.Body, blocks, schemaMap, schemaMapDependentKey, pos, langSchema.Blocks[k].DependentBody[hclschema.SchemaKey(driver.AsString())], depth+1)
				}
			}
		}
	}

	if matchingBlocks == 0 {
		var blocksByTypeArr []protocol.CompletionItem

		for k, v := range langSchema.Blocks {
			var text string
			var detail string

			if len(v.Labels) != 0 {
				text = asBlock(k, depth)
				detail = "named"
			} else {
				text = asAnonymousBlock(k, depth)
			}

			blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
				Label:            k,
				InsertText:       text,
				Kind:             protocol.CompletionItemKindInterface,
				InsertTextFormat: protocol.InsertTextFormatSnippet,
				Detail:           detail,
			})
		}

		for k, v := range langSchema.Attributes {
			if v.Constraint == nil {
				continue
			}

			h, ok := v.Constraint.(*hclschema.LiteralType)
			if !ok {
				continue
			}

			log.Printf("attr: %s", k)
			log.Printf("%+v", bodyContent.Attributes)

			if bodyContent.Attributes[k] != nil {
				continue
			}

			switch h.Type {
			case cty.String:
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: fmt.Sprintf("%s = \"$0\"", k),
					Kind:       protocol.CompletionItemKindVariable,
					Documentation: protocol.MarkupContent{
						Kind:  protocol.Markdown,
						Value: v.Description.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			case cty.List(cty.String):
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: fmt.Sprintf("%s = [\"$0\"]", k),
					Kind:       protocol.CompletionItemKindVariable,
					Documentation: protocol.MarkupContent{
						Kind:  protocol.Markdown,
						Value: v.Description.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			default:
				blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
					Label:      k,
					InsertText: fmt.Sprintf("%s = ", k),
					Kind:       protocol.CompletionItemKindVariable,
					Documentation: protocol.MarkupContent{
						Kind:  protocol.Markdown,
						Value: v.Description.Value,
					},
					InsertTextFormat: protocol.InsertTextFormatSnippet,
				})
			}
		}

		*blocks = append(*blocks, blocksByTypeArr...)
	}

	log.Printf("matching blocks: %d", matchingBlocks)
}
