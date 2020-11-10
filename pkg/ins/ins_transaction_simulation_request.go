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

// InsTransactionSimulationRequest struct for InsTransactionSimulationRequest
type InsTransactionSimulationRequest struct {
	// Unique notification message identifier
	MessageReference string `json:"MessageReference"`
	// Notification message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Notification Service Name identifier
	ServiceName string `json:"ServiceName,omitempty"`
	// Notification Code identifier
	NotificationCode string `json:"NotificationCode,omitempty"`
	// Transaction Payment Reference
	PaymentRef string `json:"PaymentRef"`
	// Posting account number
	AccountNumber string `json:"AccountNumber"`
	// Transaction Amount
	Amount float64 `json:"Amount"`
	// Posting date of the Transaction
	TransactionDate string    `json:"TransactionDate"`
	EventType       EventType `json:"EventType"`
	Currency        Currency  `json:"Currency"`
	// Exchange Rate
	ExchangeRate float64 `json:"ExchangeRate"`
	// Transaction Posting account narration
	Narration string   `json:"Narration"`
	CustMemo  CustMemo `json:"CustMemo"`
	// Transaction Posting Value Date
	ValueDate string `json:"ValueDate"`
	// Transaction Posting Entry Date
	EntryDate string `json:"EntryDate"`
	// Co-operative Bank's processed transaction number
	TransactionId string `json:"TransactionId"`
}
