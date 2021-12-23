/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StorageGroupV150StorageGroup - Collection of resources that are managed and exposed to hosts as a group.
type StorageGroupV150StorageGroup struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	AccessState EndpointGroupAccessState `json:"AccessState,omitempty"`

	Actions StorageGroupV150Actions `json:"Actions,omitempty"`

	AuthenticationMethod StorageGroupV150AuthenticationMethod `json:"AuthenticationMethod,omitempty"`

	// The credential information used to authenticate the endpoints in this StorageGroup.
	ChapInfo []StorageGroupV150ChapInformation `json:"ChapInfo,omitempty"`

	// Groups of client endpoints in this storage group.
	ClientEndpointGroups []OdataV4IdRef `json:"ClientEndpointGroups,omitempty"`

	// The number of items in a collection.
	ClientEndpointGroupsodataCount int64 `json:"ClientEndpointGroups@odata.count,omitempty"`

	// The credential information used to authenticate the endpoints in this StorageGroup for DHCHAP.
	DHChapInfo []StorageGroupV150DhchapInformation `json:"DHChapInfo,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	Identifier ResourceIdentifier `json:"Identifier,omitempty"`

	Links StorageGroupV150Links `json:"Links,omitempty"`

	// Mapped Volumes in this storage group.
	MappedVolumes []StorageGroupMappedVolume `json:"MappedVolumes,omitempty"`

	// Members are kept in a consistent state.
	MembersAreConsistent bool `json:"MembersAreConsistent,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	ReplicaInfo StorageReplicaInfoReplicaInfo `json:"ReplicaInfo,omitempty"`

	// The resources that are target replicas of this source.
	ReplicaTargets []OdataV4IdRef `json:"ReplicaTargets,omitempty"`

	// The number of items in a collection.
	ReplicaTargetsodataCount int64 `json:"ReplicaTargets@odata.count,omitempty"`

	// Server endpoints in this storage group.
	ServerEndpoints []OdataV4IdRef `json:"ServerEndpoints,omitempty"`

	// The number of items in a collection.
	ServerEndpointsodataCount int64 `json:"ServerEndpoints@odata.count,omitempty"`

	// Groups of server endpoints in this storage group.
	ServerEndpointGroups []OdataV4IdRef `json:"ServerEndpointGroups,omitempty"`

	// The number of items in a collection.
	ServerEndpointGroupsodataCount int64 `json:"ServerEndpointGroups@odata.count,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// Volumes in this storage group.
	Volumes []OdataV4IdRef `json:"Volumes,omitempty"`

	// The number of items in a collection.
	VolumesodataCount int64 `json:"Volumes@odata.count,omitempty"`

	// Storage volumes are exposed to paths defined by the client and server endpoints.
	VolumesAreExposed bool `json:"VolumesAreExposed,omitempty"`
}
