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

// ExchangeRateRequest struct for ExchangeRateRequest
type ExchangeRateRequest struct {
	// Your unique transaction request message identifier
	MessageReference string           `json:"MessageReference"`
	FromCurrencyCode FromCurrencyCode `json:"FromCurrencyCode"`
	ToCurrencyCode   ToCurrencyCode   `json:"ToCurrencyCode"`
}
