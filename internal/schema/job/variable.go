package job

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// TODO: update docs
var VariableSchema = &schema.BodySchema{
	Description: lang.Markdown("Input variables serve as parameters for a Nomad job, allowing aspects of the job to be customized without altering the job's own source code.\nWhen you declare variables in the same file as the job specification, you can set their values using CLI options and environment variables."),
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description: lang.Markdown("The type of HCL variable: `string`, `number`, `bool`."),
			Constraint:  schema.TypeDeclaration{},
			IsOptional:  true,
		},
		"default": {
			Description: lang.Markdown("The default value used when no value for this variable is provided."),
			Constraint:  schema.LiteralType{Type: cty.DynamicPseudoType},
			IsOptional:  true,
		},
		"description": {
			Description: lang.Markdown("variable description"),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"validation": {
			Description: lang.PlainText("Input variables support specifying arbitrary custom validation rules for a particular variable using a `validation` block nested within the corresponding `variable` block"),
			Body:        ValidationSchema,
		},
	},
}

var ValidationSchema = &schema.BodySchema{
	Description: lang.Markdown("Input variables serve as parameters for a Nomad job, allowing aspects of the job to be customized without altering the job's own source code.\nWhen you declare variables in the same file as the job specification, you can set their values using CLI options and environment variables."),
	Attributes: map[string]*schema.AttributeSchema{
		"condition": {
			Description: lang.Markdown("The condition argument is an expression that must use the value of the variable to return true if the value is valid, or false if it is invalid. The expression can refer only to the variable that the condition applies to, and must not produce errors."),
			Constraint:  schema.AnyExpression{OfType: cty.DynamicPseudoType},
			IsRequired:  true,
		},
		"error_message": {
			Description: lang.Markdown("The error message string should be at least one full sentence explaining the constraint that failed, starting with an uppercase letter ( if the alphabet permits it ) and ending with a period or question mark."),
			Constraint:  schema.LiteralType{Type: cty.String},
			IsRequired:  true,
		},
	},
}
