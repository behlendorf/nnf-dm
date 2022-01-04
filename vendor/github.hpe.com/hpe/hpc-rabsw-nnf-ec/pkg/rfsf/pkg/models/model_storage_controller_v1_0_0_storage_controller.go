/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StorageControllerV100StorageController - The StorageController schema describes a storage controller and its properties.  A storage controller represents a physical or virtual storage device that produces volumes.
type StorageControllerV100StorageController struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions StorageControllerV100Actions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this storage controller.
	AssetTag string `json:"AssetTag,omitempty"`

	CacheSummary StorageControllerV100CacheSummary `json:"CacheSummary,omitempty"`

	ControllerRates StorageControllerV100Rates `json:"ControllerRates,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The firmware version of this storage controller.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable names for the storage controller.
	Identifiers []ResourceIdentifier `json:"Identifiers,omitempty"`

	Links StorageControllerV100Links `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// The manufacturer of this storage controller.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model number for the storage controller.
	Model string `json:"Model,omitempty"`

	NVMeControllerProperties StorageControllerV100NvMeControllerProperties `json:"NVMeControllerProperties,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeInterface PcIeDevicePcIeInterface `json:"PCIeInterface,omitempty"`

	// The part number for this storage controller.
	PartNumber string `json:"PartNumber,omitempty"`

	Ports OdataV4IdRef `json:"Ports,omitempty"`

	// The SKU for this storage controller.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this storage controller.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The maximum speed of the storage controller's device interface.
	SpeedGbps *float32 `json:"SpeedGbps,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The supported set of protocols for communicating to this storage controller.
	SupportedControllerProtocols []ProtocolProtocol `json:"SupportedControllerProtocols,omitempty"`

	// The protocols that the storage controller can use to communicate with attached devices.
	SupportedDeviceProtocols []ProtocolProtocol `json:"SupportedDeviceProtocols,omitempty"`

	// The set of RAID types supported by the storage controller.
	SupportedRAIDTypes []VolumeRAIDType `json:"SupportedRAIDTypes,omitempty"`
}
