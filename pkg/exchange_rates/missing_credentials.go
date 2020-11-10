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

// MissingCredentials struct for MissingCredentials
type MissingCredentials struct {
	Fault MissingCredentialsFault `json:"fault,omitempty"`
}
