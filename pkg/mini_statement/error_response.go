/*
 * AccountMiniStatement
 *
 * Account Mini Statement Enquiry API will enable you to enquire about your own Co-operative Bank accounts' Mini statement for the specified account number
 *
 * API version: 1.0.0
 *
 *
 */

package mini_statement

import (
	"time"
)

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// Your Unique Request Message Identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement/Response Message Creation Timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Acknowledgement/Response Message Code
	MessageCode string `json:"MessageCode"`
	// Acknowledgement/Response Message Code Description
	MessageDescription string `json:"MessageDescription"`
}
