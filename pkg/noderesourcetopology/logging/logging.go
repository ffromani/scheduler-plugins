/*
Copyright 2021 The Kubernetes Authors.

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

package logging

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2/klogr"

	"github.com/go-logr/logr"
	topologyv1alpha2 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha2"
)

var (
	nrtlog logr.Logger
)

func init() {
	nrtlog = klogr.NewWithOptions(klogr.WithFormat(klogr.FormatKlog))
}

func Set(handle logr.Logger) {
	nrtlog = handle
}

func Get() logr.Logger {
	return nrtlog
}

func WithName(name string) logr.Logger {
	return nrtlog.WithName(name)
}

func PodRef(pod *corev1.Pod) string {
	if pod == nil {
		return ""
	}
	return pod.Namespace + "/" + pod.Name
}

func NRT(lh logr.Logger, desc string, nrtObj *topologyv1alpha2.NodeResourceTopology) {
	if !lh.V(6).Enabled() {
		// avoid the expensive marshal operation
		return
	}

	ntrJson, err := json.MarshalIndent(nrtObj, "", " ")
	if err != nil {
		lh.V(6).Error(err, "failed to marshal noderesourcetopology object")
		return
	}
	lh.V(6).Info(desc, "noderesourcetopology", string(ntrJson))
}
