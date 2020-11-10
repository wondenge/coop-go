/*
 * AccountMiniStatement
 *
 * Account Mini Statement Enquiry API will enable you to enquire about your own Co-operative Bank accounts' Mini statement for the specified account number
 *
 * API version: 1.0.0
 *
 *
 */

package mini_statement

// AccountMiniStatementRequest struct for AccountMiniStatementRequest
type AccountMiniStatementRequest struct {
	// Your Unique Request Message Identifier
	MessageReference string `json:"MessageReference"`
	// Posting Account Number
	AccountNumber string `json:"AccountNumber"`
}
