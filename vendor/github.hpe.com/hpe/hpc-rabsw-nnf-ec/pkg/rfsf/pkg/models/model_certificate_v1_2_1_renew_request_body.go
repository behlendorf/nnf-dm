/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CertificateV121RenewRequestBody - This action generates a certificate signing request by using the existing information and key-pair of the certificate.
type CertificateV121RenewRequestBody struct {

	// The challenge password to apply to the certificate for revocation requests.
	ChallengePassword string `json:"ChallengePassword,omitempty"`
}
