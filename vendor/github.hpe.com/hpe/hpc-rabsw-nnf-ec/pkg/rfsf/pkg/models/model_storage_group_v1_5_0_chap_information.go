/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StorageGroupV150ChapInformation - User name and password for CHAP authentication.
type StorageGroupV150ChapInformation struct {

	// The password for CHAP authentication.
	CHAPPassword string `json:"CHAPPassword,omitempty"`

	// The username for CHAP authentication.
	CHAPUser string `json:"CHAPUser,omitempty"`

	// The shared secret for Mutual (2-way) CHAP authentication by the initiator.
	InitiatorCHAPPassword string `json:"InitiatorCHAPPassword,omitempty"`

	// The Initiator username for Mutual (2-way) CHAP authentication by the initiator.
	InitiatorCHAPUser string `json:"InitiatorCHAPUser,omitempty"`

	// The Target CHAP Secret for Mutual (2-way) CHAP authentication by the target.
	TargetCHAPPassword string `json:"TargetCHAPPassword,omitempty"`

	// The Target CHAP Username for Mutual (2-way) CHAP authentication by the target.
	TargetCHAPUser string `json:"TargetCHAPUser,omitempty"`

	// This property is deprecated in favor of TargetCHAPPassword.
	TargetPassword string `json:"TargetPassword,omitempty"`
}
