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

// MissingCredentialsFault struct for MissingCredentialsFault
type MissingCredentialsFault struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
