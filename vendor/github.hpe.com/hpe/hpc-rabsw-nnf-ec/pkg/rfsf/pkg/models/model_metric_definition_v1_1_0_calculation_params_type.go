/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MetricDefinitionV110CalculationParamsType - The usage of the parameter in the calculation.
type MetricDefinitionV110CalculationParamsType struct {

	// The link to a metric property that stores the result of the calculation.  If the link has wildcards, the wildcards are substituted as specified in the Wildcards array property.
	ResultMetric string `json:"ResultMetric,omitempty"`

	// The metric property used as the input into the calculation.  If the link has wildcards, the wildcards are substituted as specified in the Wildcards array property.
	SourceMetric string `json:"SourceMetric,omitempty"`
}
