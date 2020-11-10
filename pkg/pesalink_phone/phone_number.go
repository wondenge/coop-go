/*
 * PesaLinkSendToPhone
 *
 * PesaLink Send to Phone Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account in an IPSL participating bank
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_phone

// PhoneNumber Recipient phone number linked to a bank account in an IPSL participating bank
type PhoneNumber string

// List of PhoneNumber
const (
	_07XXXXXXXX   PhoneNumber = "07xxxxxxxx"
	_2547XXXXXXXX PhoneNumber = "2547xxxxxxxx"
	_2547XXXXXXXX PhoneNumber = "+2547xxxxxxxx"
)
