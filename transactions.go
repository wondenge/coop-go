package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 4. Account Transactions Enquiry API will enable you to enquire about your own
// Co-operative Bank accounts' latest transactions for the specified account number
// and number of transactions to be returned.
//
// Post an Account Transactions Enquiry Request
func (c *Client) AccountTransactions(ctx context.Context, p *connect.AccountTransactionsPayload) (res *connect.AccountTransactionsSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, "/Enquiry/AccountTransactions/1.0.0"), bytes.NewReader(b))
	res = &connect.AccountTransactionsSuccessResponse{}
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
