/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MetricReportDefinitionV133MetricReportDefinition - The MetricReportDefinition schema describes set of metrics that are collected into a metric report.
type MetricReportDefinitionV133MetricReportDefinition struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions MetricReportDefinitionV133Actions `json:"Actions,omitempty"`

	// The maximum number of entries that can be appended to a metric report.  When the metric report reaches its limit, its behavior is dictated by the ReportUpdates property.
	AppendLimit int64 `json:"AppendLimit,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	Links MetricReportDefinitionV133Links `json:"Links,omitempty"`

	// The list of URIs with wildcards and property identifiers to include in the metric report.  If a URI has wildcards, the wildcards are substituted as specified in the Wildcards property.
	MetricProperties []string `json:"MetricProperties,omitempty"`

	MetricReport OdataV4IdRef `json:"MetricReport,omitempty"`

	// An indication of whether the generation of new metric reports is enabled.
	MetricReportDefinitionEnabled bool `json:"MetricReportDefinitionEnabled,omitempty"`

	MetricReportDefinitionType MetricReportDefinitionV133MetricReportDefinitionType `json:"MetricReportDefinitionType,omitempty"`

	// The interval at which to send the complete metric report because the Redfish client wants refreshed metric data even when the data has not changed.  This property value is always greater than the recurrence interval of a metric report, and it only applies when the SuppressRepeatedMetricValue property is `true`.
	MetricReportHeartbeatInterval string `json:"MetricReportHeartbeatInterval,omitempty"`

	// The list of metrics to include in the metric report.  The metrics might include metric properties or calculations applied to a metric property.
	Metrics []MetricReportDefinitionV133Metric `json:"Metrics,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The set of actions to perform when a metric report is generated.
	ReportActions []MetricReportDefinitionV133ReportActionsEnum `json:"ReportActions,omitempty"`

	// The maximum timespan that a metric report can cover.
	ReportTimespan string `json:"ReportTimespan,omitempty"`

	ReportUpdates MetricReportDefinitionV133ReportUpdatesEnum `json:"ReportUpdates,omitempty"`

	Schedule ScheduleSchedule `json:"Schedule,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// An indication of whether any metrics are suppressed from the generated metric report.  If `true`, any metric that equals the same value in the previously generated metric report is suppressed from the current report.  Also, duplicate metrics are suppressed.  If `false`, no metrics are suppressed from the current report.  The current report might contain no metrics if all metrics equal the values in the previously generated metric report.
	SuppressRepeatedMetricValue bool `json:"SuppressRepeatedMetricValue,omitempty"`

	// The set of wildcards and their substitution values for the entries in the MetricProperties property.
	Wildcards []MetricReportDefinitionV133Wildcard `json:"Wildcards,omitempty"`
}
