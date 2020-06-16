package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post a Transaction Status Enquiry Request
func (s *connectsrvc) TransactionStatus(ctx context.Context, p *connect.FTTransactionStatusPayload) (res *connect.SuccessResponse, err error) {
	res = &connect.SuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.TransactionStatus"))
	return
}