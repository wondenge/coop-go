package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 5. Account Validation Enquiry API will enable you to validate a Co-operative Bank account number.
//
// Post an Account Validation Enquiry Request
func (s *connectsrvc) AccountValidation(ctx context.Context, p *connect.AccountValidationPayload) (res *connect.AccountValidationSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/Enquiry/Validation/Account/1.0.0"), bytes.NewReader(b))
	res = &connect.AccountValidationSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountValidation"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
