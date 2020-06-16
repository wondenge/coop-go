package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post an Account Balance Enquiry Request
func (s *connectsrvc) AccountBalance(ctx context.Context, p *connect.AccountBalancePayload) (res *connect.AccountBalanceSuccessResponse, err error) {
	res = &connect.AccountBalanceSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountBalance"))
	return
}
