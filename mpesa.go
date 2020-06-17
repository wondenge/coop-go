package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 11. Send to M-Pesa Funds Transfer API will enable you to transfer funds from
// your own Co-operative Bank account to an M-Pesa account recipient.
//
// Post a Send To M-Pesa Funds Transfer Transaction
func (s *connectsrvc) SendToMPesa(ctx context.Context, p *connect.SendToMpesaTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/FundsTransfer/External/A2M/Mpesa/v1.0.0"), bytes.NewReader(b))
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.SendToMPesa"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
