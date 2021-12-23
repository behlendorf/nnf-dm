/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ExternalAccountProviderV113LdapSearchSettings - The settings to search a generic LDAP service.
type ExternalAccountProviderV113LdapSearchSettings struct {

	// The base distinguished names to use to search an external LDAP service.
	BaseDistinguishedNames []string `json:"BaseDistinguishedNames,omitempty"`

	// The attribute name that contains the LDAP group name entry.
	GroupNameAttribute string `json:"GroupNameAttribute,omitempty"`

	// The attribute name that contains the groups for a user on the LDAP user entry.
	GroupsAttribute string `json:"GroupsAttribute,omitempty"`

	// The attribute name that contains the LDAP user name entry.
	UsernameAttribute string `json:"UsernameAttribute,omitempty"`
}
