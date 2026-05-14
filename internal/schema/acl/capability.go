package acl

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	schemautils "github.com/loczek/nomad-ls/internal/schemaUtils"
	"github.com/zclconf/go-cty/cty"
)

var DenyCapability = schema.LiteralValue{
	Value:       cty.StringVal("deny"),
	Description: lang.Markdown("When multiple policies are associated with a token, deny will take precedence and prevent any capabilities."),
}

var ListJobsCapability = schema.LiteralValue{
	Value:       cty.StringVal("list-jobs"),
	Description: lang.Markdown("Allows listing the jobs and seeing coarse grain status. This implicitly grants csi-list-volume."),
}

var ParseJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("parse-job"),
	Description: lang.Markdown("Allows parsing a job from HCL to JSON."),
}

var ReadJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("read-job"),
	Description: lang.Markdown("Allows inspecting a job and seeing fine grain status. This implicitly grants csi-read-volume."),
}

var SubmitJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("submit-job"),
	Description: lang.Markdown("Allows jobs to be submitted, updated, or stopped."),
}

var DispatchJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("dispatch-job"),
	Description: lang.Markdown("Allows jobs to be dispatched"),
}

var ReadLogsCapability = schema.LiteralValue{
	Value:       cty.StringVal("read-logs"),
	Description: lang.Markdown("Allows the logs associated with a job to be viewed."),
}

var ReadFsCapability = schema.LiteralValue{
	Value:       cty.StringVal("read-fs"),
	Description: lang.Markdown("Allows the filesystem of allocations associated to be viewed. Implicitly grants read-logs."),
}

var AllocExecCapability = schema.LiteralValue{
	Value:       cty.StringVal("alloc-exec"),
	Description: lang.Markdown("Allows an operator to connect and run commands in running allocations."),
}

var AllocNodeExecCapability = schema.LiteralValue{
	Value:       cty.StringVal("alloc-node-exec"),
	Description: lang.Markdown("Allows an operator to connect and run commands in allocations running without filesystem isolation, for example, raw_exec jobs."),
}

var AllocLifecycleCapability = schema.LiteralValue{
	Value:       cty.StringVal("alloc-lifecycle"),
	Description: lang.Markdown("Allows an operator to stop individual allocations manually."),
}

var CsiRegisterPluginCapability = schema.LiteralValue{
	Value:       cty.StringVal("csi-register-plugin"),
	Description: lang.Markdown("Allows jobs to be submitted that register themselves as CSI plugins."),
}

var CsiWriteVolumeCapability = schema.LiteralValue{
	Value:       cty.StringVal("csi-write-volume"),
	Description: lang.Markdown("Allows CSI volumes to be registered or deregistered. This implicitly grants csi-read-volume."),
}

var CsiReadVolumeCapability = schema.LiteralValue{
	Value:       cty.StringVal("csi-read-volume"),
	Description: lang.Markdown("Allows inspecting a CSI volume, seeing fine grain status, and listing external volumes and snapshots. This implicitly grants csi-list-volume."),
}

var CsiListVolumeCapability = schema.LiteralValue{
	Value:       cty.StringVal("csi-list-volume"),
	Description: lang.Markdown("Allows listing CSI volumes, seeing coarse grain status, and listing external volumes and snapshots."),
}

var CsiMountVolumeCapability = schema.LiteralValue{
	Value:       cty.StringVal("csi-mount-volume"),
	Description: lang.Markdown("Allows jobs to be submitted that claim a CSI volume. This implicitly grants csi-read-volume."),
}

var HostVolumeCreateCapability = schema.LiteralValue{
	Value:       cty.StringVal("host-volume-create"),
	Description: lang.Markdown("Allows creating dynamic host volumes. This implicitly grants host-volume-read."),
}

var HostVolumeDeleteCapability = schema.LiteralValue{
	Value:       cty.StringVal("host-volume-delete"),
	Description: lang.Markdown("Allows deleting dynamic host volumes."),
}

var HostVolumeReadCapability = schema.LiteralValue{
	Value:       cty.StringVal("host-volume-read"),
	Description: lang.Markdown("Allows inspecting dynamic host volumes."),
}

var HostVolumeRegisterCapability = schema.LiteralValue{
	Value:       cty.StringVal("host-volume-register"),
	Description: lang.Markdown("Allows registering dynamic host volumes that have been created without a plugin. This implicitly grants host-volume-read and host-volume-create."),
}

