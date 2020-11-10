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

// NotFound struct for NotFound
type NotFound struct {
	Timestamp string `json:"timestamp,omitempty"`
	Status    string `json:"status,omitempty"`
	Error     string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
	Path      string `json:"path,omitempty"`
}
