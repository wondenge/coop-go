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

// DestinationAccountCallbackRequest struct for DestinationAccountCallbackRequest
type DestinationAccountCallbackRequest struct {
	// Unique posting reference for the transaction leg
	ReferenceNumber string `json:"ReferenceNumber"`
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// Posting account bank code
	BankCode string `json:"BankCode,omitempty"`
	// Transaction Amount
	Amount              int32               `json:"Amount"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `json:"Narration"`
	// Co-operative Bank's processed transaction number
	TransactionID string `json:"TransactionID,omitempty"`
	// Posting leg response code
	ResponseCode string `json:"ResponseCode"`
	// Posting leg response description
	ResponseDescription string `json:"ResponseDescription"`
}
