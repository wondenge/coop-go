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

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Message Response Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string `json:"MessageDescription"`
}
