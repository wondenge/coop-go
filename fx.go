package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 6. Exchange Rate Enquiry API will enable you to enquire about the current
// SPOT exchange rate for the specified currencies.
//
// Post an Exchange Rate Enquiry Request
func (c *Client) ExchangeRate(ctx context.Context, p *connect.ExchangeRatePayload) (res *connect.ExchangeRateSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, "/Enquiry/ExchangeRate/1.0.0"), bytes.NewReader(b))
	res = &connect.ExchangeRateSuccessResponse{}
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
