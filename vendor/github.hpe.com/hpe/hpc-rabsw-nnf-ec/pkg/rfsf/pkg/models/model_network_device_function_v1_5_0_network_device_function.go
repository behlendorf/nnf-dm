/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NetworkDeviceFunctionV150NetworkDeviceFunction - The NetworkDeviceFunction schema represents a logical interface that a network adapter exposes.
type NetworkDeviceFunctionV150NetworkDeviceFunction struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions NetworkDeviceFunctionV150Actions `json:"Actions,omitempty"`

	// An array of physical ports to which this network device function can be assigned.
	AssignablePhysicalNetworkPorts []OdataV4IdRef `json:"AssignablePhysicalNetworkPorts,omitempty"`

	// The number of items in a collection.
	AssignablePhysicalNetworkPortsodataCount int64 `json:"AssignablePhysicalNetworkPorts@odata.count,omitempty"`

	// An array of physical ports to which this network device function can be assigned.
	AssignablePhysicalPorts []OdataV4IdRef `json:"AssignablePhysicalPorts,omitempty"`

	// The number of items in a collection.
	AssignablePhysicalPortsodataCount int64 `json:"AssignablePhysicalPorts@odata.count,omitempty"`

	BootMode NetworkDeviceFunctionV150BootMode `json:"BootMode,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether the network device function is enabled.
	DeviceEnabled bool `json:"DeviceEnabled,omitempty"`

	Ethernet NetworkDeviceFunctionV150Ethernet `json:"Ethernet,omitempty"`

	FibreChannel NetworkDeviceFunctionV150FibreChannel `json:"FibreChannel,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	InfiniBand NetworkDeviceFunctionV150InfiniBand `json:"InfiniBand,omitempty"`

	Links NetworkDeviceFunctionV150Links `json:"Links,omitempty"`

	// The number of virtual functions that are available for this network device function.
	MaxVirtualFunctions int64 `json:"MaxVirtualFunctions,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// An array of capabilities for this network device function.
	NetDevFuncCapabilities []NetworkDeviceFunctionV150NetworkDeviceTechnology `json:"NetDevFuncCapabilities,omitempty"`

	NetDevFuncType NetworkDeviceFunctionV150NetworkDeviceTechnology `json:"NetDevFuncType,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PhysicalNetworkPortAssignment OdataV4IdRef `json:"PhysicalNetworkPortAssignment,omitempty"`

	PhysicalPortAssignment OdataV4IdRef `json:"PhysicalPortAssignment,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// An indication of whether single root input/output virtualization (SR-IOV) virtual functions are enabled for this network device function.
	VirtualFunctionsEnabled bool `json:"VirtualFunctionsEnabled,omitempty"`

	ISCSIBoot NetworkDeviceFunctionV150IScsiBoot `json:"iSCSIBoot,omitempty"`
}
