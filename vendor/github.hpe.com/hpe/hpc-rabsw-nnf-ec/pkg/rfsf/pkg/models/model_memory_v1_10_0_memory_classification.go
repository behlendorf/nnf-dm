/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MemoryV1100MemoryClassification string

// List of Memory_v1_10_0_MemoryClassification
const (
	VOLATILE_MV1100MC MemoryV1100MemoryClassification = "Volatile"
	BYTE_ACCESSIBLE_PERSISTENT_MV1100MC MemoryV1100MemoryClassification = "ByteAccessiblePersistent"
	BLOCK_MV1100MC MemoryV1100MemoryClassification = "Block"
)
