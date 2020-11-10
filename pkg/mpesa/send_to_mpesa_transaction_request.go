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

// SendToMpesaTransactionRequest struct for SendToMpesaTransactionRequest
type SendToMpesaTransactionRequest struct {
	// Your callback endpoint that will receive transaction processing results
	CallBackUrl  string                                 `json:"CallBackUrl"`
	Destinations []DestinationAccountTransactionRequest `json:"Destinations"`
	// Your unique transaction request message identifier
	MessageReference string                          `json:"MessageReference"`
	Source           SourceAccountTransactionRequest `json:"Source"`
}
