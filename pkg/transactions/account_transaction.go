/*
 * AccountTransactions
 *
 * Account Transactions Enquiry API will enable you to enquire about your own Co-operative Bank accounts' latest transactions for the specified account number and number of transactions to be returned
 *
 * API version: 1.0.0
 *
 *
 */

package transactions

// AccountTransaction struct for AccountTransaction
type AccountTransaction struct {
	// Unique Transaction Posting Identifier
	TransactionID string `json:"TransactionID,omitempty"`
	// Transaction Date
	TransactionDate string `json:"TransactionDate,omitempty"`
	// Value Date
	ValueDate string `json:"ValueDate,omitempty"`
	// Transaction Narration
	Narration string `json:"Narration,omitempty"`
	// Transaction Type
	TransactionType string `json:"TransactionType,omitempty"`
	// Transaction Service Point
	ServicePoint string `json:"ServicePoint,omitempty"`
	// Transaction Reference
	TransactionReference string `json:"TransactionReference,omitempty"`
	// Transaction Credit Amount
	CreditAmount float64 `json:"CreditAmount,omitempty"`
	// Transaction Debit Amount
	DebitAmount float64 `json:"DebitAmount,omitempty"`
	// Running Cleared Balance
	RunningClearedBalance float64 `json:"RunningClearedBalance,omitempty"`
	// Running Book Balance
	RunningBookBalance float64 `json:"RunningBookBalance,omitempty"`
	// Account Debit Limit
	DebitLimit float64 `json:"DebitLimit,omitempty"`
	// Account Debit Limit Expiry Date
	LimitExpiryDate string `json:"LimitExpiryDate,omitempty"`
}
