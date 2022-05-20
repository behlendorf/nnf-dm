/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ChassisV1140Chassis - The Chassis schema represents the physical components of a system.  This resource represents the sheet-metal confined spaces and logical zones such as racks, enclosures, chassis and all other containers.  Subsystems, such as sensors, that operate outside of a system's data plane are linked either directly or indirectly through this resource.  A subsystem that operates outside of a system's data plane are not accessible to software that runs on the system.
type ChassisV1140Chassis struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ChassisV1140Actions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The user-assigned asset tag of this chassis.
	AssetTag string `json:"AssetTag,omitempty"`

	ChassisType ChassisV1140ChassisType `json:"ChassisType"`

	// The depth of the chassis.
	DepthMm *float32 `json:"DepthMm,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	Drives OdataV4IdRef `json:"Drives,omitempty"`

	EnvironmentalClass ChassisV1140EnvironmentalClass `json:"EnvironmentalClass,omitempty"`

	// The height of the chassis.
	HeightMm *float32 `json:"HeightMm,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	IndicatorLED ChassisV1140IndicatorLED `json:"IndicatorLED,omitempty"`

	Links ChassisV1140Links `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	LogServices OdataV4IdRef `json:"LogServices,omitempty"`

	// The manufacturer of this chassis.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The upper bound of the total power consumed by the chassis.
	MaxPowerWatts *float32 `json:"MaxPowerWatts,omitempty"`

	MediaControllers OdataV4IdRef `json:"MediaControllers,omitempty"`

	Memory OdataV4IdRef `json:"Memory,omitempty"`

	MemoryDomains OdataV4IdRef `json:"MemoryDomains,omitempty"`

	// The lower bound of the total power consumed by the chassis.
	MinPowerWatts *float32 `json:"MinPowerWatts,omitempty"`

	// The model number of the chassis.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	NetworkAdapters OdataV4IdRef `json:"NetworkAdapters,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeDevices OdataV4IdRef `json:"PCIeDevices,omitempty"`

	PCIeSlots OdataV4IdRef `json:"PCIeSlots,omitempty"`

	// The part number of the chassis.
	PartNumber string `json:"PartNumber,omitempty"`

	PhysicalSecurity ChassisV1140PhysicalSecurity `json:"PhysicalSecurity,omitempty"`

	Power OdataV4IdRef `json:"Power,omitempty"`

	PowerState ChassisV1140PowerState `json:"PowerState,omitempty"`

	// The SKU of the chassis.
	SKU string `json:"SKU,omitempty"`

	Sensors OdataV4IdRef `json:"Sensors,omitempty"`

	// The serial number of the chassis.
	SerialNumber string `json:"SerialNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	Thermal OdataV4IdRef `json:"Thermal,omitempty"`

	UUID string `json:"UUID,omitempty"`

	// The weight of the chassis.
	WeightKg *float32 `json:"WeightKg,omitempty"`

	// The width of the chassis.
	WidthMm *float32 `json:"WidthMm,omitempty"`
}