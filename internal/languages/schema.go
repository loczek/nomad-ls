package languages

import (
	hclSchema "github.com/hashicorp/hcl-lang/schema"
	schemaACL "github.com/loczek/nomad-ls/internal/schema/acl"
	schemaJob "github.com/loczek/nomad-ls/internal/schema/job"
	schemaNamespace "github.com/loczek/nomad-ls/internal/schema/namespace"
	schemaNodePool "github.com/loczek/nomad-ls/internal/schema/node-pool"
	schemaResourceQuota "github.com/loczek/nomad-ls/internal/schema/resource-quota"
	schemaVariable "github.com/loczek/nomad-ls/internal/schema/variable"
	schemaVolumeCSI "github.com/loczek/nomad-ls/internal/schema/volume/csi"
	schemaVolumeDynamic "github.com/loczek/nomad-ls/internal/schema/volume/dynamic"
)

var schemaMap = map[LanguageID]hclSchema.BodySchema{
	NomadACL:           schemaACL.RootSchema,
	NomadJob:           schemaJob.RootBodySchema, // TODO: rename this
	NomadNapespace:     schemaNamespace.RootSchema,
	NomadNodePool:      schemaNodePool.RootSchema,
	NomadResourceQuota: schemaResourceQuota.RootSchema,
	NomadVariable:      schemaVariable.RootSchema,
	NomadVolumeCSI:     schemaVolumeCSI.RootSchema,
	NomadVolumeDynamic: schemaVolumeDynamic.RootSchema,
}

func ToSchema(lang LanguageID) hclSchema.BodySchema {
	if schema, ok := schemaMap[lang]; ok {
		return schema
	}
	return schemaJob.RootBodySchema
}
