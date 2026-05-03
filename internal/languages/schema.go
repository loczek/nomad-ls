package languages

import (
	hclSchema "github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/schema"
)

var schemaMap = map[LanguageID]hclSchema.BodySchema{
	NomadACL:           schema.NomadACL,
	NomadAgent:         schema.NomadAgent,
	NomadJob:           schema.NomadJob,
	NomadNapespace:     schema.NomadNamespace,
	NomadNodePool:      schema.NomadNodePool,
	NomadResourceQuota: schema.NomadResourceQuota,
	NomadVariable:      schema.NomadVariable,
	NomadVolumeCSI:     schema.NomadVolumeCSI,
	NomadVolumeDynamic: schema.NomadVolumeDynamic,
}

func ToSchema(lang LanguageID) hclSchema.BodySchema {
	if schema, ok := schemaMap[lang]; ok {
		return schema
	}
	return schema.NomadJob
}
