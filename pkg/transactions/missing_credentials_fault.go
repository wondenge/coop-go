/*
 * AccountTransactions
 *
 * Account Transactions Enquiry API will enable you to enquire about your own Co-operative Bank accounts' latest transactions for the specified account number and number of transactions to be returned
 *
 * API version: 1.0.0
 *
 *
 */

package transactions

// MissingCredentialsFault struct for MissingCredentialsFault
type MissingCredentialsFault struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
