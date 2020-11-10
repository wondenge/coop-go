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

// DestinationAccountCallbackRequest struct for DestinationAccountCallbackRequest
type DestinationAccountCallbackRequest struct {
	MobileNumber MobileNumber `json:"MobileNumber"`
	// Posting leg response code
	ResponseCode string `json:"ResponseCode,omitempty"`
	// Transaction posting narration
	Narration string `json:"Narration"`
	// Unique posting reference for the transaction leg
	ReferenceNumber string `json:"ReferenceNumber"`
	// Transaction Amount
	Amount int32 `json:"Amount"`
	// Posting leg response description
	ResponseDescription string `json:"ResponseDescription,omitempty"`
	// Co-operative Bank's processed transaction number concatenated with M-Pesa transaction number
	TransactionID string `json:"TransactionID,omitempty"`
}
