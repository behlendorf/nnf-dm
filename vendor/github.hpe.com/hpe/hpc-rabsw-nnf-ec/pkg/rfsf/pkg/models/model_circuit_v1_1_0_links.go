/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CircuitV110Links - The links to other resources that are related to this resource.
type CircuitV110Links struct {

	BranchCircuit OdataV4IdRef `json:"BranchCircuit,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of references to the outlets contained by this circuit.
	Outlets []OdataV4IdRef `json:"Outlets,omitempty"`

	// The number of items in a collection.
	OutletsodataCount int64 `json:"Outlets@odata.count,omitempty"`
}
