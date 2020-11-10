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

// EventType The event of the transaction
type EventType string

// List of EventType
const (
	DEBIT  EventType = "DEBIT"
	CREDIT EventType = "CREDIT"
)
