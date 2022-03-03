/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ResourceV166Orientation : The orientation for the ordering of the part location ordinal value.
type ResourceV166Orientation string

// List of Resource_v1_6_6_Orientation
const (
	FRONT_TO_BACK_RV166O ResourceV166Orientation = "FrontToBack"
	BACK_TO_FRONT_RV166O ResourceV166Orientation = "BackToFront"
	TOP_TO_BOTTOM_RV166O ResourceV166Orientation = "TopToBottom"
	BOTTOM_TO_TOP_RV166O ResourceV166Orientation = "BottomToTop"
	LEFT_TO_RIGHT_RV166O ResourceV166Orientation = "LeftToRight"
	RIGHT_TO_LEFT_RV166O ResourceV166Orientation = "RightToLeft"
)