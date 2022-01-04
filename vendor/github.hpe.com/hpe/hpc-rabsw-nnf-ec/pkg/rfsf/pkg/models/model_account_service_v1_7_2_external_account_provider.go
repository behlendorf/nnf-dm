/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AccountServiceV172ExternalAccountProvider - The external account provider services that can provide accounts for this manager to use for authentication.
type AccountServiceV172ExternalAccountProvider struct {

	AccountProviderType AccountServiceV172AccountProviderTypes `json:"AccountProviderType,omitempty"`

	Authentication AccountServiceV172Authentication `json:"Authentication,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	LDAPService AccountServiceV172LdapService `json:"LDAPService,omitempty"`

	// Indicates if the Password property is set.
	PasswordSet bool `json:"PasswordSet,omitempty"`

	// The mapping rules to convert the external account providers account information to the local Redfish role.
	RemoteRoleMapping []AccountServiceV172RoleMapping `json:"RemoteRoleMapping,omitempty"`

	// The addresses of the user account providers to which this external account provider links.  The format of this field depends on the type of external account provider.
	ServiceAddresses []string `json:"ServiceAddresses,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
}
