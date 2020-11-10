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

// DestinationAccount struct for DestinationAccount
type DestinationAccount struct {
	// Unique posting reference for the transaction leg
	ReferenceNumber string `json:"ReferenceNumber"`
	// Posting account number
	AccountNumber string       `json:"AccountNumber"`
	MobileNumber  MobileNumber `json:"MobileNumber,omitempty"`
	PhoneNumber   PhoneNumber  `json:"PhoneNumber,omitempty"`
	// Posting account bank code
	BankCode string `json:"BankCode,omitempty"`
	// Transaction Amount
	Amount              float64             `json:"Amount"`
	TransactionCurrency TransactionCurrency `json:"TransactionCurrency"`
	// Posting account transaction narration
	Narration string `json:"Narration"`
	// Co-operative Bank's processed transaction number
	TransactionID string `json:"TransactionID,omitempty"`
	// Posting leg response code
	ResponseCode string `json:"ResponseCode,omitempty"`
	// Posting leg response description
	ResponseDescription string `json:"ResponseDescription,omitempty"`
}
