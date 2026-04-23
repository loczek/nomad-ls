package main

import "github.com/loczek/nomad-ls/internal/schema"

func main() {
	if err := schema.RootBodySchema.Validate(); err != nil {
		panic(err)
	}
}
