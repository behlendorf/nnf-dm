/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// FileSystemFileProtocol : The file sharing protocols supported by the file system.
type FileSystemFileProtocol string

// List of FileSystem_FileProtocol
const (
	NF_SV3_FSFP FileSystemFileProtocol = "NFSv3"
	NF_SV4_0_FSFP FileSystemFileProtocol = "NFSv4_0"
	NF_SV4_1_FSFP FileSystemFileProtocol = "NFSv4_1"
	SM_BV2_0_FSFP FileSystemFileProtocol = "SMBv2_0"
	SM_BV2_1_FSFP FileSystemFileProtocol = "SMBv2_1"
	SM_BV3_0_FSFP FileSystemFileProtocol = "SMBv3_0"
	SM_BV3_0_2_FSFP FileSystemFileProtocol = "SMBv3_0_2"
	SM_BV3_1_1_FSFP FileSystemFileProtocol = "SMBv3_1_1"
)