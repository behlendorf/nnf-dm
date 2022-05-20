/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StorageServiceV150StorageService - Collection of resources that are managed and exposed to hosts as a group.
type StorageServiceV150StorageService struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions StorageServiceV150Actions `json:"Actions,omitempty"`

	ClassesOfService OdataV4IdRef `json:"ClassesOfService,omitempty"`

	ClientEndpointGroups OdataV4IdRef `json:"ClientEndpointGroups,omitempty"`

	ConsistencyGroups OdataV4IdRef `json:"ConsistencyGroups,omitempty"`

	DataProtectionLoSCapabilities OdataV4IdRef `json:"DataProtectionLoSCapabilities,omitempty"`

	DataSecurityLoSCapabilities OdataV4IdRef `json:"DataSecurityLoSCapabilities,omitempty"`

	DataStorageLoSCapabilities OdataV4IdRef `json:"DataStorageLoSCapabilities,omitempty"`

	DefaultClassOfService OdataV4IdRef `json:"DefaultClassOfService,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	Drives OdataV4IdRef `json:"Drives,omitempty"`

	EndpointGroups OdataV4IdRef `json:"EndpointGroups,omitempty"`

	Endpoints OdataV4IdRef `json:"Endpoints,omitempty"`

	FileSystems OdataV4IdRef `json:"FileSystems,omitempty"`

	IOConnectivityLoSCapabilities OdataV4IdRef `json:"IOConnectivityLoSCapabilities,omitempty"`

	IOPerformanceLoSCapabilities OdataV4IdRef `json:"IOPerformanceLoSCapabilities,omitempty"`

	IOStatistics IoStatisticsIoStatistics `json:"IOStatistics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	Identifier ResourceIdentifier `json:"Identifier,omitempty"`

	// The LinesOService defined for this StorageService.
	LinesOfService []OdataV4IdRef `json:"LinesOfService,omitempty"`

	// The number of items in a collection.
	LinesOfServiceodataCount int64 `json:"LinesOfService@odata.count,omitempty"`

	Links StorageServiceV150Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Redundancy information for the storage subsystem.
	Redundancy []RedundancyRedundancy `json:"Redundancy,omitempty"`

	// The number of items in a collection.
	RedundancyodataCount int64 `json:"Redundancy@odata.count,omitempty"`

	ServerEndpointGroups OdataV4IdRef `json:"ServerEndpointGroups,omitempty"`

	// An array of SpareResourceSets.
	SpareResourceSets []OdataV4IdRef `json:"SpareResourceSets,omitempty"`

	// The number of items in a collection.
	SpareResourceSetsodataCount int64 `json:"SpareResourceSets@odata.count,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	StorageGroups OdataV4IdRef `json:"StorageGroups,omitempty"`

	StoragePools OdataV4IdRef `json:"StoragePools,omitempty"`

	StorageSubsystems OdataV4IdRef `json:"StorageSubsystems,omitempty"`

	Volumes OdataV4IdRef `json:"Volumes,omitempty"`
}