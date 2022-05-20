/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AggregationSourceV100AggregationSource - The AggregationSource schema is used to represent the source of information for a subset of the resources provided by a Redfish service.  It can be thought of as a provider of information.  As such, most such interfaces have requirements to support the gathering of information like address and account used to access the information.
type AggregationSourceV100AggregationSource struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions AggregationSourceV100Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The URI of the system to be accessed.
	HostName string `json:"HostName"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	Links AggregationSourceV100Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The password for accessing the aggregation source.  The value is `null` in responses.
	Password string `json:"Password,omitempty"`

	// The user name for accessing the aggregation source.
	UserName string `json:"UserName,omitempty"`
}