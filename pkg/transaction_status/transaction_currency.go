/*
 * TransactionStatus
 *
 * Transaction Status Enquiry API will enable you to enquire about the status of a previously requested transaction for the specified transaction message reference
 *
 * API version: 2.0.0
 *
 *
 */

package transaction_status

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
