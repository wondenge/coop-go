/*
 * AccountBalance
 *
 * Account Balance Enquiry API will enable you to enquire about your own Co-operative Bank accounts' balance as at now for the specified account number
 *
 * API version: 1.0.0
 *
 */

package balances

// AccountBalanceRequest struct for AccountBalanceRequest
type AccountBalanceRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
}
