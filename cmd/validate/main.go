package main

import (
	"fmt"

	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/loczek/nomad-ls/internal/schema/job/drivers"
)

func main() {
	if err := schema.NomadACL.Validate(); err != nil {
		panic(fmt.Sprintf("ACL schema validation error: %s", err))
	}
	if err := schema.NomadAgent.Validate(); err != nil {
		panic(fmt.Sprintf("Agent schema validation error: %s", err))
	}
	if err := schema.NomadCSIVolume.Validate(); err != nil {
		panic(fmt.Sprintf("CSI Volume schema validation error: %s", err))
	}
	if err := schema.NomadDynamicHostVolume.Validate(); err != nil {
		panic(fmt.Sprintf("Dynamic Host Volume schema validation error: %s", err))
	}
	if err := schema.NomadJob.Validate(); err != nil {
		panic(fmt.Sprintf("Job schema validation failed: %s", err))
	}
	if err := schema.NomadNamespace.Validate(); err != nil {
		panic(fmt.Sprintf("Namespace schema validation error: %s", err))
	}
	if err := schema.NomadNodePool.Validate(); err != nil {
		panic(fmt.Sprintf("Nomad Pool schema validation error: %s", err))
	}
	if err := schema.NomadResourceQuota.Validate(); err != nil {
		panic(fmt.Sprintf("ResourceQuota schema validation error: %s", err))
	}
	if err := schema.NomadVariable.Validate(); err != nil {
		panic(fmt.Sprintf("Variable schema validation error: %s", err))
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
