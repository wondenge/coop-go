package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post an Account Transactions Enquiry Request
func (s *connectsrvc) AccountTransactions(ctx context.Context, p *connect.AccountTransactionsPayload) (res *connect.AccountTransactionsSuccessResponse, err error) {
	res = &connect.AccountTransactionsSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountTransactions"))
	return
}