var HostVolumeWriteCapability = schema.LiteralValue{
	Value:       cty.StringVal("host-volume-write"),
	Description: lang.Markdown("Allows all write operations on dynamic host volumes. This implicitly grants host-volume-read, host-volume-create, host-volume-register, and host-volume-delete."),
}

var ListScalingPoliciesCapability = schema.LiteralValue{
	Value:       cty.StringVal("list-scaling-policies"),
	Description: lang.Markdown("Allows listing scaling policies."),
}

var ReadScalingPolicyCapability = schema.LiteralValue{
	Value:       cty.StringVal("read-scaling-policy"),
	Description: lang.Markdown("Allows inspecting a scaling policy."),
}

var ReadJobScalingCapability = schema.LiteralValue{
	Value:       cty.StringVal("read-job-scaling"),
	Description: lang.Markdown("Allows inspecting the current scaling of a job."),
}

var ScaleJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("scale-job"),
	Description: lang.Markdown("Allows scaling a job up or down."),
}

var SentinelOverrideCapability = schema.LiteralValue{
	Value:       cty.StringVal("sentinel-override"),
	Description: lang.Markdown("Allows soft mandatory policies to be overridden."),
}

var SubmitRecommendationCapability = schema.LiteralValue{
	Value:       cty.StringVal("submit-recommendation"),
	Description: lang.Markdown("Allows submitting vertical job scaling recommendations."),
}

// https://developer.hashicorp.com/nomad/docs/secure/acl/policies#finer-grained-capabilities
var RegisterJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("register-job"),
	Description: lang.Markdown("Allows a new job to be run or registered."),
}
var RevertJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("revert-job"),
	Description: lang.Markdown("Allows a job to be reverted to a previous version."),
}
var DeregisterJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("deregister-job"),
	Description: lang.Markdown("Allows a job to be stopped, but does not allow manual purging. The job will still be purged during of garbage collection."),
}
var PurgeJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("purge-job"),
	Description: lang.Markdown("Allows a job to be stopped and purged."),
}
var EvaluateJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("evaluate-job"),
	Description: lang.Markdown("Allows an evaluation to be made for a job."),
}
var PlanJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("plan-job"),
	Description: lang.Markdown("Allows a job to be planned."),
}
var TagJobVersionCapability = schema.LiteralValue{
	Value:       cty.StringVal("tag-job-version"),
	Description: lang.Markdown("Allows a job version to be tagged."),
}
var StableJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("stable-job"),
	Description: lang.Markdown("Allows a job version to be marked stable."),
}
var FailDeploymentCapability = schema.LiteralValue{
	Value:       cty.StringVal("fail-deployment"),
	Description: lang.Markdown("Allows a deployment to be failed."),
}
var PauseDeploymentCapability = schema.LiteralValue{
	Value:       cty.StringVal("pause-deployment"),
	Description: lang.Markdown("Allows a deployment to be paused."),
}
var PromoteDeploymentCapability = schema.LiteralValue{
	Value:       cty.StringVal("promote-deployment"),
	Description: lang.Markdown("Allows a deployment to be promoted."),
}
var UnblockDeploymentCapability = schema.LiteralValue{
	Value:       cty.StringVal("unblock-deployment"),
	Description: lang.Markdown("Allows a deployment to be unblocked."),
}
var CancelDeploymentCapability = schema.LiteralValue{
	Value:       cty.StringVal("cancel-deployment"),
	Description: lang.Markdown("Allows a deployment to be cancelled."),
}
var SetAllocHealthDeploymentCapability = schema.LiteralValue{
	Value:       cty.StringVal("set-alloc-health-deployment"),
	Description: lang.Markdown("Allows the health of an allocation in a deployment to be set manually."),
}
var ForcePeriodicJobCapability = schema.LiteralValue{
	Value:       cty.StringVal("force-periodic-job"),
	Description: lang.Markdown("Allows a periodic job to be manually paused." + schemautils.Divider + schemautils.EnterpriseOnly),
}
var PauseAllocationCapability = schema.LiteralValue{
	Value:       cty.StringVal("pause-allocation"),
	Description: lang.Markdown("Allows an allocation to be paused." + schemautils.Divider + schemautils.EnterpriseOnly),
}
var GcAllocationCapability = schema.LiteralValue{
	Value:       cty.StringVal("gc-allocation"),
	Description: lang.Markdown("Allows an allocation to be manually garbage collected, usually used by operators."),
}
var DeleteServiceRegistrationCapability = schema.LiteralValue{
	Value:       cty.StringVal("delete-service-registration"),
	Description: lang.Markdown("Allows an individual service registration to be deleted, usually used by operators."),
}
