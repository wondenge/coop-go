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

// AccountValidationRequest struct for AccountValidationRequest
type AccountValidationRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
}
