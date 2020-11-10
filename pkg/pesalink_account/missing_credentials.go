/*
 * PesaLinkSendToAccount
 *
 * PesaLink Send to Account Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to Bank account(s) in IPSL participating banks
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_account

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
