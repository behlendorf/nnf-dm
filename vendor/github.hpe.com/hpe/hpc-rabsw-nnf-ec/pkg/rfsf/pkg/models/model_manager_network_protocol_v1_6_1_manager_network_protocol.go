/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ManagerNetworkProtocolV161ManagerNetworkProtocol - The network service settings for the manager.
type ManagerNetworkProtocolV161ManagerNetworkProtocol struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ManagerNetworkProtocolV161Actions `json:"Actions,omitempty"`

	DHCP ManagerNetworkProtocolV161Protocol `json:"DHCP,omitempty"`

	DHCPv6 ManagerNetworkProtocolV161Protocol `json:"DHCPv6,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The fully qualified domain name for the manager obtained by DNS including the host name and top-level domain name.
	FQDN string `json:"FQDN,omitempty"`

	HTTP ManagerNetworkProtocolV161Protocol `json:"HTTP,omitempty"`

	HTTPS ManagerNetworkProtocolV161HttpsProtocol `json:"HTTPS,omitempty"`

	// The DNS host name of this manager, without any domain information.
	HostName string `json:"HostName,omitempty"`

	IPMI ManagerNetworkProtocolV161Protocol `json:"IPMI,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	KVMIP ManagerNetworkProtocolV161Protocol `json:"KVMIP,omitempty"`

	NTP ManagerNetworkProtocolV161NtpProtocol `json:"NTP,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	RDP ManagerNetworkProtocolV161Protocol `json:"RDP,omitempty"`

	RFB ManagerNetworkProtocolV161Protocol `json:"RFB,omitempty"`

	SNMP ManagerNetworkProtocolV161SnmpProtocol `json:"SNMP,omitempty"`

	SSDP ManagerNetworkProtocolV161SsdProtocol `json:"SSDP,omitempty"`

	SSH ManagerNetworkProtocolV161Protocol `json:"SSH,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	Telnet ManagerNetworkProtocolV161Protocol `json:"Telnet,omitempty"`

	VirtualMedia ManagerNetworkProtocolV161Protocol `json:"VirtualMedia,omitempty"`
}