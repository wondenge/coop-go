package design

import (
	. "github.com/wondenge/coop-go/messages/design"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("connect", func() {
	Title("Co-op Bank Connect API")
	Description("HTTP Service for accessing Co-op Bank Connect API.")
	Version("1.0")
	TermsOfService("https://github.com/wondenge/coop-go/blob/master/LICENCE")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/coop-go/blob/master/LICENCE")
	})
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/coop-go/blob/master/README.md")
	})
	Server("connect", func() {
		Description("connect hosts Co-op Bank Connect API Services.")
		Services()
		Host("development", func() {
			Description("Development Hosts")
			URI("http://developer.co-opbank.co.ke:8280")
		})
		Host("production", func() {
			Description("Production Host")
			URI("https://developer.co-opbank.co.ke:8243")
		})
	})
})

var _ = Service("coop", func() {

	HTTP(func() {
		Path("")
	})

	// 1. Account Balance Enquiry API:
	// Enables you to enquire about your own Co-operative Bank accounts'
	// balance as at now for the specified account number
	Method("AccountBalance", func() {
		Description("Post an Account Balance Enquiry Request")
		Payload(AccountBalancePayload)
		Result(AccountBalanceSuccessResponse)

		Error("not_allowed", ErrorResult, "Method Not Allowed")
		Error("timeout", ErrorResult, "Request Timeout")

		HTTP(func() {
			POST("/Enquiry/AccountBalance/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 400
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)
		})
	})
})
