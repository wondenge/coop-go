package connectapi

import (
	"context"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
)

// Post a Debit/Credit Account Transaction Event Type Notification Simulation
// Request
func (s *connectsrvc) INSSimulation(ctx context.Context, p *connect.INSTransactionSimulationRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.INSSimulation"))
	return
}
