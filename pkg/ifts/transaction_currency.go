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

// TransactionCurrency Posting account currency in ISO Currency Code
type TransactionCurrency string

// List of TransactionCurrency
const (
	KES TransactionCurrency = "KES"
	USD TransactionCurrency = "USD"
	EUR TransactionCurrency = "EUR"
	GBP TransactionCurrency = "GBP"
	AUD TransactionCurrency = "AUD"
	CHF TransactionCurrency = "CHF"
	CAD TransactionCurrency = "CAD"
	ZAR TransactionCurrency = "ZAR"
)
