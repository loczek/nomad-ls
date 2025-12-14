package lsp

import (
	"fmt"
	"log"
	"strconv"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
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

			if len(v.Labels) != 0 {
				text = asBlock(k, depth)
			} else {
				text = asAnonymousBlock(k, depth)
			}

			blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
				Label:      k,
				Kind:       protocol.CompletionItemKindInterface,
				InsertText: text,
				Documentation: protocol.MarkupContent{
					Kind:  protocol.Markdown,
					Value: v.Description.Value,
				},
				InsertTextFormat: protocol.InsertTextFormatSnippet,
			})
		}

		for k, v := range langSchema.Attributes {
			if v.Constraint == nil {
				continue
			}

			c, ok := v.Constraint.(*hclschema.LiteralType)
			if !ok {
				continue
			}

			var insertText string

			switch c.Type {
			case cty.String:
				insertText = fmt.Sprintf("%s = \"$0\"", k)
			case cty.Number:
			case cty.Bool:
				insertText = fmt.Sprintf("%s = $0", k)
			case cty.List(cty.String):
				insertText = fmt.Sprintf("%s = [$0]", k)
			case cty.Map(cty.String):
				insertText = fmt.Sprintf("%s = {$0}", k)
			default:
				insertText = fmt.Sprintf("%s = ", k)
			}

			if bodyContent.Attributes[k] != nil {
				continue
			}

			d, ok := v.DefaultValue.(*hclschema.DefaultValue)
			if ok {
				switch c.Type {
				case cty.String:
					insertText = fmt.Sprintf("%s = \"${0:%s}\"", k, d.Value.AsString())
				case cty.Number:
					val, err := convert.Convert(d.Value, cty.String)

					if err != nil {
						continue
					}

					insertText = fmt.Sprintf("%s = ${0:%s}", k, val.AsString())
				case cty.Bool:
					insertText = fmt.Sprintf("%s = ${0:%s}", k, strconv.FormatBool(d.Value.True()))
				case cty.List(cty.String):
					var arr []string

					for _, b := range d.Value.Elements() {
						arr = append(arr, b.AsString())
					}

					insertText = fmt.Sprintf("%s = [\"${0:%s}\"]", k, arr)
				case cty.Map(cty.String):
					var arr = make(map[string]string)

					for a, b := range d.Value.Elements() {
						arr[a.AsString()] = b.AsString()
					}

					insertText = fmt.Sprintf("%s = {${0:%s}}", k, formatMap(arr))
				default:
					insertText = fmt.Sprintf("%s = ", k)
				}
			}

			blocksByTypeArr = append(blocksByTypeArr, protocol.CompletionItem{
				Label:      k,
				Kind:       protocol.CompletionItemKindVariable,
				InsertText: insertText,
				Detail:     c.FriendlyName(),
				Documentation: protocol.MarkupContent{
					Kind:  protocol.Markdown,
					Value: v.Description.Value,
				},
				InsertTextFormat: protocol.InsertTextFormatSnippet,
			})
		}

		*blocks = append(*blocks, blocksByTypeArr...)
	}

	log.Printf("matching blocks: %d", matchingBlocks)
}

func formatMap(input map[string]string) string {
	ans := "\n"

	for k, v := range input {
		ans += fmt.Sprintf("\"%s\": \"%s\"\n", k, v)
	}

	return ans
}
