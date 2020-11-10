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

// SourceAccountCallbackRequest struct for SourceAccountCallbackRequest
type SourceAccountCallbackRequest struct {
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// Transaction Amount
	Amount              int32               `json:"Amount"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `json:"Narration"`
	// Posting leg response code
	ResponseCode string `json:"ResponseCode"`
	// Posting leg response description
	ResponseDescription string `json:"ResponseDescription"`
}
