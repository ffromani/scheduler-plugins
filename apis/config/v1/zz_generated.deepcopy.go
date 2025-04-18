//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1 "k8s.io/kube-scheduler/config/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoschedulingArgs) DeepCopyInto(out *CoschedulingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PermitWaitingTimeSeconds != nil {
		in, out := &in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.PodGroupBackoffSeconds != nil {
		in, out := &in.PodGroupBackoffSeconds, &out.PodGroupBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoschedulingArgs.
func (in *CoschedulingArgs) DeepCopy() *CoschedulingArgs {
	if in == nil {
		return nil
	}
	out := new(CoschedulingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoschedulingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadVariationRiskBalancingArgs) DeepCopyInto(out *LoadVariationRiskBalancingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.TrimaranSpec.DeepCopyInto(&out.TrimaranSpec)
	if in.SafeVarianceMargin != nil {
		in, out := &in.SafeVarianceMargin, &out.SafeVarianceMargin
		*out = new(float64)
		**out = **in
	}
	if in.SafeVarianceSensitivity != nil {
		in, out := &in.SafeVarianceSensitivity, &out.SafeVarianceSensitivity
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadVariationRiskBalancingArgs.
func (in *LoadVariationRiskBalancingArgs) DeepCopy() *LoadVariationRiskBalancingArgs {
	if in == nil {
		return nil
	}
	out := new(LoadVariationRiskBalancingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadVariationRiskBalancingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LowRiskOverCommitmentArgs) DeepCopyInto(out *LowRiskOverCommitmentArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.TrimaranSpec.DeepCopyInto(&out.TrimaranSpec)
	if in.SmoothingWindowSize != nil {
		in, out := &in.SmoothingWindowSize, &out.SmoothingWindowSize
		*out = new(int64)
		**out = **in
	}
	if in.RiskLimitWeights != nil {
		in, out := &in.RiskLimitWeights, &out.RiskLimitWeights
		*out = make(map[corev1.ResourceName]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LowRiskOverCommitmentArgs.
func (in *LowRiskOverCommitmentArgs) DeepCopy() *LowRiskOverCommitmentArgs {
	if in == nil {
		return nil
	}
	out := new(LowRiskOverCommitmentArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LowRiskOverCommitmentArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricProviderSpec) DeepCopyInto(out *MetricProviderSpec) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(string)
		**out = **in
	}
	if in.InsecureSkipVerify != nil {
		in, out := &in.InsecureSkipVerify, &out.InsecureSkipVerify
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricProviderSpec.
func (in *MetricProviderSpec) DeepCopy() *MetricProviderSpec {
	if in == nil {
		return nil
	}
	out := new(MetricProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkOverheadArgs) DeepCopyInto(out *NetworkOverheadArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WeightsName != nil {
		in, out := &in.WeightsName, &out.WeightsName
		*out = new(string)
		**out = **in
	}
	if in.NetworkTopologyName != nil {
		in, out := &in.NetworkTopologyName, &out.NetworkTopologyName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkOverheadArgs.
func (in *NetworkOverheadArgs) DeepCopy() *NetworkOverheadArgs {
	if in == nil {
		return nil
	}
	out := new(NetworkOverheadArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkOverheadArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceTopologyCache) DeepCopyInto(out *NodeResourceTopologyCache) {
	*out = *in
	if in.ForeignPodsDetect != nil {
		in, out := &in.ForeignPodsDetect, &out.ForeignPodsDetect
		*out = new(ForeignPodsDetectMode)
		**out = **in
	}
	if in.ResyncMethod != nil {
		in, out := &in.ResyncMethod, &out.ResyncMethod
		*out = new(CacheResyncMethod)
		**out = **in
	}
	if in.InformerMode != nil {
		in, out := &in.InformerMode, &out.InformerMode
		*out = new(CacheInformerMode)
		**out = **in
	}
	if in.ResyncScope != nil {
		in, out := &in.ResyncScope, &out.ResyncScope
		*out = new(CacheResyncScope)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceTopologyCache.
func (in *NodeResourceTopologyCache) DeepCopy() *NodeResourceTopologyCache {
	if in == nil {
		return nil
	}
	out := new(NodeResourceTopologyCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceTopologyMatchArgs) DeepCopyInto(out *NodeResourceTopologyMatchArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ScoringStrategy != nil {
		in, out := &in.ScoringStrategy, &out.ScoringStrategy
		*out = new(ScoringStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheResyncPeriodSeconds != nil {
		in, out := &in.CacheResyncPeriodSeconds, &out.CacheResyncPeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(NodeResourceTopologyCache)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceTopologyMatchArgs.
func (in *NodeResourceTopologyMatchArgs) DeepCopy() *NodeResourceTopologyMatchArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourceTopologyMatchArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourceTopologyMatchArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesAllocatableArgs) DeepCopyInto(out *NodeResourcesAllocatableArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]configv1.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesAllocatableArgs.
func (in *NodeResourcesAllocatableArgs) DeepCopy() *NodeResourcesAllocatableArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesAllocatableArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesAllocatableArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeaksArgs) DeepCopyInto(out *PeaksArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.NodePowerModel != nil {
		in, out := &in.NodePowerModel, &out.NodePowerModel
		*out = make(map[string]PowerModel, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeaksArgs.
func (in *PeaksArgs) DeepCopy() *PeaksArgs {
	if in == nil {
		return nil
	}
	out := new(PeaksArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeaksArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerModel) DeepCopyInto(out *PowerModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerModel.
func (in *PowerModel) DeepCopy() *PowerModel {
	if in == nil {
		return nil
	}
	out := new(PowerModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreemptionTolerationArgs) DeepCopyInto(out *PreemptionTolerationArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.MinCandidateNodesPercentage != nil {
		in, out := &in.MinCandidateNodesPercentage, &out.MinCandidateNodesPercentage
		*out = new(int32)
		**out = **in
	}
	if in.MinCandidateNodesAbsolute != nil {
		in, out := &in.MinCandidateNodesAbsolute, &out.MinCandidateNodesAbsolute
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreemptionTolerationArgs.
func (in *PreemptionTolerationArgs) DeepCopy() *PreemptionTolerationArgs {
	if in == nil {
		return nil
	}
	out := new(PreemptionTolerationArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreemptionTolerationArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoringStrategy) DeepCopyInto(out *ScoringStrategy) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]configv1.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringStrategy.
func (in *ScoringStrategy) DeepCopy() *ScoringStrategy {
	if in == nil {
		return nil
	}
	out := new(ScoringStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SySchedArgs) DeepCopyInto(out *SySchedArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DefaultProfileNamespace != nil {
		in, out := &in.DefaultProfileNamespace, &out.DefaultProfileNamespace
		*out = new(string)
		**out = **in
	}
	if in.DefaultProfileName != nil {
		in, out := &in.DefaultProfileName, &out.DefaultProfileName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SySchedArgs.
func (in *SySchedArgs) DeepCopy() *SySchedArgs {
	if in == nil {
		return nil
	}
	out := new(SySchedArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SySchedArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetLoadPackingArgs) DeepCopyInto(out *TargetLoadPackingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.TrimaranSpec.DeepCopyInto(&out.TrimaranSpec)
	if in.DefaultRequests != nil {
		in, out := &in.DefaultRequests, &out.DefaultRequests
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.DefaultRequestsMultiplier != nil {
		in, out := &in.DefaultRequestsMultiplier, &out.DefaultRequestsMultiplier
		*out = new(string)
		**out = **in
	}
	if in.TargetUtilization != nil {
		in, out := &in.TargetUtilization, &out.TargetUtilization
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetLoadPackingArgs.
func (in *TargetLoadPackingArgs) DeepCopy() *TargetLoadPackingArgs {
	if in == nil {
		return nil
	}
	out := new(TargetLoadPackingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetLoadPackingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopologicalSortArgs) DeepCopyInto(out *TopologicalSortArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologicalSortArgs.
func (in *TopologicalSortArgs) DeepCopy() *TopologicalSortArgs {
	if in == nil {
		return nil
	}
	out := new(TopologicalSortArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TopologicalSortArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrimaranSpec) DeepCopyInto(out *TrimaranSpec) {
	*out = *in
	in.MetricProvider.DeepCopyInto(&out.MetricProvider)
	if in.WatcherAddress != nil {
		in, out := &in.WatcherAddress, &out.WatcherAddress
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrimaranSpec.
func (in *TrimaranSpec) DeepCopy() *TrimaranSpec {
	if in == nil {
		return nil
	}
	out := new(TrimaranSpec)
	in.DeepCopyInto(out)
	return out
}
