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

var JobConfigSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: variablesLabel},
		{Type: variableLabel, LabelNames: []string{"name"}},
		{Type: localsLabel},
		{Type: "job", LabelNames: []string{"name"}},
	},
}

var JobSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "region"},
		{Name: "datacenter"},
		{Name: "type"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "group", LabelNames: []string{"name"}},
	},
}

var JobGroupSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "count"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "ephemeral_disk"},
		{Type: "service"},
	},
}

var JobGroupEphemeralDiskSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "migrate"},
		{Name: "size"},
		{Name: "sticky"},
	},
}

var JobGroupServiceSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "name"},
		{Name: "port"},
		{Name: "provider"},
	},
}

var SchemaMap map[string]*hcl.BodySchema = map[string]*hcl.BodySchema{
	"job":            JobSchema,
	"group":          JobGroupSchema,
	"ephemeral_disk": JobGroupEphemeralDiskSchema,
	"service":        JobGroupServiceSchema,
}

var RootBodySchema = schema.BodySchema{
	Blocks: map[string]*schema.BlockSchema{
		"job": {
			Type:        schema.BlockTypeObject,
			Description: lang.PlainText("test"),
		},
	},
}

// var JobConfigSchema = &hcl.BodySchema{
// 	Blocks: []hcl.BlockHeaderSchema{
// 		{Type: variablesLabel},
// 		{Type: variableLabel, LabelNames: []string{"name"}},
// 		{Type: localsLabel},
// 		{Type: "job", LabelNames: []string{"name"}},
// 	},
// }
