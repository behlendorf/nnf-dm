/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectionV100Links - The links to other resources that are related to this resource.
type ConnectionV100Links struct {

	// An array of links to the initiator endpoint groups that are associated with this connection.
	InitiatorEndpointGroups []OdataV4IdRef `json:"InitiatorEndpointGroups,omitempty"`

	// The number of items in a collection.
	InitiatorEndpointGroupsodataCount int64 `json:"InitiatorEndpointGroups@odata.count,omitempty"`

	// An array of links to the initiator endpoints that are associated with this connection.
	InitiatorEndpoints []OdataV4IdRef `json:"InitiatorEndpoints,omitempty"`

	// The number of items in a collection.
	InitiatorEndpointsodataCount int64 `json:"InitiatorEndpoints@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the target endpoint groups that are associated with this connection.
	TargetEndpointGroups []OdataV4IdRef `json:"TargetEndpointGroups,omitempty"`

	// The number of items in a collection.
	TargetEndpointGroupsodataCount int64 `json:"TargetEndpointGroups@odata.count,omitempty"`

	// An array of links to the target endpoints that are associated with this connection.
	TargetEndpoints []OdataV4IdRef `json:"TargetEndpoints,omitempty"`

	// The number of items in a collection.
	TargetEndpointsodataCount int64 `json:"TargetEndpoints@odata.count,omitempty"`
}