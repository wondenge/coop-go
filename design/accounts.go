package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("enquiry", func() {

	// Bad Request Error Response-Synchronous:
	// -2=INVALID/MISSING PARAMETER,
	// -8=ACCOUNT AUTHORIZATION FAILURE,
	// -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH
	Error("bad_request", ErrorResult, "invalid data was sent in the request")

	// 401: Unauthorized Error Response-Synchronous:
	// 900901=Invalid credentials,
	// 900902=Missing credentials
	Error("unauthorized", ErrorResult, "900901=Invalid credentials, 900902=Missing credentials")

	// 405 : Method Not Allowed Error Response-Synchronous
	Error("not_allowed", ErrorResult, "Method Not Allowed")

	// 408: Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT
	Error("timeout", ErrorResult, "Timeout")

	HTTP(func() {
		Path("/Enquiry/AccountBalance/1.0.0")
	})

	Method("create", func() {
		Description("Post an Account Balance Enquiry Request")
		Payload(AccountBalanceRequest)
		Result(AccountBalanceSuccessResponse)
		HTTP(func() {
			POST("/")
			Response(StatusOK)
		})
	})
})
