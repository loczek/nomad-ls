package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var TLSSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"ca_file": {
			Description:  lang.Markdown("Specifies the path to the CA certificate to use for Nomad's TLS communication."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"cert_file": {
			Description:  lang.Markdown("Specifies the path to the certificate file used for Nomad's TLS communication."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"key_file": {
			Description:  lang.Markdown("Specifies the path to the key file to use for Nomad's TLS communication."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"http": {
			Description:  lang.Markdown("Specifies if TLS should be enabled on the HTTP endpoints on the Nomad agent, including the API. By default this is non-mutual TLS. You can upgrade this to mTLS by setting verify_https_client=true, but this can complicate using the Nomad UI by requiring mTLS in your browser or running a proxy in front of the Nomad UI."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"rpc": {
			Description:  lang.Markdown("Toggle the option to enable mTLS on the RPC endpoints and Raft traffic. When this setting is activated, it establishes protection both between Nomad servers and from the clients back to the servers, ensuring mutual authentication. Setting rpc=true is required for secure operation of Nomad."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"rpc_upgrade_mode": {
			Description:  lang.Markdown("This option should be used only when the cluster is being upgraded to TLS, and removed after the migration is complete. This allows the agent to accept both TLS and plaintext traffic."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"tls_cipher_suites": {
			Description:  lang.Markdown("Specifies the TLS cipher suites that will be used by the agent as a comma-separated string. Known insecure ciphers are disabled (3DES and RC4). By default, an agent is configured to use TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 and TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"tls_min_version": {
			Description:  lang.Markdown("Specifies the minimum supported version of TLS. Accepted values are \"tls10\", \"tls11\", \"tls12\", \"tls13\"."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("tls12")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"verify_https_client": {
			Description:  lang.Markdown("Specifies agents should require client certificates for all incoming HTTPS requests, effectively upgrading tls.http=true to mTLS. The client certificates must be signed by the same CA as Nomad. By default, verify_https_client is set to false, which is safe so long as ACLs are enabled. This is recommended if you are using the Nomad web UI to avoid the difficulty of distributing client certs to browsers."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"verify_server_hostname": {
			Description:  lang.Markdown("Specifies if outgoing TLS connections should verify the server's hostname."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
	},
}
