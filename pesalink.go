package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 9. PesaLink Send to Account Funds Transfer API will enable you to transfer funds
// from your own Co-operative Bank account to Bank account(s) in IPSL participating banks.
//
// Post a PesaLink Funds Transfer Send to Account Transaction
func (s *connectsrvc) PesaLinkSendToAccount(ctx context.Context, p *connect.PesaLinkSendToAccountTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/FundsTransfer/External/A2A/PesaLink/1.0.0"), bytes.NewReader(b))
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToAccount"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}

// 10. PesaLink Send to Phone Funds Transfer API will enable you to transfer funds
// from your own Co-operative Bank account to a Phone Number(s) linked to a Bank
// account in an IPSL participating bank.
//
// Post a PesaLink Funds Transfer Send to Phone Transaction
func (s *connectsrvc) PesaLinkSendToPhone(ctx context.Context, p *connect.PesaLinkSendToPhoneTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/FundsTransfer/External/A2M/PesaLink/1.0.0"), bytes.NewReader(b))
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToPhone"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
