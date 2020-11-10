/*
 * AccountFullStatement
 *
 * Account Full Statement Enquiry API will enable you to enquire about your own Co-operative Bank accounts' full statement for the specified account number, start date and end date
 *
 * API version: 1.0.0
 *
 *
 */

package full_statement

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
