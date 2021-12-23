/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SensorV111Actions - The available actions for this resource.
type SensorV111Actions struct {

	SensorResetMetrics SensorV111ResetMetrics `json:"#Sensor.ResetMetrics,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}
