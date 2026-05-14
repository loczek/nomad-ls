package main

import (
	"fmt"

	hclschema "github.com/hashicorp/hcl-lang/schema"
	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/loczek/nomad-ls/internal/schema/agent/keyring"
	plugin "github.com/loczek/nomad-ls/internal/schema/agent/plugins"
	"github.com/loczek/nomad-ls/internal/schema/job/drivers"
)

func main() {
	keyrings := map[string]*hclschema.BodySchema{
		"aws":   keyring.AWSSchema,
		"azure": keyring.AzureSchema,
		"gcp":   keyring.GCPSchema,
		"vault": keyring.VaultSchema,
	}

	agentPlugins := map[string]*hclschema.BodySchema{
		"docker":  plugin.DockerSchema,
		"exec":    plugin.ExecSchema,
		"java":    plugin.JavaSchema,
		"qemu":    plugin.QEMUSchema,
		"rawexec": plugin.RawExecSchema,
	}

	nomadObjects := map[string]*hclschema.BodySchema{
		"acl":                 schema.NomadACL,
		"agent":               schema.NomadAgent,
		"csi_volume":          schema.NomadCSIVolume,
		"dynamic_host_volume": schema.NomadDynamicHostVolume,
		"job":                 schema.NomadJob,
		"namespace":           schema.NomadNamespace,
		"node_pool":           schema.NomadNodePool,
		"resource_quota":      schema.NomadResourceQuota,
		"variable":            schema.NomadVariable,
	}

	jobDrivers := map[string]*hclschema.BodySchema{
		"docker":   drivers.DockerDriverSchema,
		"exec":     drivers.ExecDriverSchema,
		"java":     drivers.JavaDriverSchema,
		"qemu":     drivers.QemuDriverSchema,
		"raw_exec": drivers.RawExecDriverSchema,
	}

	for name, v := range keyrings {
		if err := v.Validate(); err != nil {
			panic(fmt.Errorf("Keyring \"%s\" validation error: %w", name, err))
		}
	}

	for name, v := range agentPlugins {
		if err := v.Validate(); err != nil {
			panic(fmt.Errorf("Agent Plugin \"%s\" validation error: %w", name, err))
		}
	}

	for name, v := range nomadObjects {
		if err := v.Validate(); err != nil {
			panic(fmt.Errorf("Nomad \"%s\" validation error: %w", name, err))
		}
	}

	for name, v := range jobDrivers {
		if err := v.Validate(); err != nil {
			panic(fmt.Errorf("Job Driver \"%s\" validation error: %w", name, err))
		}
	}
}
