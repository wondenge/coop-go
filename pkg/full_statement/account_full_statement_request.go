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

// AccountFullStatementRequest struct for AccountFullStatementRequest
type AccountFullStatementRequest struct {
	// Your Unique Request Message Identifier
	MessageReference string `json:"MessageReference"`
	// Posting Account Number
	AccountNumber string `json:"AccountNumber"`
	// Statement Start Date
	StartDate string `json:"StartDate"`
	// Statement End Date
	EndDate string `json:"EndDate"`
}
