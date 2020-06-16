package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post an Exchange Rate Enquiry Request
func (s *connectsrvc) ExchangeRate(ctx context.Context, p *connect.ExchangeRatePayload) (res *connect.ExchangeRateSuccessResponse, err error) {
	res = &connect.ExchangeRateSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.ExchangeRate"))
	return
}
