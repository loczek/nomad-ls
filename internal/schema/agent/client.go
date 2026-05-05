package agent

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ClientSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"alloc_dir": {
			Description:  lang.Markdown("Specifies the directory to use for allocation data. When this parameter is empty, Nomad will generate the path using the top-level data_dir suffixed with alloc, like \"/opt/nomad/alloc\". This must be an absolute path. Nomad will create the directory on the host, if it does not exist when the agent process starts."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"alloc_mounts_dir": {
			Description:  lang.Markdown("Specifies the directory to use for binding mounts for the unveil file isolation mode. When this parameter is empty, Nomad generates the path as a sibling of the top-level data_dir, with the name alloc_mounts. For example, if the data_dir is /opt/nomad/data, then the alloc mounts directory is /opt/nomad/alloc_mounts. This must be an absolute path and should not be inside the Nomad data directory. Nomad creates the directory on the host, if it does not exist when the agent"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"enabled": {
			Description:  lang.Markdown("Specifies if client mode is enabled. All other client configuration options depend on this value."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"max_kill_timeout": {
			Description:  lang.Markdown("Specifies the maximum amount of time a job is allowed to wait to exit. Individual jobs may customize their own kill timeout, but it may not exceed this value."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("30s")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"disable_remote_exec": {
			Description:  lang.Markdown("Specifies if the client should disable remote task execution to tasks running on this client."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"network_interface": {
			Description:  lang.Markdown("Specifies the name of the interface to force network fingerprinting on. When run in dev mode, this defaults to the loopback interface. When not in dev mode, the interface attached to the default route is used. The scheduler chooses from these fingerprinted IP addresses when allocating ports for tasks. This value supports go-sockaddr/template format. Nomad adds the fingerprinted interface to a \"default\" host_network. If no non-local IP addresses are found, Nomad could"),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"preferred_address_family": {
			Description:  lang.Markdown("Specifies the preferred address family for the network interface. The value can be ipv4 or ipv6. If the selected network interface has both IPv4 and IPv6 addresses, this option will select an IP address of the preferred family. When the option is not specified, the current behavior is conserved: the first IP address is selected no matter the family."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"cpu_disable_dmidecode": {
			Description:  lang.Markdown("Specifies the client should not use dmidecode as a method of cpu detection. Nomad ignores this field on all platforms except Linux."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"cpu_total_compute": {
			Description:  lang.Markdown("Specifies an override for the total CPU compute. This value should be set to # Cores * Core MHz. For example, a quad-core running at 2 GHz would have a total compute of 8000 (4 * 2000). Most clients can determine their total CPU compute automatically, and thus in most cases this should be left unset."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"memory_total_mb": {
			Description:  lang.Markdown("Specifies an override for the total memory. If set, this value overrides any detected memory."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"disk_total_mb": {
			Description:  lang.Markdown("Specifies an override for the total disk space fingerprint attribute. This value is not used by the scheduler unless you have constraints set on the attribute unique.storage.bytestotal. The actual total disk space can be determined via the Read Stats API"),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"disk_free_mb": {
			Description:  lang.Markdown("Previously specified the disk space free for scheduling allocations. If set, it would override any detected free disk space. This value has been deprecated and is now ignored by Nomad clients. Use client.reserved.disk instead."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsDeprecated: true,
			IsOptional:   true,
		},
		"min_dynamic_port": {
			Description:  lang.Markdown("Specifies the minimum dynamic port to be assigned. Individual ports and ranges of ports may be excluded from dynamic port assignment via reserved parameters."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(20000)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"max_dynamic_port": {
			Description:  lang.Markdown("Specifies the maximum dynamic port to be assigned. Individual ports and ranges of ports may be excluded from dynamic port assignment via reserved parameters."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(32000)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"node_class": {
			Description:  lang.Markdown("Specifies an arbitrary string used to logically group client nodes by user-defined class. This value can be used during job placement as an affinity or constraint attribute and other places where variable interpolation is supported."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"node_max_allocs": {
			Description:  lang.Markdown("Specifies the maximum number of allocations that may be scheduled on a client node and is not enforced if unset. This value can be seen in nomad node status under Allocated Resources."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"node_pool": {
			Description:  lang.Markdown("Specifies the node pool in which the client is registered. If the node pool does not exist yet, it will be created automatically if the node registers in the authoritative region. In non-authoritative regions, the node is kept in the initializing status until the node pool is created and replicated."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("default")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"options": {
			Description: lang.Markdown("Specifies a key-value mapping of internal configuration for clients, such as for driver configuration."),
			Constraint:  schema.LiteralType{Type: cty.Map(cty.String)},
			IsOptional:  true,
		},
		"servers": {
			Description:  lang.Markdown("Specifies an array of addresses to the Nomad servers this client should join. This list is used to register the client with the server nodes and advertise the available resources so that the agent can receive work. This may be specified as an IP address or DNS, with or without the port. If the port is omitted, the default port of 4647 is used. If you are specifying IPv6 addresses, they must be in URL format with brackets (ex. \"[2001:db8::1]\")."),
			DefaultValue: schema.DefaultValue{Value: cty.ListValEmpty(cty.String)},
			Constraint:   schema.LiteralType{Type: cty.List(cty.String)},
			IsOptional:   true,
		},
		"state_dir": {
			Description:  lang.Markdown("Specifies the directory to use to store client state. When this parameter is empty, Nomad will generate the path using the top-level data_dir suffixed with client, like \"/opt/nomad/client\". This must be an absolute path. Nomad will create the directory on the host, if it does not exist when the agent process starts."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"gc_interval": {
			Description:  lang.Markdown("Specifies the interval at which Nomad attempts to garbage collect terminal allocation directories."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("1m")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"gc_disk_usage_threshold": {
			Description:  lang.Markdown("Specifies the disk usage percent which Nomad tries to maintain by garbage collecting terminal allocations. Note that Nomad immediately garbage collects terminal allocations if garbage collected on the server."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberFloatVal(80)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"gc_inode_usage_threshold": {
			Description:  lang.Markdown("Specifies the inode usage percent which Nomad tries to maintain by garbage collecting terminal allocations. Note that Nomad immediately garbage collects terminal allocations if garbage collected on the server."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberFloatVal(70)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"gc_max_allocs": {
			Description:  lang.Markdown("Specifies the maximum number of allocations which a client will track before triggering a garbage collection of terminal allocations. This will not limit the number of allocations a node can run at a time, however after gc_max_allocs every new allocation will cause terminal allocations to be GC'd. Note that Nomad immediately garbage collects terminal allocations if garbage collected on the server."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(50)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"gc_parallel_destroys": {
			Description:  lang.Markdown("Specifies the maximum number of parallel destroys allowed by the garbage collector. This value should be relatively low to avoid high resource usage during garbage collections."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(2)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"gc_volumes_on_node_gc": {
			Description:  lang.Markdown("Specifies that the server should delete any dynamic host volumes on this node when the node is garbage collected. You should only set this to true if you know that garbage collected nodes will never rejoin the cluster, such as with ephemeral cloud hosts."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"no_host_uuid": {
			Description:  lang.Markdown("By default a random node UUID will be generated, but setting this to false will use the system's UUID."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(true)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"cni_path": {
			Description:  lang.Markdown("Sets the search path that is used for CNI plugin discovery. Multiple paths can be searched using colon delimited paths. CNI is only supported on Linux."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("/opt/cni/bin:/usr/libexec/cni")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"cni_config_dir": {
			Description:  lang.Markdown("Sets the directory where CNI network configuration is located. The client will use this path when fingerprinting CNI networks. Filenames should use the .conflist extension. Filenames with the .conf or .json extensions are loaded as individual plugin configuration."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("/opt/cni/config")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"bridge_network_name": {
			Description:  lang.Markdown("Sets the name of the bridge to be created by Nomad for allocations running with bridge networking mode on the client."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("nomad")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"bridge_network_subnet": {
			Description:  lang.Markdown("Specifies the subnet which the client will use to allocate IP addresses from."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("172.26.64.0/20")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"bridge_network_subnet_ipv6": {
			Description:  lang.Markdown("Enables IPv6 on Nomad's bridge network by specifying the subnet which the client will use to allocate IPv6 addresses."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"bridge_network_hairpin_mode": {
			Description:  lang.Markdown("Specifies if hairpin mode is enabled on the network bridge created by Nomad for allocations running with bridge networking mode on this client. You may use the corresponding node attribute nomad.bridge.hairpin_mode in constraints. When hairpin mode is enabled, allocations are able to reach their own IP and all ports bound to it. Changing this value requires a reboot of the client host to take effect."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"common_plugin_dir": {
			Description:  lang.Markdown("Specifies the directory where you should place plugins that conform to the common plugins interface. When this parameter is empty, Nomad generates the path using the top-level data-dir suffixed with common_plugins, like \"/opt/nomad/common_plugins\". This must be an absolute path."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"host_volumes_dir": {
			Description:  lang.Markdown("Specifies the directory wherein host volume plugins should place volumes. When this parameter is empty, Nomad generates the path using the top-level data_dir suffixed with host_volumes, like \"/opt/nomad/host_volumes\". This must be an absolute path."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"host_volume_plugin_dir": {
			Description:  lang.Markdown("Specifies the directory to find host volume plugins. When this parameter is empty, Nomad generates the path using the top-level data_dir suffixed with host_volume_plugins, like \"/opt/nomad/host_volume_plugins\". This must be an absolute path."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
		"cgroup_parent": {
			Description:  lang.Markdown("Specifies the cgroup parent for which cgroup subsystems managed by Nomad will be mounted under. Currently this only applies to the cpuset subsystems. This field is ignored on non Linux platforms."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("/nomad")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
	Blocks: map[string]*schema.BlockSchema{
		// TODO: fix this
		"chroot_env": {
			Description: lang.Markdown("Specifies a key-value mapping that defines the chroot environment for jobs using the Exec and Java drivers."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Constraint: schema.LiteralType{Type: cty.String},
					IsOptional: true,
				},
			},
		},
		"meta": {
			Description: lang.Markdown("Specifies a key-value map that annotates with user-defined metadata."),
			Body: &schema.BodySchema{
				AnyAttribute: &schema.AttributeSchema{
					Description: lang.Markdown("A user-defined key-value pair for metadata."),
					Constraint:  schema.LiteralType{Type: cty.String},
					IsOptional:  true,
				},
			},
		},
		"reserved": {
			Description: lang.Markdown("Specifies that Nomad should reserve a portion of the node's resources from receiving tasks. This can be used to target a certain capacity usage for the node. For example, a value equal to 20% of the node's CPU could be reserved to target a CPU utilization of 80%."),
			Body:        ReservedSchema,
		},
		"server_join": {
			Description: lang.Markdown("Specifies how the Nomad client will connect to Nomad servers. The start_join field is not supported on the client. The retry_join fields may directly specify the server address or use go-discover syntax for auto-discovery."),
			Body:        ServerJoinSchema,
		},
		"artifact": {
			Description: lang.Markdown("Specifies controls on the behavior of task artifact blocks."),
			Body:        ArtifactSchema,
		},
		"template": {
			Description: lang.Markdown("Specifies controls on the behavior of task template blocks."),
			Body:        TemplateSchema,
		},
		"fingerprint": {
			Description: lang.Markdown("Provides configuration for fingerprinters used by the client and applies to \"env_aws\", \"env_azure\", \"env_digitalocean\", and \"env_gce\" fingerprinters."),
			Body:        FingerprintSchema,
		},
		"host_volume": {
			Description: lang.Markdown("Exposes paths from the host as volumes that can be mounted into jobs."),
			Body:        HostVolumeSchema,
		},
		"host_network": {
			Description: lang.Markdown("Registers additional host networks with the node that can be selected when port mapping. Although this parameter defaults to nil, Nomad creates a host network named \"default\" using the network_interface field."),
			Body:        HostNetworkSchema,
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
		},
		"drain_on_shutdown": {
			Description: lang.Markdown("Controls the behavior of the client when leave_on_interrupt or leave_on_terminate are set and the client receives the appropriate signal."),
			Body:        DrainOnShutdownSchema,
		},
		"users": {
			Description: lang.Markdown("Specifies options concerning Nomad client's use of operating system users."),
			Body:        UsersSchema,
		},
	},
}

var ReservedSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cpu": {
			Description:  lang.Markdown("Specifies the amount of CPU to reserve, in MHz."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"cores": {
			Description:  lang.Markdown("Specifies the cpuset of CPU cores to reserve. Only supported on Linux."),
			DefaultValue: schema.DefaultValue{Value: cty.BoolVal(false)},
			Constraint:   schema.LiteralType{Type: cty.Bool},
			IsOptional:   true,
		},
		"memory": {
			Description:  lang.Markdown("Specifies the amount of memory to reserve, in MB."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"disk": {
			Description:  lang.Markdown("Specifies the amount of disk to reserve, in MB."),
			DefaultValue: schema.DefaultValue{Value: cty.NumberIntVal(0)},
			Constraint:   schema.LiteralType{Type: cty.Number},
			IsOptional:   true,
		},
		"reserved_ports": {
			Description:  lang.Markdown("Specifies a comma-separated list of ports to reserve on all fingerprinted network devices. Ranges can be specified by using a hyphen separating the two inclusive ends. Refer to host_network for reserving ports on specific host networks."),
			DefaultValue: schema.DefaultValue{Value: cty.StringVal("")},
			Constraint:   schema.LiteralType{Type: cty.String},
			IsOptional:   true,
		},
	},
}
