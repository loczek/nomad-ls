package references

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/reference"
	"github.com/loczek/nomad-ls/internal/scope"
	"github.com/zclconf/go-cty/cty"
)

func CommonBuiltinReferences() reference.Targets {
	return reference.Targets{
		{
			Addr: lang.Address{
				lang.RootStep{Name: "node"},
				lang.AttrStep{Name: "unique"},
				lang.AttrStep{Name: "id"},
			},
			Type:        cty.String,
			ScopeId:     scope.BuiltinScope,
			Description: lang.Markdown("36 character unique client identifier\n\n**Example**: `9afa5da1-8f39-25a2-48dc-ba31fd7c0023`"),
		},
		{
			Addr: lang.Address{
				lang.RootStep{Name: "node"},
				lang.AttrStep{Name: "region"},
			},
			Type:        cty.String,
			ScopeId:     scope.BuiltinScope,
			Description: lang.Markdown("Client's region\n\n**Example**: `global`"),
		},
		{
			Addr: lang.Address{
				lang.RootStep{Name: "node"},
				lang.AttrStep{Name: "datacenter"},
			},
			Type:        cty.String,
			ScopeId:     scope.BuiltinScope,
			Description: lang.Markdown("Client's datacenter\n\n**Example**: `dc1`"),
		},
		{
			Addr: lang.Address{
				lang.RootStep{Name: "node"},
				lang.AttrStep{Name: "unique"},
				lang.AttrStep{Name: "name"},
			},
			Type:        cty.String,
			ScopeId:     scope.BuiltinScope,
			Description: lang.Markdown("Client's name\n\n**Example**: `nomad-client-10-1-2-4`"),
		},
		{
			Addr: lang.Address{
				lang.RootStep{Name: "node"},
				lang.AttrStep{Name: "class"},
			},
			Type:        cty.String,
			ScopeId:     scope.BuiltinScope,
			Description: lang.Markdown("Client's class\n\n**Example**: `linux-64bit`"),
		},
		{
			Addr: lang.Address{
				lang.RootStep{Name: "node"},
				lang.AttrStep{Name: "pool"},
			},
			Type:        cty.String,
			ScopeId:     scope.BuiltinScope,
			Description: lang.Markdown("Client's node pool\n\n**Example**: `prod`"),
		},
	}
}
