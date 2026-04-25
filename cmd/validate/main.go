package main

import (
	"fmt"

	"github.com/loczek/nomad-ls/internal/schema"
	"github.com/loczek/nomad-ls/internal/schema/drivers"
)

func main() {
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
	if err := schema.RootBodySchema.Validate(); err != nil {
		panic(fmt.Sprintf("Nomad Job schema validation error: %s", err))
	}
}
