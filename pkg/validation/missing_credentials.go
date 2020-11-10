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

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
