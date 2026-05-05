package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var RPCSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"accept_backlog": {
			Description:  lang.Markdown("Limits how many RPC streams (requests) can be waiting for acceptance by the RPC server."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(256)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"keep_alive_interval": {
			Description:  lang.Markdown("How often to perform the keep alive"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"connection_write_timeout": {
			Description:  lang.Markdown("Safety valve timeout after which you should suspect a problem with the underlying connection and close it. This is only applied to writes, where there is generally an expectation that things move along quickly."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("10s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"max_stream_window_size": {
			Description:  lang.Markdown("Controls the maximum window size allowed for a stream, in bytes."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(262144)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"stream_open_timeout": {
			Description:  lang.Markdown("Maximum amount of time that a stream remains in pending state while waiting for an ACK from the peer. Once the timeout is reached, Nomad gracefully closes the session. A zero value disables the stream open timeout, allowing unbounded blocking on open stream calls."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("75s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"stream_close_timeout": {
			Description:  lang.Markdown("When `close` is called, maximum time that a stream is in a half-closed state before forcibly closing the connection. Forcibly closed connections empty the receive buffer, drop any future packets received for that stream, and send an RST to the remote side."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("5m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
