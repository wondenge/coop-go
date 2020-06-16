package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post a PesaLink Funds Transfer Send to Account Transaction
func (s *connectsrvc) PesaLinkSendToAccount(ctx context.Context, p *connect.PesaLinkSendToAccountTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToAccount"))
	return
}

// Post a PesaLink Funds Transfer Send to Phone Transaction
func (s *connectsrvc) PesaLinkSendToPhone(ctx context.Context, p *connect.PesaLinkSendToPhoneTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToPhone"))
	return
}
