/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// VLanNetworkInterfaceVlan - The attributes of a VLAN.
type VLanNetworkInterfaceVlan struct {

	// An indication of whether this VLAN is enabled for this VLAN network interface.
	VLANEnable bool `json:"VLANEnable,omitempty"`

	VLANId int64 `json:"VLANId,omitempty"`
}