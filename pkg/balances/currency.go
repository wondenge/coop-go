/*
 * AccountBalance
 *
 * Account Balance Enquiry API will enable you to enquire about your own Co-operative Bank accounts' balance as at now for the specified account number
 *
 * API version: 1.0.0
 *
 *
 */

package balances

// Currency Account currency in ISO Currency Code
type Currency string

// List of Currency
const (
	KES Currency = "KES"
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	AUD Currency = "AUD"
	CHF Currency = "CHF"
	CAD Currency = "CAD"
	ZAR Currency = "ZAR"
)
