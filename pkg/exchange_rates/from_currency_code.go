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

// FromCurrencyCode From Currency Code in ISO Currency Code
type FromCurrencyCode string

// List of FromCurrencyCode
const (
	KES FromCurrencyCode = "KES"
	USD FromCurrencyCode = "USD"
	EUR FromCurrencyCode = "EUR"
	GBP FromCurrencyCode = "GBP"
	AUD FromCurrencyCode = "AUD"
	CHF FromCurrencyCode = "CHF"
	CAD FromCurrencyCode = "CAD"
	ZAR FromCurrencyCode = "ZAR"
)
