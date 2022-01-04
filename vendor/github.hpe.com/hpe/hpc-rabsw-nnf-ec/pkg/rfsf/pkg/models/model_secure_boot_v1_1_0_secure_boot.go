/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SecureBootV110SecureBoot - The SecureBoot schema contains UEFI Secure Boot information and represents properties for managing the UEFI Secure Boot functionality of a system.
type SecureBootV110SecureBoot struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions SecureBootV110Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	SecureBootCurrentBoot SecureBootV110SecureBootCurrentBootType `json:"SecureBootCurrentBoot,omitempty"`

	SecureBootDatabases OdataV4IdRef `json:"SecureBootDatabases,omitempty"`

	// An indication of whether UEFI Secure Boot is enabled.
	SecureBootEnable bool `json:"SecureBootEnable,omitempty"`

	SecureBootMode SecureBootV110SecureBootModeType `json:"SecureBootMode,omitempty"`
}
