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

// DestinationAccountTransactionRequest struct for DestinationAccountTransactionRequest
type DestinationAccountTransactionRequest struct {
	MobileNumber MobileNumber `json:"MobileNumber"`
	// Transaction posting narration
	Narration string `json:"Narration"`
	// Unique posting reference for the transaction leg
	ReferenceNumber string `json:"ReferenceNumber"`
	// Transaction Amount
	Amount int32 `json:"Amount"`
}
