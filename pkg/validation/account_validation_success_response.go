/*
 * AccountValidation
 *
 * Account Validation Enquiry API will enable you to validate a Co-operative Bank account number
 *
 * API version: 1.0.0
 *
 *
 */

package validation

import (
	"time"
)

// AccountValidationSuccessResponse struct for AccountValidationSuccessResponse
type AccountValidationSuccessResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Message Response Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string `json:"MessageDescription"`
}
