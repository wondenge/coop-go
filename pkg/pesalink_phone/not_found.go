/*
 * PesaLinkSendToPhone
 *
 * PesaLink Send to Phone Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account in an IPSL participating bank
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_phone

// NotFound struct for NotFound
type NotFound struct {
	Timestamp string `json:"timestamp,omitempty"`
	Status    string `json:"status,omitempty"`
	Error     string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
	Path      string `json:"path,omitempty"`
}
