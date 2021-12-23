/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// IoPerformanceLineOfServiceV111IoPerformanceLineOfService - Describe service option within the IO performance line of service.
type IoPerformanceLineOfServiceV111IoPerformanceLineOfService struct {

	Actions IoPerformanceLineOfServiceV111Actions `json:"Actions,omitempty"`

	// Expected average IO latency.
	AverageIOOperationLatencyMicroseconds int64 `json:"AverageIOOperationLatencyMicroseconds,omitempty"`

	// Limit the IOPS.
	IOOperationsPerSecondIsLimited bool `json:"IOOperationsPerSecondIsLimited,omitempty"`

	IOWorkload IoPerformanceLoSCapabilitiesV100IoWorkload `json:"IOWorkload,omitempty"`

	// The amount of IOPS a volume of a given committed size can support.
	MaxIOOperationsPerSecondPerTerabyte int64 `json:"MaxIOOperationsPerSecondPerTerabyte,omitempty"`

	// Sampling period over which average values are calculated.
	SamplePeriod string `json:"SamplePeriod,omitempty"`
}
