/*
 * PesaLinkSendToAccount
 *
 * PesaLink Send to Account Funds Transfer API will enable you to transfer funds from your own Co-operative Bank account to Bank account(s) in IPSL participating banks
 *
 * API version: 1.0.0
 *
 *
 */

package pesalink_account

// PesaLinkSendToAccountTransactionRequest struct for PesaLinkSendToAccountTransactionRequest
type PesaLinkSendToAccountTransactionRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Your callback URL that will receive transaction processing results
	CallBackUrl  string                                 `json:"CallBackUrl"`
	Source       SourceAccountTransactionRequest        `json:"Source"`
	Destinations []DestinationAccountTransactionRequest `json:"Destinations"`
}
