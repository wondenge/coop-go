package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)


// Post an Account Full Statement Enquiry Request
func (s *connectsrvc) AccountFullStatement(ctx context.Context, p *connect.AccountFullStatementPayload) (res *connect.AccountFullStatementSuccessResponse, err error) {
	res = &connect.AccountFullStatementSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountFullStatement"))
	return
}

// Post an Account Mini Statement Enquiry Request
func (s *connectsrvc) AccountMiniStatement(ctx context.Context, p *connect.AccountMiniStatementPayload) (res *connect.AccountMiniStatementSuccessResponse, err error) {
	res = &connect.AccountMiniStatementSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountMiniStatement"))
	return
}