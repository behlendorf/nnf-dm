/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PortV130ConfiguredNetworkLink - A set of link settings that a port is configured to use for autonegotiation.
type PortV130ConfiguredNetworkLink struct {

	// The link speed per lane this port is configured to use for autonegotiation.
	ConfiguredLinkSpeedGbps *float32 `json:"ConfiguredLinkSpeedGbps,omitempty"`

	// The link width this port is configured to use for autonegotiation in conjunction with the link speed.
	ConfiguredWidth int64 `json:"ConfiguredWidth,omitempty"`
}
