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

// MissingCredentialsFault struct for MissingCredentialsFault
type MissingCredentialsFault struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
