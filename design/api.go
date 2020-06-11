package design

import (
	. "github.com/wondenge/coop-go/responses/design"
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
		Path("/")
	})

	// 1. Account Balance Enquiry API: Enables you to enquire about your own
	// Co-operative Bank accounts' balance as at now for the specified account number.
	Method("AccountBalance", func() {
		Description("Post an Account Balance Enquiry Request")
		Payload(AccountBalancePayload)
		Result(AccountBalanceSuccessResponse)

		Error("not_allowed", ErrorResult, "Method Not Allowed")
		Error("timeout", ErrorResult, "Request Timeout")

		HTTP(func() {
			POST("Enquiry/AccountBalance/1.0.0")

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

	// 2. Account Full Statement Enquiry API will enable you to enquire about your
	// own Co-operative Bank accounts' full statement for the specified account
	// number, start date and end date.
	Method("AccountFullStatement", func() {})

	// 3. Account Mini Statement Enquiry API will enable you to enquire about your
	// own Co-operative Bank accounts' Mini statement for the specified account number.
	Method("AccountMiniStatement", func() {})

	// 4. Account Transactions Enquiry API will enable you to enquire about your
	// own Co-operative Bank accounts' latest transactions for the specified account
	// number and number of transactions to be returned.
	Method("AccountTransactions", func() {})

	// 5. Account Validation Enquiry API will enable you to validate a Co-operative
	// Bank account number.
	Method("AccountValidation", func() {})

	// 6. Exchange Rate Enquiry API will enable you to enquire about the current
	// SPOT exchange rate for the specified currencies.
	Method("ExchangeRate", func() {})

	// 7. Internal Funds Transfer Account to Account API will enable you to transfer
	// funds from your own Co-operative Bank account to other Co-operative Bank account(s).
	Method("IFTAccountToAccount", func() {})

	// 8. INSSimulation: This API is used to give instant notifications or alerts
	// on account activities(Debits,Credits) to the customer so that they can
	// reflect this in their accounting backend.
	Method("INSSimulation", func() {})

	// 9. PesaLink Send to Account Funds Transfer API will enable you to transfer funds
	// from your own Co-operative Bank account to Bank account(s) in IPSL participating banks.
	Method("PesaLinkSendToAccount", func() {})

	// 10. PesaLink Send to Phone Funds Transfer API will enable you to transfer funds
	// from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account
	// in an IPSL participating bank.
	Method("PesaLinkSendToPhone", func() {})

	// 11. Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own
	// Co-operative Bank account to an M-Pesa account recipient.
	Method("SendToMPesa", func() {})

	// 12. Transaction Status Enquiry API will enable you to enquire about the status of a
	// previously requested transaction for the specified transaction message reference.
	Method("TransactionStatus", func() {})
})
