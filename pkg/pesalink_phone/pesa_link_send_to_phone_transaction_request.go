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

// PesaLinkSendToPhoneTransactionRequest struct for PesaLinkSendToPhoneTransactionRequest
type PesaLinkSendToPhoneTransactionRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Your callback URL that will receive transaction processing results
	CallBackUrl  string                                 `json:"CallBackUrl"`
	Source       SourceAccountTransactionRequest        `json:"Source"`
	Destinations []DestinationAccountTransactionRequest `json:"Destinations"`
}
