/*
 * SendToM-Pesa
 *
 * Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to an M-Pesa mobile wallet recipient
 *
 * API version: 1.0.0
 *
 *
 */

package mpesa

import (
	"time"
)

// SendToMpesaCallbackRequest struct for SendToMpesaCallbackRequest
type SendToMpesaCallbackRequest struct {
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Acknowledgement/Response Message Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string                              `json:"MessageDescription"`
	Destinations       []DestinationAccountCallbackRequest `json:"Destinations"`
	// Your unique transaction request message identifier
	MessageReference string                       `json:"MessageReference"`
	Source           SourceAccountCallbackRequest `json:"Source"`
}
