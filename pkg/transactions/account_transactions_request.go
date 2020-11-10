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

// AccountTransactionsRequest struct for AccountTransactionsRequest
type AccountTransactionsRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// No Of Latest Transactions To Be Returned
	NoOfTransactions float32 `json:"NoOfTransactions"`
}
