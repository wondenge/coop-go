/*
 * INSSimulation
 *
 * INS Simulation API enables you to simulate instant notifications reception on your specified notification endpoint
 *
 * API version: 1.0.0
 *
 *
 */

package ins

// CustMemo struct for CustMemo
type CustMemo struct {
	// CustMemo CustMemoLine1
	CustMemoLine1 string `json:"CustMemoLine1,omitempty"`
	// CustMemo CustMemoLine2
	CustMemoLine2 string `json:"CustMemoLine2,omitempty"`
	// CustMemo CustMemoLine2
	CustMemoLine3 string `json:"CustMemoLine3,omitempty"`
}
