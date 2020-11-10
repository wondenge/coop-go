/*
 * INSSimulation
 *
 * INS Simulation API enables you to simulate instant notifications reception on your specified notification endpoint
 *
 * API version: 1.0.0
 *
 *
 */

package ins

import (
	"time"
)

// SuccessAcknowledgement struct for SuccessAcknowledgement
type SuccessAcknowledgement struct {
	// Unique notification message identifier
	MessageReference string `json:"MessageReference"`
	// Notification message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Transaction request message code
	MessageCode string `json:"MessageCode"`
	// Transaction request message code description
	MessageDescription string `json:"MessageDescription"`
}
