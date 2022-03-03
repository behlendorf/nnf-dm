/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CircuitV110EnergySensors - The energy readings for this circuit.
type CircuitV110EnergySensors struct {

	Line1ToLine2 SensorSensorEnergykWhExcerpt `json:"Line1ToLine2,omitempty"`

	Line1ToNeutral SensorSensorEnergykWhExcerpt `json:"Line1ToNeutral,omitempty"`

	Line2ToLine3 SensorSensorEnergykWhExcerpt `json:"Line2ToLine3,omitempty"`

	Line2ToNeutral SensorSensorEnergykWhExcerpt `json:"Line2ToNeutral,omitempty"`

	Line3ToLine1 SensorSensorEnergykWhExcerpt `json:"Line3ToLine1,omitempty"`

	Line3ToNeutral SensorSensorEnergykWhExcerpt `json:"Line3ToNeutral,omitempty"`
}