// Code generated by goa v3.1.3, DO NOT EDIT.
//
// health service
//
// Command:
// $ goa gen github.com/wondenge/coop-go/design

package health

import (
	"context"
)

// Service is the health service interface.
type Service interface {
	// Health check endpoint.
	Show(context.Context) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "health"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"show"}