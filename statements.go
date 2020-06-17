package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 2. Account Full Statement Enquiry API will enable you to enquire about your own
// Co-operative Bank accounts' full statement for the specified account number,
// start date and end date.
//
// Post an Account Full Statement Enquiry Request
func (s *connectsrvc) AccountFullStatement(ctx context.Context, p *connect.AccountFullStatementPayload) (res *connect.AccountFullStatementSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/Enquiry/FullStatement/Account/1.0.0"), bytes.NewReader(b))
	res = &connect.AccountFullStatementSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountFullStatement"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}

// 3. Account Mini Statement Enquiry API will enable you to enquire about your
// own Co-operative Bank accounts' Mini statement for the specified account number.
//
// Post an Account Mini Statement Enquiry Request
func (s *connectsrvc) AccountMiniStatement(ctx context.Context, p *connect.AccountMiniStatementPayload) (res *connect.AccountMiniStatementSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/MiniStatement/Account/1.0.0"), bytes.NewReader(b))
	res = &connect.AccountMiniStatementSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountMiniStatement"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
