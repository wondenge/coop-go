/*
 * SendToM-Pesa
 *
 * Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to an M-Pesa mobile wallet recipient
 *
 * API version: 1.0.0
 *
 *
 */

package mpesa

// MobileNumber Recipient M-Pesa mobile number
type MobileNumber string

// List of MobileNumber
const (
	_2547XXXXXXXX MobileNumber = "2547xxxxxxxx"
	_07XXXXXXXX   MobileNumber = "07xxxxxxxx"
	_2547XXXXXXXX MobileNumber = "+2547xxxxxxxx"
)
