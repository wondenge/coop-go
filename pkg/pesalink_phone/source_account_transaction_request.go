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

// SourceAccountTransactionRequest struct for SourceAccountTransactionRequest
type SourceAccountTransactionRequest struct {
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// Transaction Amount
	Amount              int32               `json:"Amount"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `json:"Narration"`
}
