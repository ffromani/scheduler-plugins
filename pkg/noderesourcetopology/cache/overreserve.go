/*
Copyright 2022 The Kubernetes Authors.

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

package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	podlisterv1 "k8s.io/client-go/listers/core/v1"
	k8scache "k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	"github.com/go-logr/logr"
	topologyv1alpha2 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha2"
	listerv1alpha2 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/listers/topology/v1alpha2"

	apiconfig "sigs.k8s.io/scheduler-plugins/apis/config"
	nrtlog "sigs.k8s.io/scheduler-plugins/pkg/noderesourcetopology/logging"
	"sigs.k8s.io/scheduler-plugins/pkg/noderesourcetopology/stringify"

	"github.com/k8stopologyawareschedwg/podfingerprint"
)

type OverReserve struct {
	lock             sync.Mutex
	nrts             *nrtStore
	assumedResources map[string]*resourceStore // nodeName -> resourceStore
	// nodesMaybeOverreserved counts how many times a node is filtered out. This is used as trigger condition to try
	// to resync nodes. See The documentation of Resync() below for more details.
	nodesMaybeOverreserved counter
	nodesWithForeignPods   counter
	nrtLister              listerv1alpha2.NodeResourceTopologyLister
	podLister              podlisterv1.PodLister
	resyncMethod           apiconfig.CacheResyncMethod
	lh                     logr.Logger
}

func NewOverReserve(cfg *apiconfig.NodeResourceTopologyCache, nrtLister listerv1alpha2.NodeResourceTopologyLister, podLister podlisterv1.PodLister, lh logr.Logger) (*OverReserve, error) {
	if nrtLister == nil || podLister == nil {
		return nil, fmt.Errorf("nrtcache: received nil references")
	}

	resyncMethod := getCacheResyncMethod(cfg)

	nrtObjs, err := nrtLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	lh = lh.WithName("nrtcache")

	lh.V(3).Info("initializing", "objects", len(nrtObjs), "method", resyncMethod)
	obj := &OverReserve{
		nrts:                   newNrtStore(nrtObjs, lh),
		assumedResources:       make(map[string]*resourceStore),
		nodesMaybeOverreserved: newCounter(),
		nodesWithForeignPods:   newCounter(),
		nrtLister:              nrtLister,
		podLister:              podLister,
		resyncMethod:           resyncMethod,
		lh:                     lh,
	}
	return obj, nil
}

func (ov *OverReserve) GetCachedNRTCopy(nodeName string, pod *corev1.Pod) (*topologyv1alpha2.NodeResourceTopology, bool) {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	if ov.nodesWithForeignPods.IsSet(nodeName) {
		return nil, false
	}

	nrt := ov.nrts.GetNRTCopyByNodeName(nodeName)
	if nrt == nil {
		return nil, true
	}
	nodeAssumedResources, ok := ov.assumedResources[nodeName]
	if !ok {
		return nrt, true
	}

	ov.lh.V(6).Info("NRT", "logID", nrtlog.PodRef(pod), "vanilla", stringify.NodeResourceTopologyResources(nrt))
	nodeAssumedResources.UpdateNRT(nrtlog.PodRef(pod), nrt)

	ov.lh.V(5).Info("NRT", "logID", nrtlog.PodRef(pod), "updated", stringify.NodeResourceTopologyResources(nrt))
	return nrt, true
}

func (ov *OverReserve) NodeMaybeOverReserved(nodeName string, pod *corev1.Pod) {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	val := ov.nodesMaybeOverreserved.Incr(nodeName)
	ov.lh.V(4).Info("mark discarded", "logID", nrtlog.PodRef(pod), "node", nodeName, "count", val)
}

func (ov *OverReserve) NodeHasForeignPods(nodeName string, pod *corev1.Pod) {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	if !ov.nrts.Contains(nodeName) {
		ov.lh.V(5).Info("ignoring foreign pods", "logID", nrtlog.PodRef(pod), "node", nodeName, "nrtinfo", "missing")
		return
	}
	val := ov.nodesWithForeignPods.Incr(nodeName)
	ov.lh.V(4).Info("marked with foreign pods", "logID", nrtlog.PodRef(pod), "node", nodeName, "count", val)
}

func (ov *OverReserve) ReserveNodeResources(nodeName string, pod *corev1.Pod) {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	nodeAssumedResources, ok := ov.assumedResources[nodeName]
	if !ok {
		nodeAssumedResources = newResourceStore(ov.lh)
		ov.assumedResources[nodeName] = nodeAssumedResources
	}

	nodeAssumedResources.AddPod(pod)
	ov.lh.V(5).Info("post reserve", "logID", nrtlog.PodRef(pod), "node", nodeName, "assumedResources", nodeAssumedResources.String())

	ov.nodesMaybeOverreserved.Delete(nodeName)
	ov.lh.V(6).Info("reset discard counter", "logID", nrtlog.PodRef(pod), "node", nodeName)
}

func (ov *OverReserve) UnreserveNodeResources(nodeName string, pod *corev1.Pod) {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	nodeAssumedResources, ok := ov.assumedResources[nodeName]
	if !ok {
		// this should not happen, so we're vocal about it
		// we don't return error because not much to do to recover anyway
		ov.lh.V(3).Info("no resources tracked", "logID", nrtlog.PodRef(pod), "node", nodeName)
		return
	}

	nodeAssumedResources.DeletePod(pod)
	ov.lh.V(5).Info("post release", "logID", nrtlog.PodRef(pod), "node", nodeName, "assumedResources", nodeAssumedResources.String())
}

// NodesMaybeOverReserved returns a slice of all the node names which have been discarded previously,
// so which are supposed to be `dirty` in the cache.
// A node can be discarded for two reasons:
// 1. it legitmately cannot fit containers because it has not enough free resources
// 2. it was pessimistically overallocated, so the node is a candidate for resync
// This function enables the caller to know the slice of nodes should be considered for resync,
// avoiding the need to rescan the full node list.
func (ov *OverReserve) NodesMaybeOverReserved(lh logr.Logger) []string {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	// this is intentionally aggressive. We don't yet make any attempt to find out if the
	// node was discarded because pessimistically overrserved (which should indeed trigger
	// a resync) or if it was discarded because the actual resources on the node really were
	// exhausted. We do like this because this is the safest approach. We will optimize
	// the node selection logic later on to make the resync procedure less aggressive but
	// still correct.
	nodes := ov.nodesWithForeignPods.Clone()
	foreignCount := nodes.Len()

	for _, node := range ov.nodesMaybeOverreserved.Keys() {
		nodes.Incr(node)
	}

	if nodes.Len() > 0 {
		lh.V(4).Info("found dirty nodes", "foreign", foreignCount, "discarded", nodes.Len()-foreignCount, "total", nodes.Len())
	}
	return nodes.Keys()
}

// Resync implements the cache resync loop step. This function checks if the latest available NRT information received matches the
// state of a dirty node, for all the dirty nodes. If this is the case, the cache of a node can be Flush()ed.
// The trigger for attempting to resync a node is not just that we overallocated it. If a node was overallocated but still has capacity,
// we keep using it. But we cannot predict when the capacity is too low, because that would mean predicting the future workload requests.
// The best heuristic found so far is count how many times the node was skipped *AND* crosscheck with its overallocation state.
// If *both* a node has pessimistic overallocation accounted to it *and* was discarded "too many" (how much is too much is a runtime parameter
// which needs to be set and tuned) times, then it becomes a candidate for resync. Just using one of these two factors would lead to
// too aggressive resync attempts, so to more, likely unnecessary, computation work on the scheduler side.
func (ov *OverReserve) Resync() {
	// we are not working with a specific pod, so we need a unique key to track this flow
	logID := logIDFromTime()
	lh_ := ov.lh.WithName("resync").WithValues("logID", logID)

	nodeNames := ov.NodesMaybeOverReserved(lh_)
	// avoid as much as we can unnecessary work and logs.
	if len(nodeNames) == 0 {
		lh_.V(6).Info("no dirty nodes detected")
		return
	}

	// node -> pod identifier (namespace, name)
	nodeToObjsMap, err := makeNodeToPodDataMap(ov.podLister)
	if err != nil {
		lh_.Error(err, "cannot find the mapping between running pods and nodes")
		return
	}

	lh_.V(6).Info("NodeTopology cache sync starting")
	defer lh_.V(6).Info("NodeTopology cache sync complete")

	var nrtUpdates []*topologyv1alpha2.NodeResourceTopology
	for _, nodeName := range nodeNames {
		nlh := lh_.WithValues("node", nodeName)

		nrtCandidate, err := ov.nrtLister.Get(nodeName)
		if err != nil {
			nlh.V(3).Error(err, "failed to get NodeTopology")
			continue
		}
		if nrtCandidate == nil {
			nlh.V(3).Info("missing NodeTopology")
			continue
		}

		objs, ok := nodeToObjsMap[nodeName]
		if !ok {
			// this really should never happen
			nlh.V(3).Info("cannot find any pod for node")
			continue
		}

		pfpExpected, onlyExclRes := podFingerprintForNodeTopology(nrtCandidate, ov.resyncMethod)
		if pfpExpected == "" {
			nlh.V(3).Info("missing NodeTopology podset fingerprint data")
			continue
		}

		nlh.V(6).Info("trying to resync NodeTopology", "fingerprint", pfpExpected, "onlyExclusiveResources", onlyExclRes)

		err = checkPodFingerprintForNode(lh_, objs, nodeName, pfpExpected, onlyExclRes)
		if errors.Is(err, podfingerprint.ErrSignatureMismatch) {
			// can happen, not critical
			nlh.V(5).Info("NodeTopology podset fingerprint mismatch")
			continue
		}
		if err != nil {
			// should never happen, let's be vocal
			nlh.V(3).Error(err, "nrtcache: checking NodeTopology podset fingerprint")
			continue
		}

		nlh.V(4).Info("overriding cached info")
		nrtUpdates = append(nrtUpdates, nrtCandidate)
	}

	ov.FlushNodes(lh_, nrtUpdates...)
}

// FlushNodes drops all the cached information about a given node, resetting its state clean.
func (ov *OverReserve) FlushNodes(lh logr.Logger, nrts ...*topologyv1alpha2.NodeResourceTopology) {
	ov.lock.Lock()
	defer ov.lock.Unlock()
	for _, nrt := range nrts {
		lh.V(4).Info("flushing", "node", nrt.Name)
		ov.nrts.Update(nrt)
		delete(ov.assumedResources, nrt.Name)
		ov.nodesMaybeOverreserved.Delete(nrt.Name)
		ov.nodesWithForeignPods.Delete(nrt.Name)
	}
}

func InformerFromHandle(handle framework.Handle) (k8scache.SharedIndexInformer, podlisterv1.PodLister) {
	podHandle := handle.SharedInformerFactory().Core().V1().Pods() // shortcut
	return podHandle.Informer(), podHandle.Lister()
}

// to be used only in tests
func (ov *OverReserve) Store() *nrtStore {
	return ov.nrts
}

func logIDFromTime() string {
	return fmt.Sprintf("resync%v", time.Now().UnixMilli())
}

func getCacheResyncMethod(cfg *apiconfig.NodeResourceTopologyCache) apiconfig.CacheResyncMethod {
	var resyncMethod apiconfig.CacheResyncMethod
	if cfg != nil && cfg.ResyncMethod != nil {
		resyncMethod = *cfg.ResyncMethod
	} else { // explicitly set to nil?
		resyncMethod = apiconfig.CacheResyncAutodetect
		nrtlog.Get().Info("cache resync method missing", "fallback", resyncMethod)
	}
	return resyncMethod
}

func (ov *OverReserve) PostBind(nodeName string, pod *corev1.Pod) {}
