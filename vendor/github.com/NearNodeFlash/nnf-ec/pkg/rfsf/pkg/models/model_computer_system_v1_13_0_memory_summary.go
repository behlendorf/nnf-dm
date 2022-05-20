/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ComputerSystemV1130MemorySummary - The memory of the system in general detail.
type ComputerSystemV1130MemorySummary struct {

	MemoryMirroring ComputerSystemV1130MemoryMirroring `json:"MemoryMirroring,omitempty"`

	Metrics OdataV4IdRef `json:"Metrics,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The total configured operating system-accessible memory (RAM), measured in GiB.
	TotalSystemMemoryGiB *float32 `json:"TotalSystemMemoryGiB,omitempty"`

	// The total configured, system-accessible persistent memory, measured in GiB.
	TotalSystemPersistentMemoryGiB *float32 `json:"TotalSystemPersistentMemoryGiB,omitempty"`
}