/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// VirtualMediaV132InsertMedia - This action attaches remote media to virtual media.
type VirtualMediaV132InsertMedia struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}