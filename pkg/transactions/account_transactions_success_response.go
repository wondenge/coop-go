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

import (
	"time"
)

// AccountTransactionsSuccessResponse struct for AccountTransactionsSuccessResponse
type AccountTransactionsSuccessResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Message Response Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string `json:"MessageDescription"`
	// Posting account number
	AccountNumber string `json:"AccountNumber,omitempty"`
	// Account Name
	AccountName string `json:"AccountName,omitempty"`
	// No Of Latest Transactions To Be Returned
	NoOfTransactions float32 `json:"NoOfTransactions,omitempty"`
	// Total Credits Amount
	TotalCredits float64 `json:"TotalCredits,omitempty"`
	// Total Debits Amount
	TotalDebits  float64              `json:"TotalDebits,omitempty"`
	Transactions []AccountTransaction `json:"Transactions,omitempty"`
}
