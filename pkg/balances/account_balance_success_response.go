/*
 * AccountBalance
 *
 * Account Balance Enquiry API will enable you to enquire about your own Co-operative Bank accounts' balance as at now for the specified account number
 *
 * API version: 1.0.0
 *
 *
 */

package balances

import (
	"time"
)

// AccountBalanceSuccessResponse struct for AccountBalanceSuccessResponse
type AccountBalanceSuccessResponse struct {
	// Your unique transaction request message identifier
	MessageReference string `json:"MessageReference"`
	// Acknowledgement message creation timestamp
	MessageDateTime time.Time `json:"MessageDateTime"`
	// Message Response Code
	MessageCode string `json:"MessageCode"`
	// Message Code description
	MessageDescription string `json:"MessageDescription"`
	// Posting account number
	AccountNumber string `json:"AccountNumber,omitempty"`
	// Account Name
	AccountName string   `json:"AccountName,omitempty"`
	Currency    Currency `json:"Currency,omitempty"`
	// Posting account number
	ProductName string `json:"ProductName,omitempty"`
	// Cleared Balance Amount
	ClearedBalance float64 `json:"ClearedBalance,omitempty"`
	// Cleared Balance Amount
	BookedBalance float64 `json:"BookedBalance,omitempty"`
	// Cleared Balance Amount
	BlockedBalance float64 `json:"BlockedBalance,omitempty"`
	// Cleared Balance Amount
	AvailableBalance float64 `json:"AvailableBalance,omitempty"`
	// Cleared Balance Amount
	ArrearsAmount float64 `json:"ArrearsAmount,omitempty"`
	// Posting account number
	BranchName string `json:"BranchName,omitempty"`
	// Posting account number
	BranchSortCode string `json:"BranchSortCode,omitempty"`
	// Cleared Balance Amount
	AverageBalance float64 `json:"AverageBalance,omitempty"`
	// Cleared Balance Amount
	UnclearedBalance float64 `json:"UnclearedBalance,omitempty"`
	// Cleared Balance Amount
	ODLimit float64 `json:"ODLimit,omitempty"`
	// Cleared Balance Amount
	CreditLimit float64 `json:"CreditLimit,omitempty"`
}
