// Code generated by goa v3.1.3, DO NOT EDIT.
//
// swagger endpoints
//
// Command:
// $ goa gen github.com/wondenge/coop-go/design

package swagger

import (
	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "swagger" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "swagger" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "swagger" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
}
