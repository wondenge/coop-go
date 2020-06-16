package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 1. Account Balance Enquiry API enables you to enquire about your own
// Co-operative Bank accounts' balance as at now for the specified account number.
//
// Post an Account Balance Enquiry Request
func (c *Client) AccountBalance(ctx context.Context, p *connect.AccountBalancePayload) (res *connect.AccountBalanceSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, "/Enquiry/AccountBalance/1.0.0"), bytes.NewReader(b))
	res = &connect.AccountBalanceSuccessResponse{}
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
