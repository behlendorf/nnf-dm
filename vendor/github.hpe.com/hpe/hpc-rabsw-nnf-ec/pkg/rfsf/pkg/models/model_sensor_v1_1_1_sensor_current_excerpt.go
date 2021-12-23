/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SensorV111SensorCurrentExcerpt - The Sensor schema describes a sensor and its properties.
type SensorV111SensorCurrentExcerpt struct {

	// The crest factor for this sensor.
	CrestFactor *float32 `json:"CrestFactor,omitempty"`

	// The link to the resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The sensor value.
	Reading *float32 `json:"Reading,omitempty"`

	// The total harmonic distortion (THD).
	THDPercent *float32 `json:"THDPercent,omitempty"`
}
