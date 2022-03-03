/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AggregationServiceV100AggregationService - The AggregationService schema contains properties for managing aggregation operations, either on ad hoc combinations of resources or on defined sets of resources called aggregates.  Access points define the properties needed to access the entity being aggregated and connection methods describe the protocol or other semantics of the connection.
type AggregationServiceV100AggregationService struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions AggregationServiceV100Actions `json:"Actions,omitempty"`

	Aggregates OdataV4IdRef `json:"Aggregates,omitempty"`

	AggregationSources OdataV4IdRef `json:"AggregationSources,omitempty"`

	ConnectionMethods OdataV4IdRef `json:"ConnectionMethods,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether the aggregation service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`
}