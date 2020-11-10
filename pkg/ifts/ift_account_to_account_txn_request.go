/*
 * IFTAccountToAccount
 *
 * Internal Funds Transfer Account to Account API will enable you to transfer funds from your own Co-operative Bank account to other Co-operative Bank account(s)
 *
 * API version: 2.0.0
 *
 *
 */

package ifts

// IftAccountToAccountTxnRequest struct for IftAccountToAccountTxnRequest
type IftAccountToAccountTxnRequest struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Your callback URL that will receive transaction processing results
	CallBackUrl  string                         `json:"CallBackUrl"`
	Source       SourceAccountTxnRequest        `json:"Source"`
	Destinations []DestinationAccountTxnRequest `json:"Destinations"`
}
