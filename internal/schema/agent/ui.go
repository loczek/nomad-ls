package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var UISchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"enabled": {
			Description:  lang.Markdown("Specifies whether the web UI is enabled. If disabled, the `/ui/` path will return an empty web page."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"show_cli_hints": {
			Description:  lang.Markdown("Controls whether CLI commands display hints about equivalent UI pages. For example, when running `nomad server members`, the CLI shows a message indicating where to find server information in the web UI. Set to `false` to disable these hints."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		"content_security_policy": {
			Body: ContentSecurityPolicySchema,
		},
		"consul": {
			Body: ConsulSchema,
		},
		"vault": {
			Body: VaultSchema,
		},
		"label": {
			Body: LabelSchema,
		},
	},
}

var ContentSecurityPolicySchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"connect_src": {
			Description: lang.Markdown("Specifies the valid sources for `connect-src` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("*"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"default_src": {
			Description: lang.Markdown("Specifies the valid sources for `default-src` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("'none'"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"form_action": {
			Description: lang.Markdown("Specifies the valid sources for `form-action` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("'none'"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"frame_ancestors": {
			Description: lang.Markdown("Specifies the valid sources for `frame-ancestors` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("'none'"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"img_src": {
			Description: lang.Markdown("Specifies the valid sources for `img-src` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("'self'"),
				cty.StringVal("data:"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"script_src": {
			Description: lang.Markdown("Specifies the valid sources for `script-src` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("'self'"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
		"style_src": {
			Description: lang.Markdown("Specifies the valid sources for `style-src` in the Content Security Policy header."),
			DefaultValue: schema.DefaultValue{Value: cty.ListVal([]cty.Value{
				cty.StringVal("'none'"),
				cty.StringVal("'unsafe-inline'"),
			})},
			Constraint: schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional: true,
		},
	},
}

var ConsulSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"ui_url": {
			Description:  lang.Markdown("Specifies the full base URL to a Consul web UI (for example: `https://consul.example.com:8501/ui`. This URL is used to build links from the Nomad web UI to a Consul web UI. Note that this URL will not typically be the same one used for the agent's [`consul.address`](https://developer.hashicorp.com/nomad/docs/configuration/consul#address); the `consul.address` is the URL used by the Nomad to communicate with Consul, whereas the `ui.consul.ui_url` is the URL you'll visit in your browser. If this field is omitted, this integration will be disabled."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
var VaultSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"ui_url": {
			Description:  lang.Markdown("Specifies the full base URL to a Vault web UI (for example: `https://vault.example.com:8200/ui`. This URL is used to build links from the Nomad web UI to a Vault web UI. Note that this URL will not typically be the same one used for the agent's [`vault.address`](https://developer.hashicorp.com/nomad/docs/configuration/vault#address); the `vault.address` is the URL used by the Nomad to communicate with Vault, whereas the `ui.vault.ui_url` is the URL you'll visit in your browser. If this field is omitted, this integration will be disabled."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}

var LabelSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"text": {
			Description:  lang.Markdown("Specifies the text of the label that will be displayed in the header of the Web UI."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"background_color": {
			Description:  lang.Markdown("The background color of the label to be displayed. The Web UI defaults to a black background. HEX values may be used"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"text_color": {
			Description:  lang.Markdown("The text color of the label to be displayed. The Web UI defaults to white text. HEX values may be used."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
