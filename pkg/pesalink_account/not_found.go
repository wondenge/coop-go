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

// NotFound struct for NotFound
type NotFound struct {
	Timestamp string `json:"timestamp,omitempty"`
	Status    string `json:"status,omitempty"`
	Error     string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
	Path      string `json:"path,omitempty"`
}
