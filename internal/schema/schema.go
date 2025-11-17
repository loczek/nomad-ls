package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
)

const (
	variablesLabel = "variables"
	variableLabel  = "variable"
	localsLabel    = "locals"
	vaultLabel     = "vault"
	taskLabel      = "task"
	secretLabel    = "secret"

	inputVariablesAccessor = "var"
	localsAccessor         = "local"
)

var SchemaMapBetter map[string]*hcl.BodySchema = map[string]*hcl.BodySchema{
	"root":  RootBodySchema.Copy().ToHCLSchema(),
	"job":   JobSchemaBetter.Copy().ToHCLSchema(),
	"group": GroupSchema.Copy().ToHCLSchema(),
	// "ephemeral_disk": ephemera,
	"spread": SpreadSchema.Copy().ToHCLSchema(),
	"target": TargetSchema.Copy().ToHCLSchema(),
	"update": UpdateSchema.Copy().ToHCLSchema(),
}

var RootBodySchema = schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		variablesLabel: {},
		variableLabel:  {},
		localsLabel:    {},
	},
	Blocks: map[string]*schema.BlockSchema{
		"job": {
			Description: lang.Markdown("## h2\ntest"),
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: JobSchemaBetter,
		},
	},
}
