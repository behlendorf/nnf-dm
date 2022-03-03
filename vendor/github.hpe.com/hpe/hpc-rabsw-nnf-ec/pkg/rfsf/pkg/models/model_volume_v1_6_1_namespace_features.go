/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VolumeV161NamespaceFeatures struct {

	// Indicates that the NVM fields for Namespace preferred write granularity (NPWG), write alignment (NPWA), deallocate granularity (NPDG), deallocate alignment (NPDA) and optimimal write size (NOWS)  are defined for this namespace and should be used by the host for I/O optimization.
	SupportsAtomicTransactionSize bool `json:"SupportsAtomicTransactionSize,omitempty"`

	// This property indicates that the controller supports deallocated or unwritten logical block error for this namespace.
	SupportsDeallocatedOrUnwrittenLBError bool `json:"SupportsDeallocatedOrUnwrittenLBError,omitempty"`

	// Indicates that the Namespace Atomic Write Unit Normal (NAWUN), Namespace Atomic Write Unit Power Fail (NAWUPF), and Namespace Atomic Compare and Write Unit (NACWU) fields are defined for this namespace and should be used by the host for this namespace instead of the controller-level properties AWUN, AWUPF, and ACWU.
	SupportsIOPerformanceHints bool `json:"SupportsIOPerformanceHints,omitempty"`

	// This property indicates that the namespace supports the use of an NGUID (namespace globally unique identifier) value.
	SupportsNGUIDReuse bool `json:"SupportsNGUIDReuse,omitempty"`

	// This property indicates whether or not the NVMe Namespace supports thin provisioning.
	SupportsThinProvisioning bool `json:"SupportsThinProvisioning,omitempty"`
}