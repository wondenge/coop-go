/*
 * PesaLinkSendToPhone
 *
 * PesaLink Send to Phone Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account in an IPSL participating bank
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_phone

import (
	"time"
)

// PesaLinkSendToPhoneCallbackRequest struct for PesaLinkSendToPhoneCallbackRequest
type PesaLinkSendToPhoneCallbackRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Acknowledgement/Response Message Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string                              `json:"MessageDescription"`
	Source             SourceAccountCallbackRequest        `json:"Source"`
	Destinations       []DestinationAccountCallbackRequest `json:"Destinations"`
}
