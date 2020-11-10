/*
 * SendToM-Pesa
 *
 * Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to an M-Pesa mobile wallet recipient
 *
 * API version: 1.0.0
 *
 *
 */

package mpesa

// MissingCredentialsFault struct for MissingCredentialsFault
type MissingCredentialsFault struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}
