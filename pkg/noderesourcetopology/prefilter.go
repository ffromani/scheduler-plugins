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
	"k8s.io/klog/v2"
	fwk "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	"sigs.k8s.io/scheduler-plugins/pkg/noderesourcetopology/logging"
)

const preFilterStateKey = "PreFilter" + Name

type preFilterState struct {
	// TODO: implement here custom NodeResourceTopology logic:
	// add fields to carry NUMA topology data across the scheduling cycle.
}

func (s *preFilterState) Clone() fwk.StateData {
	return s
}

func getPreFilterState(cycleState fwk.CycleState) (*preFilterState, error) {
	c, err := cycleState.Read(preFilterStateKey)
	if err != nil {
		return nil, fmt.Errorf("error reading %q from cycleState: %w", preFilterStateKey, err)
	}

	s, ok := c.(*preFilterState)
	if !ok {
		return nil, fmt.Errorf("%+v convert to preFilterState error", c)
	}
	return s, nil
}

func (tm *TopologyMatch) PreFilter(ctx context.Context, state fwk.CycleState, pod *corev1.Pod, nodes []fwk.NodeInfo) (*framework.PreFilterResult, *fwk.Status) {
	lh := klog.FromContext(klog.NewContext(ctx, tm.logger)).WithValues(logging.KeyPod, klog.KObj(pod), logging.KeyPodUID, logging.PodUID(pod))
	lh.V(4).Info(logging.FlowBegin)
	defer lh.V(4).Info(logging.FlowEnd)

	// TODO: implement here custom NodeResourceTopology logic:
	// populate preFilterState with a per-cycle NRT snapshot.
	state.Write(preFilterStateKey, &preFilterState{})
	return nil, fwk.NewStatus(fwk.Success)
}

func (tm *TopologyMatch) PreFilterExtensions() framework.PreFilterExtensions {
	return tm
}

func (tm *TopologyMatch) AddPod(ctx context.Context, state fwk.CycleState, podToSchedule *corev1.Pod, podInfoToAdd fwk.PodInfo, nodeInfo fwk.NodeInfo) *fwk.Status {
	// TODO: implement here custom NodeResourceTopology logic:
	// update the NUMA topology state when simulating a pod addition on a node.
	return fwk.NewStatus(fwk.Success)
}

func (tm *TopologyMatch) RemovePod(ctx context.Context, state fwk.CycleState, podToSchedule *corev1.Pod, podInfoToRemove fwk.PodInfo, nodeInfo fwk.NodeInfo) *fwk.Status {
	// TODO: implement here custom NodeResourceTopology logic:
	// update the NUMA topology state when simulating a pod removal from a node.
	return fwk.NewStatus(fwk.Success)
}
