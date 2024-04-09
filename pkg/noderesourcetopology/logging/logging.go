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

	"github.com/go-logr/logr"

	topologyv1alpha2 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha2"
)

// before to replace with FromContext(), at least in filter and score,
// we would need a way to inject a logger instance (preferably a
// per-plugin logger instance) when we create the Scheduler
// (with app.NewSchedulerCommand)

var logh logr.Logger

func SetLogger(lh logr.Logger) {
	logh = lh
}

func Log() logr.Logger {
	return logh
}

func NRT(desc string, nrtObj *topologyv1alpha2.NodeResourceTopology) {
	lh := Log()
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
