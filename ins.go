package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"net/http"
)

// 8. INSSimulation: This API is used to give instant notifications or alerts on
// account activities(Debits,Credits) to the customer so that they can reflect
// this in their accounting backend.
// Post a Debit/Credit Account Transaction Event Type Notification Simulation Request
func (s *connectsrvc) INSSimulation(ctx context.Context, p *connect.INSTransactionSimulationRequest) (res *connect.SuccessAcknowledgement, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "/Notifications/INS/Simulation/1.0.0"), bytes.NewReader(b))
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.INSSimulation"))
	if err != nil {
		return nil, err
	}
	err = c.SendWithAuth(req, res)
	return res, err
}
