/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ComputerSystemV1130SerialConsoleProtocol - The information about a serial console service that this system provides.
type ComputerSystemV1130SerialConsoleProtocol struct {

	// The command string passed to the service to select or enter the system's serial console.
	ConsoleEntryCommand string `json:"ConsoleEntryCommand,omitempty"`

	// The hotkey sequence available for the user to exit the serial console session.
	HotKeySequenceDisplay string `json:"HotKeySequenceDisplay,omitempty"`

	// The protocol port.
	Port int64 `json:"Port,omitempty"`

	// An indication of whether the service is enabled for this system.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// Indicates whether the serial console service is shared with access to the manager's command-line interface (CLI).
	SharedWithManagerCLI bool `json:"SharedWithManagerCLI,omitempty"`
}