/*
 * SendToM-Pesa
 *
 * Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to an M-Pesa mobile wallet recipient
 *
 * API version: 1.0.0
 *
 *
 */

package mpesa

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
