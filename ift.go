package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 7. Internal Funds Transfer Account to Account API will enable you to transfer
// funds from your own Co-operative Bank account to other Co-operative Bank account(s).
//
// Post an Internal Funds Transfer Account to Account Transaction
func (s *connectsrvc) IFTAccountToAccount(ctx context.Context, p *connect.IFTAccountToAccountTXNRequest) (res *connect.SuccessAcknowledgement, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/FundsTransfer/Internal/A2A/2.0.0"), bytes.NewReader(b))
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.IFTAccountToAccount"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
