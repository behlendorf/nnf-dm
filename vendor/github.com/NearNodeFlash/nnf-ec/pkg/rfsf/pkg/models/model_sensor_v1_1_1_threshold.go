/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SensorV111Threshold - The threshold definition for a sensor.
type SensorV111Threshold struct {

	Activation SensorV111ThresholdActivation `json:"Activation,omitempty"`

	// The duration the sensor value must violate the threshold before the threshold is activated.
	DwellTime string `json:"DwellTime,omitempty"`

	// The threshold value.
	Reading *float32 `json:"Reading,omitempty"`
}