/*
 * AccountFullStatement
 *
 * Account Full Statement Enquiry API will enable you to enquire about your own Co-operative Bank accounts' full statement for the specified account number, start date and end date
 *
 * API version: 1.0.0
 *
 *
 */

package full_statement

// AccountTransaction struct for AccountTransaction
type AccountTransaction struct {
	// Unique Transaction Posting Identifier
	TransactionId string `json:"TransactionId,omitempty"`
	// Transaction Date
	TransactionDate string `json:"TransactionDate,omitempty"`
	// Transaction Value Date
	ValueDate string `json:"ValueDate,omitempty"`
	// Transaction Narration
	Narration string `json:"Narration,omitempty"`
	// Transaction Type, C for Credit or D for Debit
	TransactionType string `json:"TransactionType,omitempty"`
	// Transaction Service Point
	ServicePoint string `json:"ServicePoint,omitempty"`
	// Unique Transaction Reference
	TransactionReference string `json:"TransactionReference,omitempty"`
	// Transaction Credit Amount
	CreditAmount float64 `json:"CreditAmount,omitempty"`
	// Transaction Debit Amount
	DebitAmount float64 `json:"DebitAmount,omitempty"`
	// Account Running Cleared Balance
	RunningClearedBalance float64 `json:"RunningClearedBalance,omitempty"`
	// Account Running Book Balance
	RunningBookBalance float64 `json:"RunningBookBalance,omitempty"`
	// Account Debit Limit
	DebitLimit float64 `json:"DebitLimit,omitempty"`
	// Account Debit Limit Expiry Date
	LimitExpiryDate string `json:"LimitExpiryDate,omitempty"`
}
