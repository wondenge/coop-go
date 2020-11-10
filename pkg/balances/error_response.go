/*
 * AccountBalance
 *
 * Account Balance Enquiry API will enable you to enquire about your own Co-operative Bank accounts' balance as at now for the specified account number
 *
 * API version: 1.0.0
 *
 *
 */

package balances

import (
	"time"
)

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Message Response Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string `json:"MessageDescription"`
}
