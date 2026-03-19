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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	fwk "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/preemption"

	"sigs.k8s.io/scheduler-plugins/pkg/noderesourcetopology/logging"
)

func (tm *TopologyMatch) PostFilter(ctx context.Context, state fwk.CycleState, pod *corev1.Pod, filteredNodeStatusMap framework.NodeToStatusReader) (*framework.PostFilterResult, *fwk.Status) {
	lh := klog.FromContext(klog.NewContext(ctx, tm.logger)).WithValues(logging.KeyPod, klog.KObj(pod), logging.KeyPodUID, logging.PodUID(pod))
	lh.V(4).Info(logging.FlowBegin)
	defer lh.V(4).Info(logging.FlowEnd)

	pe := preemption.NewEvaluator(
		tm.Name(),
		tm.fh,
		&preemptor{
			logger: lh,
			fh:     tm.fh,
			state:  state,
		},
		false,
	)

	return pe.Preempt(ctx, state, pod, filteredNodeStatusMap)
}
