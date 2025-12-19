package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var NumaSchema = &schema.BodySchema{
	DocsLink: &schema.DocsLink{
		URL: "https://developer.hashicorp.com/nomad/docs/job-specification/numa",
	},
	Attributes: map[string]*schema.AttributeSchema{
		"affinity": {
			Description:  lang.Markdown("Specifies the strategy Nomad will use when selecting CPU cores to assign to a task. Possible values are `none`, `prefer`, or `require`.\n\n- [`none`](https://developer.hashicorp.com/nomad/docs/job-specification/numa#none) - Nomad is free to allocate CPU cores using any strategy. Nomad uses this freedom to allocate cores in such a way that minimizes the amount of fragmentation of core availability per NUMA node. It does so by bin-packing the chosen cores onto the NUMA nodes with the fewest number of cores available.\n- [`prefer`](https://developer.hashicorp.com/nomad/docs/job-specification/numa#prefer) - Nomad will select the set of CPU cores on a node that minimizes the total distance between those cores, but does not limit those CPU core selections to come from a single NUMA node.\n- [`require`](https://developer.hashicorp.com/nomad/docs/job-specification/numa#require) - Nomad will select a set of CPU cores that are strictly colocated on the same hardware NUMA node. If there are no Nomad nodes with a sufficient number of available cores in a compatible configuration, task placement will fail due to exhausted resources."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("none")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsRequired:   true,
		},
		"devices": {
			Description: lang.Markdown("Specifies which devices must be colocated on the name NUMA node, along with allocated CPU cores. Must be a subset of the devices listed in the [resources](https://developer.hashicorp.com/nomad/docs/job-specification/resources) block. May only be used with `affinity` set to `require`."),
			Constraint:  schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:  true,
		},
	},
}
