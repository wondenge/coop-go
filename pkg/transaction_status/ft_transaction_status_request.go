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

// FtTransactionStatusRequest struct for FtTransactionStatusRequest
type FtTransactionStatusRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
}
