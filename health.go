package connectapi

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	health "github.com/wondenge/coop-go/gen/health"
)

// health service example implementation.
// The example methods log the requests and return zero values.
type healthsrvc struct {
	logger log.Logger
}

// NewHealth returns the health service implementation.
func NewHealth(logger log.Logger) health.Service {
	return &healthsrvc{logger}
}

// Health check endpoint.
func (s *healthsrvc) Show(ctx context.Context) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("health.show"))
	return
}
