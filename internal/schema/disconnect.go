package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var DisconnectSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"lost_after": {
			Description:  lang.Markdown("Specifies a duration during which a Nomad client will attempt to reconnect allocations after it fails to heartbeat in the [`heartbeat_grace`](https://developer.hashicorp.com/nomad/docs/configuration/server#heartbeat_grace) window. It defaults to \"\", which is equivalent to having the disconnect block be nil.\n\nYou cannot use `lost_after` and `stop_on_client_after` in the same `disconnect` block.\n\nRefer to [the Lost After section](https://developer.hashicorp.com/nomad/docs/job-specification/disconnect#lost-after) for more details."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"replace": {
			Description:  lang.Markdown("Specifies if Nomad should replace the disconnected allocation with a new one rescheduled on a different node. Nomad considers the replacement allocation a reschedule and obeys the job's [`reschedule`](https://developer.hashicorp.com/nomad/docs/job-specification/reschedule) block. If false and the node the allocation is running on disconnects or goes down, Nomad does not replace this allocation and reports `unknown` until the node reconnects, or until you manually stop the allocation.\n\n`nomad alloc stop  <alloc ID>`\n\nIf true, a new alloc will be placed immediately upon the node becoming disconnected."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"stop_on_client_after": {
			Description:  lang.Markdown("Specifies a duration after which a disconnected Nomad client will stop its allocations. Setting `stop_on_client_after` shorter than `lost_after` and `replace = false` at the same time is not permitted and will cause a validation error, because this would lead to a state where no allocations can be scheduled.\n\nThe Nomad client process must be running for this to occur.\n\nYou cannot use `stop_on_client_after` and `lost_after` in the same `disconnect` block.\n\nRefer to [the Stop After section](https://developer.hashicorp.com//nomad/docs/job-specification/disconnect#stop-after) for more details."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"reconcile": {
			Description:  lang.Markdown("Specifies which allocation to keep once the previously disconnected node regains connectivity. It has four possible values which are described below:\n\n* [`keep_original`](https://developer.hashicorp.com/nomad/docs/job-specification/disconnect#keep_original): Always keep the original allocation. Bear in mind when choosing this option, it can have crashed while the client was disconnected.\n* [`keep_replacement`](https://developer.hashicorp.com/nomad/docs/job-specification/disconnect#keep_replacement): Always keep the allocation that was replaced to replace the disconnected one.\n* [`best_score`](https://developer.hashicorp.com/nomad/docs/job-specification/disconnect#best_score): Keep the allocation running on the node with the best score.\n* [`longest_running`](https://developer.hashicorp.com/nomad/docs/job-specification/disconnect#longest_running): Keep the allocation that has been up and running continuously for the longest time."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("best_score")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
