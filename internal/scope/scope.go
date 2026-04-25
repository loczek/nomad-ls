package scope

import (
	"github.com/hashicorp/hcl-lang/lang"
)

var (
	BuiltinScope  = lang.ScopeId("builtin")
	LocalScope    = lang.ScopeId("local")
	ModuleScope   = lang.ScopeId("module")
	VariableScope = lang.ScopeId("variable")
	ListScope     = lang.ScopeId("list")
	ActionScope   = lang.ScopeId("action")
	MetaScope     = lang.ScopeId("meta")
)
