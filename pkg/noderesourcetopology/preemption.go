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

package noderesourcetopology

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1"
	"k8s.io/klog/v2"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
	fwk "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/preemption"
)

var _ preemption.Interface = &preemptor{}

type preemptor struct {
	logger klog.Logger
	fh     framework.Handle
	state  fwk.CycleState
}

func (p *preemptor) GetOffsetAndNumCandidates(numNodes int32) (int32, int32) {
	return 0, numNodes
}

func (p *preemptor) CandidatesToVictimsMap(candidates []preemption.Candidate) map[string]*extenderv1.Victims {
	m := make(map[string]*extenderv1.Victims, len(candidates))
	for _, c := range candidates {
		m[c.Name()] = c.Victims()
	}
	return m
}

func (p *preemptor) PodEligibleToPreemptOthers(ctx context.Context, pod *corev1.Pod, nominatedNodeStatus *fwk.Status) (bool, string) {
	if pod.Spec.PreemptionPolicy != nil && *pod.Spec.PreemptionPolicy == corev1.PreemptNever {
		return false, "not eligible due to preemptionPolicy=Never."
	}
	// TODO: implement here custom NodeResourceTopology logic:
	// enable preemption (return true) and add NRT-specific eligibility checks.
	return false, "preemption is not yet supported"
}

func (p *preemptor) SelectVictimsOnNode(
	ctx context.Context,
	state fwk.CycleState,
	pod *corev1.Pod,
	nodeInfo fwk.NodeInfo,
	pdbs []*policy.PodDisruptionBudget,
) ([]*corev1.Pod, int, *fwk.Status) {
	// TODO: implement here custom NodeResourceTopology logic:
	// select victims on this node using NUMA-aware resource accounting.
	return nil, 0, fwk.NewStatus(fwk.UnschedulableAndUnresolvable, fmt.Sprintf("preemption is not yet supported on node %v", nodeInfo.Node().Name))
}

func (p *preemptor) OrderedScoreFuncs(ctx context.Context, nodesToVictims map[string]*extenderv1.Victims) []func(node string) int64 {
	return nil
}
