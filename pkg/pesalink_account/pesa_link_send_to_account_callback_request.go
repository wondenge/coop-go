/*
 * PesaLinkSendToAccount
 *
 * PesaLink Send to Account Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to Bank account(s) in IPSL participating banks
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_account

import (
	"time"
)

// PesaLinkSendToAccountCallbackRequest struct for PesaLinkSendToAccountCallbackRequest
type PesaLinkSendToAccountCallbackRequest struct {
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
