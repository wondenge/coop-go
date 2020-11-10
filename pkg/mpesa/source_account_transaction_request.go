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

// SourceAccountTransactionRequest struct for SourceAccountTransactionRequest
type SourceAccountTransactionRequest struct {
	// Posting leg response code
	ResponseCode string `json:"ResponseCode,omitempty"`
	// Transaction posting narration
	Narration string `json:"Narration"`
	// Transaction Amount
	Amount int32 `json:"Amount"`
	// Posting leg response description
	ResponseDescription string              `json:"ResponseDescription,omitempty"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Your posting account number
	AccountNumber string `json:"AccountNumber"`
}
