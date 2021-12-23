/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EventDestinationV190SNMPAuthenticationProtocols string

// List of EventDestination_v1_9_0_SNMPAuthenticationProtocols
const (
	NONE_EDV190SNMPAP EventDestinationV190SNMPAuthenticationProtocols = "None"
	COMMUNITY_STRING_EDV190SNMPAP EventDestinationV190SNMPAuthenticationProtocols = "CommunityString"
	HMAC_MD5_EDV190SNMPAP EventDestinationV190SNMPAuthenticationProtocols = "HMAC_MD5"
	HMAC_SHA96_EDV190SNMPAP EventDestinationV190SNMPAuthenticationProtocols = "HMAC_SHA96"
)
