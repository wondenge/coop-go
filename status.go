package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 12. Transaction Status Enquiry API will enable you to enquire about the status
// of a previously requested transaction for the specified transaction message reference.
//
// Post a Transaction Status Enquiry Request
func (s *connectsrvc) TransactionStatus(ctx context.Context, p *connect.FTTransactionStatusPayload) (res *connect.SuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/Enquiry/TransactionStatus/2.0.0"), bytes.NewReader(b))
	res = &connect.SuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.TransactionStatus"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
