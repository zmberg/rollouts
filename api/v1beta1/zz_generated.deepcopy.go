//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/apps/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchRelease) DeepCopyInto(out *BatchRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchRelease.
func (in *BatchRelease) DeepCopy() *BatchRelease {
	if in == nil {
		return nil
	}
	out := new(BatchRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseCanaryStatus) DeepCopyInto(out *BatchReleaseCanaryStatus) {
	*out = *in
	if in.BatchReadyTime != nil {
		in, out := &in.BatchReadyTime, &out.BatchReadyTime
		*out = (*in).DeepCopy()
	}
	if in.NoNeedUpdateReplicas != nil {
		in, out := &in.NoNeedUpdateReplicas, &out.NoNeedUpdateReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseCanaryStatus.
func (in *BatchReleaseCanaryStatus) DeepCopy() *BatchReleaseCanaryStatus {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseCanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseList) DeepCopyInto(out *BatchReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BatchRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseList.
func (in *BatchReleaseList) DeepCopy() *BatchReleaseList {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseSpec) DeepCopyInto(out *BatchReleaseSpec) {
	*out = *in
	out.WorkloadRef = in.WorkloadRef
	in.ReleasePlan.DeepCopyInto(&out.ReleasePlan)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseSpec.
func (in *BatchReleaseSpec) DeepCopy() *BatchReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseStatus) DeepCopyInto(out *BatchReleaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RolloutCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CanaryStatus.DeepCopyInto(&out.CanaryStatus)
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseStatus.
func (in *BatchReleaseStatus) DeepCopy() *BatchReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
	*out = *in
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStep) DeepCopyInto(out *CanaryStep) {
	*out = *in
	in.TrafficRoutingStrategy.DeepCopyInto(&out.TrafficRoutingStrategy)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Pause.DeepCopyInto(&out.Pause)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStep.
func (in *CanaryStep) DeepCopy() *CanaryStep {
	if in == nil {
		return nil
	}
	out := new(CanaryStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStrategy) DeepCopyInto(out *CanaryStrategy) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]CanaryStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrafficRoutings != nil {
		in, out := &in.TrafficRoutings, &out.TrafficRoutings
		*out = make([]TrafficRoutingRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.PatchPodTemplateMetadata != nil {
		in, out := &in.PatchPodTemplateMetadata, &out.PatchPodTemplateMetadata
		*out = new(PatchPodTemplateMetadata)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStrategy.
func (in *CanaryStrategy) DeepCopy() *CanaryStrategy {
	if in == nil {
		return nil
	}
	out := new(CanaryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentExtraStatus) DeepCopyInto(out *DeploymentExtraStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentExtraStatus.
func (in *DeploymentExtraStatus) DeepCopy() *DeploymentExtraStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentExtraStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStrategy) DeepCopyInto(out *DeploymentStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(v1.RollingUpdateDeployment)
		(*in).DeepCopyInto(*out)
	}
	out.Partition = in.Partition
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStrategy.
func (in *DeploymentStrategy) DeepCopy() *DeploymentStrategy {
	if in == nil {
		return nil
	}
	out := new(DeploymentStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayTrafficRouting) DeepCopyInto(out *GatewayTrafficRouting) {
	*out = *in
	if in.HTTPRouteName != nil {
		in, out := &in.HTTPRouteName, &out.HTTPRouteName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayTrafficRouting.
func (in *GatewayTrafficRouting) DeepCopy() *GatewayTrafficRouting {
	if in == nil {
		return nil
	}
	out := new(GatewayTrafficRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpRouteMatch) DeepCopyInto(out *HttpRouteMatch) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]v1alpha2.HTTPHeaderMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpRouteMatch.
func (in *HttpRouteMatch) DeepCopy() *HttpRouteMatch {
	if in == nil {
		return nil
	}
	out := new(HttpRouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTrafficRouting) DeepCopyInto(out *IngressTrafficRouting) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTrafficRouting.
func (in *IngressTrafficRouting) DeepCopy() *IngressTrafficRouting {
	if in == nil {
		return nil
	}
	out := new(IngressTrafficRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchPodTemplateMetadata) DeepCopyInto(out *PatchPodTemplateMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchPodTemplateMetadata.
func (in *PatchPodTemplateMetadata) DeepCopy() *PatchPodTemplateMetadata {
	if in == nil {
		return nil
	}
	out := new(PatchPodTemplateMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBatch) DeepCopyInto(out *ReleaseBatch) {
	*out = *in
	out.CanaryReplicas = in.CanaryReplicas
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBatch.
func (in *ReleaseBatch) DeepCopy() *ReleaseBatch {
	if in == nil {
		return nil
	}
	out := new(ReleaseBatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlan) DeepCopyInto(out *ReleasePlan) {
	*out = *in
	if in.Batches != nil {
		in, out := &in.Batches, &out.Batches
		*out = make([]ReleaseBatch, len(*in))
		copy(*out, *in)
	}
	if in.BatchPartition != nil {
		in, out := &in.BatchPartition, &out.BatchPartition
		*out = new(int32)
		**out = **in
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.PatchPodTemplateMetadata != nil {
		in, out := &in.PatchPodTemplateMetadata, &out.PatchPodTemplateMetadata
		*out = new(PatchPodTemplateMetadata)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlan.
func (in *ReleasePlan) DeepCopy() *ReleasePlan {
	if in == nil {
		return nil
	}
	out := new(ReleasePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rollout) DeepCopyInto(out *Rollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rollout.
func (in *Rollout) DeepCopy() *Rollout {
	if in == nil {
		return nil
	}
	out := new(Rollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutCondition) DeepCopyInto(out *RolloutCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutCondition.
func (in *RolloutCondition) DeepCopy() *RolloutCondition {
	if in == nil {
		return nil
	}
	out := new(RolloutCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutList) DeepCopyInto(out *RolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutList.
func (in *RolloutList) DeepCopy() *RolloutList {
	if in == nil {
		return nil
	}
	out := new(RolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutPause) DeepCopyInto(out *RolloutPause) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutPause.
func (in *RolloutPause) DeepCopy() *RolloutPause {
	if in == nil {
		return nil
	}
	out := new(RolloutPause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutSpec) DeepCopyInto(out *RolloutSpec) {
	*out = *in
	out.WorkloadRef = in.WorkloadRef
	in.Strategy.DeepCopyInto(&out.Strategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutSpec.
func (in *RolloutSpec) DeepCopy() *RolloutSpec {
	if in == nil {
		return nil
	}
	out := new(RolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStatus) DeepCopyInto(out *RolloutStatus) {
	*out = *in
	if in.CanaryStatus != nil {
		in, out := &in.CanaryStatus, &out.CanaryStatus
		*out = new(CanaryStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RolloutCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStatus.
func (in *RolloutStatus) DeepCopy() *RolloutStatus {
	if in == nil {
		return nil
	}
	out := new(RolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(CanaryStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficRoutingRef) DeepCopyInto(out *TrafficRoutingRef) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressTrafficRouting)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(GatewayTrafficRouting)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomNetworkRefs != nil {
		in, out := &in.CustomNetworkRefs, &out.CustomNetworkRefs
		*out = make([]ObjectRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficRoutingRef.
func (in *TrafficRoutingRef) DeepCopy() *TrafficRoutingRef {
	if in == nil {
		return nil
	}
	out := new(TrafficRoutingRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficRoutingStrategy) DeepCopyInto(out *TrafficRoutingStrategy) {
	*out = *in
	if in.Traffic != nil {
		in, out := &in.Traffic, &out.Traffic
		*out = new(string)
		**out = **in
	}
	if in.RequestHeaderModifier != nil {
		in, out := &in.RequestHeaderModifier, &out.RequestHeaderModifier
		*out = new(v1alpha2.HTTPRequestHeaderFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make([]HttpRouteMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficRoutingStrategy.
func (in *TrafficRoutingStrategy) DeepCopy() *TrafficRoutingStrategy {
	if in == nil {
		return nil
	}
	out := new(TrafficRoutingStrategy)
	in.DeepCopyInto(out)
	return out
}
