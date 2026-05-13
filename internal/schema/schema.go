package schema

import (
	"github.com/loczek/nomad-ls/internal/schema/acl"
	"github.com/loczek/nomad-ls/internal/schema/agent"
	"github.com/loczek/nomad-ls/internal/schema/job"
	"github.com/loczek/nomad-ls/internal/schema/namespace"
	nodePool "github.com/loczek/nomad-ls/internal/schema/node-pool"
	resourceQuota "github.com/loczek/nomad-ls/internal/schema/resource-quota"
	"github.com/loczek/nomad-ls/internal/schema/variable"
	"github.com/loczek/nomad-ls/internal/schema/volume/csi"
	"github.com/loczek/nomad-ls/internal/schema/volume/dynamic"
)

var NomadACL = acl.RootSchema
var NomadAgent = agent.RootSchema
var NomadCSIVolume = csi.RootSchema
var NomadDynamicHostVolume = dynamic.RootSchema
var NomadJob = job.RootSchema
var NomadNamespace = namespace.RootSchema
var NomadNodePool = nodePool.RootSchema
var NomadResourceQuota = resourceQuota.RootSchema
var NomadVariable = variable.RootSchema
