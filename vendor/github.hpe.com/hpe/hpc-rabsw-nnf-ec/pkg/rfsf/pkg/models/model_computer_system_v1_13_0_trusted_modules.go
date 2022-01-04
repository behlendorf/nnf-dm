/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ComputerSystemV1130TrustedModules - The Trusted Module installed in the system.
type ComputerSystemV1130TrustedModules struct {

	// The firmware version of this Trusted Module.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The second firmware version of this Trusted Module, if applicable.
	FirmwareVersion2 string `json:"FirmwareVersion2,omitempty"`

	InterfaceType ComputerSystemV1130InterfaceType `json:"InterfaceType,omitempty"`

	InterfaceTypeSelection ComputerSystemV1130InterfaceTypeSelection `json:"InterfaceTypeSelection,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`
}
