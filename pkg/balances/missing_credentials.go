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

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
