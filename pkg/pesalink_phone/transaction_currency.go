/*
 * PesaLinkSendToPhone
 *
 * PesaLink Send to Phone Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account in an IPSL participating bank
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_phone

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
