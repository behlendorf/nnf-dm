/*
 * Swordfish API
 *
 * This contains the definition of the Swordfish extensions to a Redfish service.
 *
 * API version: v1.2.c
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EventServiceV170SMTPAuthenticationMethods string

// List of EventService_v1_7_0_SMTPAuthenticationMethods
const (
	NONE_ESV170SMTPAM EventServiceV170SMTPAuthenticationMethods = "None"
	AUTO_DETECT_ESV170SMTPAM EventServiceV170SMTPAuthenticationMethods = "AutoDetect"
	PLAIN_ESV170SMTPAM EventServiceV170SMTPAuthenticationMethods = "Plain"
	LOGIN_ESV170SMTPAM EventServiceV170SMTPAuthenticationMethods = "Login"
	CRAM_MD5_ESV170SMTPAM EventServiceV170SMTPAuthenticationMethods = "CRAM_MD5"
)