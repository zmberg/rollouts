/*
Copyright 2022 The Kruise Authors.

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

package workloads

import (
	"github.com/openkruise/rollouts/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WorkloadChangeEventType string

const (
	IgnoreWorkloadEvent        WorkloadChangeEventType = "workload-not-cared"
	WorkloadRollback           WorkloadChangeEventType = "workload-rollback"
	WorkloadPodTemplateChanged WorkloadChangeEventType = "workload-pod-template-changed"
	WorkloadReplicasChanged    WorkloadChangeEventType = "workload-replicas-changed"
	WorkloadStillReconciling   WorkloadChangeEventType = "workload-is-reconciling"
	WorkloadUnHealthy          WorkloadChangeEventType = "workload-is-unhealthy"
)

// WorkloadController is the interface that all type of cloneSet controller implements
type WorkloadController interface {
	// IfNeedToProgress makes sure that the resources can be upgraded according to the release plan.
	// it returns 'true' if the verification is succeeded.
	// it returns 'false' if the verification should retry.
	// it returns not-empty error if the verification has something wrong, and should not retry.
	IfNeedToProgress() (bool, error)

	// PrepareBeforeProgress make sure that the resource is ready to be progressed.
	// this function is tasked to do any initialization work on the resources.
	// it returns 'true' if the verification is succeeded.
	// it returns 'false' if the verification should retry.
	// it returns not-empty error if the verification has something wrong, and should not retry.
	PrepareBeforeProgress() (bool, error)

	// ProgressOneBatchReplicas tries to upgrade pods in the resources following the rollout plan
	// it will upgrade pods as the rollout plan allows at once
	// it returns if the upgrade actionable items succeeded/failed or should continue
	ProgressOneBatchReplicas() (bool, error)

	// CheckOneBatchReplicas checks how many pods are ready to serve requests in the current batch
	// it returns whether the number of pods upgraded in this round satisfies the rollout plan
	CheckOneBatchReplicas() (bool, error)

	// FinalizeOneBatch makes sure that the rollout can start the next batch
	// it returns if the finalization of this batch succeeded/failed or should retry
	FinalizeOneBatch() (bool, error)

	// FinalizeProgress makes sure the resources are in a good final state.
	// It might depend on if the rollout succeeded or not.
	// For example, we may remove the source object to prevent scalar traits to ever work
	// and the finalize rollout web hooks will be called after this call succeeds
	FinalizeProgress(pause *bool, cleanup bool) bool

	// SyncWorkload will watch and compare the status recorded in release.status and the real-time
	// workload status. If workload status is inconsistent with that recorded in release.status,
	// will return the corresponding WorkloadChangeEventType and info.
	SyncWorkload() (WorkloadChangeEventType, *WorkloadAccessor, error)
}

type Status struct {
	Replicas             int32
	ReadyReplicas        int32
	UpdatedReplicas      int32
	UpdatedReadyReplicas int32
	ObservedGeneration   int64
}

type WorkloadAccessor struct {
	Replicas       *int32
	Paused         bool
	Status         *Status
	UpdateRevision *string
	Metadata       *metav1.ObjectMeta
}

type workloadController struct {
	client           client.Client
	recorder         record.EventRecorder
	parentController *v1alpha1.BatchRelease

	releasePlan   *v1alpha1.ReleasePlan
	releaseStatus *v1alpha1.BatchReleaseStatus
}
