package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var operatorValues = "`=`, `!=`, `>`, `>=`, `<`, `<=`, `distinct_hosts`, `distinct_property`, `regexp`, `set_contains`, `set_contains_any`, `version`, `semver`, `is_set`, `is_not_set`"

var ConstraintSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		// Standard attributes
		"attribute": {
			Description: lang.Markdown("Specifies the name or reference of the attribute to examine for the constraint. This can be any of the [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars).\n\n**Required** for most operators. **Not used** with `distinct_hosts` operator. **Optional** when using `distinct_property` shorthand."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsOptional: true,
		},
		"operator": {
			Description: lang.Markdown("Specifies the comparison operator. Defaults to `=` if not specified.\n\nIf the operator is one of `>`, `>=`, `<`, `<=`, the ordering is compared numerically if the operands are both integers or both floats, and lexically otherwise.\n\nPossible values: " + operatorValues + "\n\n**Note:** When using shorthand attributes (`distinct_hosts`, `distinct_property`, `set_contains`, `set_contains_any`), this attribute should be omitted."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal("="),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsOptional: true,
		},
		"value": {
			Description: lang.Markdown("Specifies the value to compare the attribute against using the specified operation. This can be a literal value, another attribute, or any [Nomad interpolated values](https://developer.hashicorp.com/nomad/docs/reference/runtime-variable-interpolation#interpreted_node_vars).\n\n**Required** for most operators. **Not required** when using `is_set`, `is_not_set` operators. **Optional** for `distinct_hosts` and `distinct_property` operators."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.StringVal(""),
			},
			Constraint: &schema.LiteralType{Type: cty.String},
			IsOptional: true,
		},

		// Shorthand attributes - these provide a more compact representation
		"distinct_hosts": {
			Description: lang.Markdown("Shorthand for `operator = \"distinct_hosts\"`. Instructs the scheduler to not co-locate any groups on the same machine.\n\nWhen specified as a job constraint, it applies to all groups in the job. When specified as a group constraint, the effect is constrained to that group. This constraint **cannot** be specified at the task level.\n\n```hcl\nconstraint {\n  distinct_hosts = true\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  operator = \"distinct_hosts\"\n  value    = \"true\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.Bool},
			IsOptional:  true,
		},
		"distinct_property": {
			Description: lang.Markdown("Shorthand for `operator = \"distinct_property\"`. Instructs the scheduler to select nodes that have a distinct value of the specified property (attribute).\n\nThe `value` parameter specifies how many allocations are allowed to share the value of a property. The `value` must be 1 or greater and if omitted, defaults to 1.\n\nWhen specified as a job constraint, it applies to all groups in the job. When specified as a group constraint, the effect is constrained to that group. This constraint **cannot** be specified at the task level.\n\n```hcl\nconstraint {\n  distinct_property = \"${meta.rack}\"\n  value             = \"3\"\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  operator  = \"distinct_property\"\n  attribute = \"${meta.rack}\"\n  value     = \"3\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"set_contains": {
			Description: lang.Markdown("Shorthand for `operator = \"set_contains\"`. Specifies a contains constraint against the attribute. The attribute and the list being checked are split using commas. This will check that the given attribute contains **all** of the specified elements.\n\n```hcl\nconstraint {\n  attribute    = \"${meta.cached_binaries}\"\n  set_contains = \"redis,cypress,nginx\"\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  attribute = \"${meta.cached_binaries}\"\n  operator  = \"set_contains\"\n  value     = \"redis,cypress,nginx\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"set_contains_any": {
			Description: lang.Markdown("Shorthand for `operator = \"set_contains_any\"`. Specifies a contains constraint against the attribute. The attribute and the list being checked are split using commas. This will check that the given attribute contains **any** of the specified elements.\n\n```hcl\nconstraint {\n  attribute        = \"${meta.cached_binaries}\"\n  set_contains_any = \"redis,cypress,nginx\"\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  attribute = \"${meta.cached_binaries}\"\n  operator  = \"set_contains_any\"\n  value     = \"redis,cypress,nginx\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"regexp": {
			Description: lang.Markdown("Shorthand for `operator = \"regexp\"`. Specifies a regular expression constraint against the attribute. The syntax of the regular expressions accepted is the same general syntax used by Perl, Python, and many other languages (RE2 syntax).\n\n```hcl\nconstraint {\n  attribute = \"${attr.kernel.name}\"\n  regexp    = \"linux|darwin\"\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  attribute = \"${attr.kernel.name}\"\n  operator  = \"regexp\"\n  value     = \"linux|darwin\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"version": {
			Description: lang.Markdown("Shorthand for `operator = \"version\"`. Specifies a version constraint against the attribute. This supports a comma-separated list of constraints, including the pessimistic operator.\n\n`version` will not consider a prerelease (eg `1.6.0-beta`) sufficient to match a non-prerelease constraint (eg `>= 1.0`). Use `semver` for strict Semantic Versioning 2.0 ordering.\n\n```hcl\nconstraint {\n  attribute = \"${attr.kernel.version}\"\n  version   = \">= 3.19\"\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  attribute = \"${attr.kernel.version}\"\n  operator  = \"version\"\n  value     = \">= 3.19\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
		"semver": {
			Description: lang.Markdown("Shorthand for `operator = \"semver\"`. Specifies a Semantic Versioning 2.0 constraint against the attribute. Unlike `version`, this operator considers prereleases (eg `1.6.0-beta`) sufficient to satisfy non-prerelease constraints (eg `>= 1.0`).\n\n```hcl\nconstraint {\n  attribute = \"${attr.vault.version}\"\n  semver    = \">= 1.0.0\"\n}\n```\n\nEquivalent to:\n```hcl\nconstraint {\n  attribute = \"${attr.vault.version}\"\n  operator  = \"semver\"\n  value     = \">= 1.0.0\"\n}\n```"),
			Constraint:  &schema.LiteralType{Type: cty.String},
			IsOptional:  true,
		},
	},
}
