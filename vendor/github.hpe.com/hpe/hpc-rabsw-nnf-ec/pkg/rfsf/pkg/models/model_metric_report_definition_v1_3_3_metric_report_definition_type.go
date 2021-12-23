/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MetricReportDefinitionV133MetricReportDefinitionType : Indicates when the metric report is generated.
type MetricReportDefinitionV133MetricReportDefinitionType string

// List of MetricReportDefinition_v1_3_3_MetricReportDefinitionType
const (
	PERIODIC_MRDV133MRDT MetricReportDefinitionV133MetricReportDefinitionType = "Periodic"
	ON_CHANGE_MRDV133MRDT MetricReportDefinitionV133MetricReportDefinitionType = "OnChange"
	ON_REQUEST_MRDV133MRDT MetricReportDefinitionV133MetricReportDefinitionType = "OnRequest"
)
