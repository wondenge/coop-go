/*
 * TransactionStatus
 *
 * Transaction Status Enquiry API will enable you to enquire about the status of a previously requested transaction for the specified transaction message reference
 *
 * API version: 2.0.0
 *
 *
 */

package transaction_status

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
