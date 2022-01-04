/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ManagerV1100CommandConnectTypesSupported string

// List of Manager_v1_10_0_CommandConnectTypesSupported
const (
	SSH_MV1100CCTS ManagerV1100CommandConnectTypesSupported = "SSH"
	TELNET_MV1100CCTS ManagerV1100CommandConnectTypesSupported = "Telnet"
	IPMI_MV1100CCTS ManagerV1100CommandConnectTypesSupported = "IPMI"
	OEM_MV1100CCTS ManagerV1100CommandConnectTypesSupported = "Oem"
)
