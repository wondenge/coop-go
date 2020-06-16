package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post a Send To M-Pesa Funds Transfer Transaction
func (s *connectsrvc) SendToMPesa(ctx context.Context, p *connect.SendToMpesaTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.SendToMPesa"))
	return
}
