/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DataProtectionLoSCapabilitiesRecoveryAccessScope : An enumeration that represents the relative time required to make a replica available as a source.
type DataProtectionLoSCapabilitiesRecoveryAccessScope string

// List of DataProtectionLoSCapabilities_RecoveryAccessScope
const (
	ONLINE_ACTIVE_DPLSCRAS DataProtectionLoSCapabilitiesRecoveryAccessScope = "OnlineActive"
	ONLINE_PASSIVE_DPLSCRAS DataProtectionLoSCapabilitiesRecoveryAccessScope = "OnlinePassive"
	NEARLINE_DPLSCRAS DataProtectionLoSCapabilitiesRecoveryAccessScope = "Nearline"
	OFFLINE_DPLSCRAS DataProtectionLoSCapabilitiesRecoveryAccessScope = "Offline"
)
