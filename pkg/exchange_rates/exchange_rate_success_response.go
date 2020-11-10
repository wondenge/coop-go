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

import (
	"time"
)

// ExchangeRateSuccessResponse struct for ExchangeRateSuccessResponse
type ExchangeRateSuccessResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Message Response Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string           `json:"MessageDescription"`
	FromCurrencyCode   FromCurrencyCode `json:"FromCurrencyCode,omitempty"`
	ToCurrencyCode     ToCurrencyCode   `json:"ToCurrencyCode,omitempty"`
	// Exchange rate type
	RateType string `json:"RateType,omitempty"`
	// Exchange rate amount
	Rate float64 `json:"Rate,omitempty"`
	// Exchange rate tolerance
	Tolerance float64 `json:"Tolerance,omitempty"`
	// Exchange rate is a multiply or a divide
	MultiplyDivide string `json:"MultiplyDivide,omitempty"`
}
