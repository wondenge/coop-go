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

// MissingCredentialsFault struct for MissingCredentialsFault
type MissingCredentialsFault struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
