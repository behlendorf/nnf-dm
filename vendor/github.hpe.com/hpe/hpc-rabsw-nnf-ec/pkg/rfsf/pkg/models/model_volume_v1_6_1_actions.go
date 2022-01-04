/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VolumeV161Actions struct {

	VolumeAssignReplicaTarget VolumeV161AssignReplicaTarget `json:"#Volume.AssignReplicaTarget,omitempty"`

	VolumeChangeRAIDLayout VolumeV161ChangeRaidLayout `json:"#Volume.ChangeRAIDLayout,omitempty"`

	VolumeCheckConsistency VolumeV161CheckConsistency `json:"#Volume.CheckConsistency,omitempty"`

	VolumeCreateReplicaTarget VolumeV161CreateReplicaTarget `json:"#Volume.CreateReplicaTarget,omitempty"`

	VolumeForceEnable VolumeV161ForceEnable `json:"#Volume.ForceEnable,omitempty"`

	VolumeInitialize VolumeV161Initialize `json:"#Volume.Initialize,omitempty"`

	VolumeRemoveReplicaRelationship VolumeV161RemoveReplicaRelationship `json:"#Volume.RemoveReplicaRelationship,omitempty"`

	VolumeResumeReplication VolumeV161ResumeReplication `json:"#Volume.ResumeReplication,omitempty"`

	VolumeReverseReplicationRelationship VolumeV161ReverseReplicationRelationship `json:"#Volume.ReverseReplicationRelationship,omitempty"`

	VolumeSplitReplication VolumeV161SplitReplication `json:"#Volume.SplitReplication,omitempty"`

	VolumeSuspendReplication VolumeV161SuspendReplication `json:"#Volume.SuspendReplication,omitempty"`

	Oem map[string]interface{} `json:"Oem,omitempty"`
}
