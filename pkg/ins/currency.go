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

// Currency Transaction Posting account currency in ISO Currency Code
type Currency string

// List of Currency
const (
	KES Currency = "KES"
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	AUD Currency = "AUD"
	CHF Currency = "CHF"
	CAD Currency = "CAD"
	ZAR Currency = "ZAR"
)
