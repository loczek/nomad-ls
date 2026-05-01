package main

import (
	"fmt"

	schemaACL "github.com/loczek/nomad-ls/internal/schema/acl"
	schemaAgent "github.com/loczek/nomad-ls/internal/schema/agent"
	schema "github.com/loczek/nomad-ls/internal/schema/job"
	"github.com/loczek/nomad-ls/internal/schema/job/drivers"
	schemaNamespace "github.com/loczek/nomad-ls/internal/schema/namespace"
	schemaNodePool "github.com/loczek/nomad-ls/internal/schema/node-pool"
	schemaResourceQuota "github.com/loczek/nomad-ls/internal/schema/resource-quota"
	schemaVariable "github.com/loczek/nomad-ls/internal/schema/variable"
	"github.com/loczek/nomad-ls/internal/schema/volume/csi"
	"github.com/loczek/nomad-ls/internal/schema/volume/dynamic"
)

func main() {
	if err := schemaACL.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("ACL schema validation error: %s", err))
	}
	if err := schemaAgent.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Agent schema validation error: %s", err))
	}
	if err := schema.RootBodySchema.Validate(); err != nil {
		panic(fmt.Sprintf("Job schema validation failed: %s", err))
	}
	if err := schemaNamespace.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Namespace schema validation error: %s", err))
	}
	if err := schemaNodePool.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Nomad Pool schema validation error: %s", err))
	}
	if err := schemaResourceQuota.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("ResourceQuota schema validation error: %s", err))
	}
	if err := schemaVariable.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Variable schema validation error: %s", err))
	}
	if err := csi.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("CSI Volume schema validation error: %s", err))
	}
	if err := dynamic.RootSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Dynamic Host Volume schema validation error: %s", err))
	}

	validateDrivers()
}

func validateDrivers() {
	if err := drivers.DockerDriverSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Docker driver validation error: %s", err))
	}
	if err := drivers.ExecDriverSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Exec driver validation error: %s", err))
	}
	if err := drivers.JavaDriverSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Java driver validation error: %s", err))
	}
	if err := drivers.QemuDriverSchema.Validate(); err != nil {
		panic(fmt.Sprintf("Qemu driver validation error: %s", err))
	}
	if err := drivers.RawExecDriverSchema.Validate(); err != nil {
		panic(fmt.Sprintf("RawExec driver validation error: %s", err))
	}
}
