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

import (
	"time"
)

// AccountFullStatementSuccessResponse struct for AccountFullStatementSuccessResponse
type AccountFullStatementSuccessResponse struct {
	// Your Unique Request Message Identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement/Response Message Creation Timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Acknowledgement/Response Message Code
	MessageCode string `json:"MessageCode"`
	// Acknowledgement/Response Message Code Description
	MessageDescription string `json:"MessageDescription"`
	// Posting Account Number
	AccountNumber string `json:"AccountNumber,omitempty"`
	// Posting Account Name
	AccountName  string               `json:"AccountName,omitempty"`
	Transactions []AccountTransaction `json:"Transactions,omitempty"`
}
