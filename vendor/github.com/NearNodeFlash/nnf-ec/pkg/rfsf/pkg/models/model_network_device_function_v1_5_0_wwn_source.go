/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NetworkDeviceFunctionV150WWNSource string

// List of NetworkDeviceFunction_v1_5_0_WWNSource
const (
	CONFIGURED_LOCALLY_NDFV150WWNS NetworkDeviceFunctionV150WWNSource = "ConfiguredLocally"
	PROVIDED_BY_FABRIC_NDFV150WWNS NetworkDeviceFunctionV150WWNSource = "ProvidedByFabric"
)