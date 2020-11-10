/*
 * IFTAccountToAccount
 *
 * Internal Funds Transfer Account to Account API will enable you to transfer funds from your own Co-operative Bank account to other Co-operative Bank account(s)
 *
 * API version: 2.0.0
 *
 *
 */

package ifts

// SourceAccountTxnRequest struct for SourceAccountTxnRequest
type SourceAccountTxnRequest struct {
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// Transaction Amount
	Amount              float64             `json:"Amount"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `json:"Narration"`
}
