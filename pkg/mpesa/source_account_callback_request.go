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

// SourceAccountCallbackRequest struct for SourceAccountCallbackRequest
type SourceAccountCallbackRequest struct {
	// Posting leg response code
	ResponseCode string `json:"ResponseCode"`
	// Transaction posting narration
	Narration string `json:"Narration"`
	// Transaction Amount
	Amount int32 `json:"Amount"`
	// Posting leg response description
	ResponseDescription string              `json:"ResponseDescription"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Your posting account number
	AccountNumber string `json:"AccountNumber"`
}
