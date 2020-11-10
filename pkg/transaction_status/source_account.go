/*
 * TransactionStatus
 *
 * Transaction Status Enquiry API will enable you to enquire about the status of a previously requested transaction for the specified transaction message reference
 *
 * API version: 2.0.0
 *
 *
 */

package transaction_status

// SourceAccount struct for SourceAccount
type SourceAccount struct {
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// Transaction Amount
	Amount              float64             `json:"Amount"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `json:"Narration"`
	// Posting leg response code
	ResponseCode string `json:"ResponseCode,omitempty"`
	// Posting leg response description
	ResponseDescription string `json:"ResponseDescription,omitempty"`
}
