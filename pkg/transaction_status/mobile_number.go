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

// MobileNumber Recipient phone number linked to a bank account in an IPSL participating bank
type MobileNumber string

// List of MobileNumber
const (
	_2547XXXXXXXX MobileNumber = "2547xxxxxxxx"
	_07XXXXXXXX   MobileNumber = "07xxxxxxxx"
	_2547XXXXXXXX MobileNumber = "+2547xxxxxxxx"
)
