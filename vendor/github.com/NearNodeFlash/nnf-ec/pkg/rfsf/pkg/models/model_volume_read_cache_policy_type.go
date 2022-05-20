/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VolumeReadCachePolicyType string

// List of Volume_ReadCachePolicyType
const (
	READ_AHEAD_VRCPT VolumeReadCachePolicyType = "ReadAhead"
	ADAPTIVE_READ_AHEAD_VRCPT VolumeReadCachePolicyType = "AdaptiveReadAhead"
	OFF_VRCPT VolumeReadCachePolicyType = "Off"
)