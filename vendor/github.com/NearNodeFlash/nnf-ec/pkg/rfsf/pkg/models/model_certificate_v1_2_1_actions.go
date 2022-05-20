/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CertificateV121Actions - The available actions for this resource.
type CertificateV121Actions struct {

	CertificateRekey CertificateV121Rekey `json:"#Certificate.Rekey,omitempty"`

	CertificateRenew CertificateV121Renew `json:"#Certificate.Renew,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}