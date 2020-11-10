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

// InvalidCredentials struct for InvalidCredentials
type InvalidCredentials struct {
	Fault InvalidCredentialsFault `json:"fault,omitempty"`
}
