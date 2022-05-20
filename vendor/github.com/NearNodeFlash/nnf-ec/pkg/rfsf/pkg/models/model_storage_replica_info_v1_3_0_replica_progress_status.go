/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// StorageReplicaInfoV130ReplicaProgressStatus : Values of ReplicaProgressStatus describe the status of the session with respect to Replication activity.
type StorageReplicaInfoV130ReplicaProgressStatus string

// List of StorageReplicaInfo_v1_3_0_ReplicaProgressStatus
const (
	COMPLETED_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Completed"
	DORMANT_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Dormant"
	INITIALIZING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Initializing"
	PREPARING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Preparing"
	SYNCHRONIZING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Synchronizing"
	RESYNCING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Resyncing"
	RESTORING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Restoring"
	FRACTURING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Fracturing"
	SPLITTING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Splitting"
	FAILING_OVER_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "FailingOver"
	FAILING_BACK_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "FailingBack"
	DETACHING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Detaching"
	ABORTING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Aborting"
	MIXED_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Mixed"
	SUSPENDING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Suspending"
	REQUIRES_FRACTURE_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "RequiresFracture"
	REQUIRES_RESYNC_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "RequiresResync"
	REQUIRES_ACTIVATE_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "RequiresActivate"
	PENDING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Pending"
	REQUIRES_DETACH_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "RequiresDetach"
	TERMINATING_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "Terminating"
	REQUIRES_SPLIT_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "RequiresSplit"
	REQUIRES_RESUME_SRIV130RPS StorageReplicaInfoV130ReplicaProgressStatus = "RequiresResume"
)