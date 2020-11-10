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

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
