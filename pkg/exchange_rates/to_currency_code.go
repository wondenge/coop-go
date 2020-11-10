/*
 * ExchangeRate
 *
 * Exchange Rate Enquiry API will enable you to enquire about the current SPOT exchange rate for the specified currencies
 *
 * API version: 1.0.0
 *
 *
 */

package exchange_rates

// ToCurrencyCode To Currency Code in ISO Currency Code
type ToCurrencyCode string

// List of ToCurrencyCode
const (
	USD ToCurrencyCode = "USD"
	KES ToCurrencyCode = "KES"
	EUR ToCurrencyCode = "EUR"
	GBP ToCurrencyCode = "GBP"
	AUD ToCurrencyCode = "AUD"
	CHF ToCurrencyCode = "CHF"
	CAD ToCurrencyCode = "CAD"
	ZAR ToCurrencyCode = "ZAR"
)
