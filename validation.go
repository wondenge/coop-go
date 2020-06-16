package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post an Account Validation Enquiry Request
func (s *connectsrvc) AccountValidation(ctx context.Context, p *connect.AccountValidationPayload) (res *connect.AccountValidationSuccessResponse, err error) {
	res = &connect.AccountValidationSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountValidation"))
	return
}
