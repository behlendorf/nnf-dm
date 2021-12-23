/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MemoryV1100SecurityCapabilities - This type contains security capabilities of a memory device.
type MemoryV1100SecurityCapabilities struct {

	// An indication of whether this memory device supports the locking, or freezing, of the configuration.
	ConfigurationLockCapable bool `json:"ConfigurationLockCapable,omitempty"`

	// An indication of whether this memory device supports data locking.
	DataLockCapable bool `json:"DataLockCapable,omitempty"`

	// Maximum number of passphrases supported for this memory device.
	MaxPassphraseCount int64 `json:"MaxPassphraseCount,omitempty"`

	// An indication of whether the memory device is passphrase capable.
	PassphraseCapable bool `json:"PassphraseCapable,omitempty"`

	// The maximum number of incorrect passphrase attempts allowed before memory device is locked.
	PassphraseLockLimit int64 `json:"PassphraseLockLimit,omitempty"`

	// Security states supported by the memory device.
	SecurityStates []MemoryV1100SecurityStates `json:"SecurityStates,omitempty"`
}
