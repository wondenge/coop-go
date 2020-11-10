/*
 * TransactionStatus
 *
 * Transaction Status Enquiry API will enable you to enquire about the status of a previously requested transaction for the specified transaction message reference
 *
 * API version: 2.0.0
 *
 *
 */

package transaction_status

// PhoneNumber Recipient phone number linked to a bank account in an IPSL participating bank
type PhoneNumber string

// List of PhoneNumber
const (
	_2547XXXXXXXX PhoneNumber = "2547xxxxxxxx"
	_07XXXXXXXX   PhoneNumber = "07xxxxxxxx"
	_2547XXXXXXXX PhoneNumber = "+2547xxxxxxxx"
)
