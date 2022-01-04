/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NetworkAdapterV150Npiv - N_Port ID Virtualization (NPIV) capabilities for a controller.
type NetworkAdapterV150Npiv struct {

	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxDeviceLogins int64 `json:"MaxDeviceLogins,omitempty"`

	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical port on this controller.
	MaxPortLogins int64 `json:"MaxPortLogins,omitempty"`
}
