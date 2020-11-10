/*
 * TransactionStatus
 *
 * Transaction Status Enquiry API will enable you to enquire about the status of a previously requested transaction for the specified transaction message reference
 *
 * API version: 2.0.0
 *
 *
 */

package transaction_status

import (
	"time"
)

// SuccessResponse struct for SuccessResponse
type SuccessResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Transaction request message code
	MessageCode string `json:"MessageCode"`
	// Transaction request message code description
	MessageDescription string               `json:"MessageDescription"`
	Source             SourceAccount        `json:"Source"`
	Destinations       []DestinationAccount `json:"Destinations"`
}
