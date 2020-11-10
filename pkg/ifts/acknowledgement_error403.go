/*
 * IFTAccountToAccount
 *
 * Internal Funds Transfer Account to Account API will enable you to transfer funds from your own Co-operative Bank account to other Co-operative Bank account(s)
 *
 * API version: 2.0.0
 *
 *
 */

package ifts

import (
	"time"
)

// AcknowledgementError403 struct for AcknowledgementError403
type AcknowledgementError403 struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Acknowledgement/Response Message Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string `json:"MessageDescription"`
}
