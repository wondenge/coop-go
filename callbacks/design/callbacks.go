package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("connect", func() {

	HTTP(func() {
		Path("/callbacks/coop-bank")
	})

	// StatusMultiStatus
	// Status 207
	// RFC 4918, 11.1
	Error("partial_success", ErrorResult, "1 - PARTIAL SUCCESS")
	Error("full_failure", ErrorResult, "2 - FULL FAILURE")

	// 7. Internal Funds Transfer Account to Account API will enable you to transfer
	// funds from your own Co-operative Bank account to other Co-operative Bank account(s).
	Method("IFTAccountToAccount", func() {
		Description("IFTAccountToAccount Early Hints Callback Request")
		Payload(IFTAccountToAccountCallbackRequest)
		Result(String)
		HTTP(func() {
			POST("/FundsTransfer/Internal/A2A/2.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 207
			// RFC 4918, 11.1
			Response("partial_success", StatusMultiStatus)

			// Status 207
			// RFC 4918, 11.1
			Response("full_failure", StatusMultiStatus)
		})

	})

	// 9. PesaLink Send to Account Funds Transfer API will enable you to transfer funds
	// from your own Co-operative Bank account to Bank account(s) in IPSL participating banks.
	Method("PesaLinkSendToAccount", func() {
		Description("PesaLinkSendToAccount Early Hints Callback Request")
		Payload(PesaLinkSendToAccountCallbackRequest)
		Result(String)
		HTTP(func() {
			POST("/FundsTransfer/External/A2A/PesaLink/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 207
			// RFC 4918, 11.1
			Response("partial_success", StatusMultiStatus)

			// Status 207
			// RFC 4918, 11.1
			Response("full_failure", StatusMultiStatus)
		})

	})

	// 10. PesaLink Send to Phone Funds Transfer API will enable you to transfer funds
	// from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account
	// in an IPSL participating bank.
	Method("PesaLinkSendToPhone", func() {
		Description("PesaLinkSendToPhone Early Hints Callback Request")
		Payload(PesaLinkSendToPhoneCallbackRequest)
		Result(String)
		HTTP(func() {
			POST("/FundsTransfer/External/A2M/PesaLink/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 207
			// RFC 4918, 11.1
			Response("partial_success", StatusMultiStatus)

			// Status 207
			// RFC 4918, 11.1
			Response("full_failure", StatusMultiStatus)
		})

	})

	// 11. Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own
	// Co-operative Bank account to an M-Pesa account recipient.
	Method("SendToMPesa", func() {
		Description("SendToMPesa Early Hints Callback Request")
		Payload(SendToMpesaCallbackRequest)
		Result(String)
		HTTP(func() {
			POST("/FundsTransfer/External/A2M/Mpesa/v1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 207
			// RFC 4918, 11.1
			Response("partial_success", StatusMultiStatus)

			// Status 207
			// RFC 4918, 11.1
			Response("full_failure", StatusMultiStatus)
		})
	})
})
