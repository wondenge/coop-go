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

	// 400
	Error("bad_request", ErrorResult, "Bad Request")

	// 404
	Error("not_found", ErrorResult, "Not Found")

	// 405
	Error("not_allowed", ErrorResult, "Method Not Allowed")

	// 408
	Error("timeout", ErrorResult, "Request Timeout")

	// 1. Account Balance Enquiry API: Enables you to enquire about your own
	// Co-operative Bank accounts' balance as at now for the specified account number.
	Method("AccountBalance", func() {
		Description("Post an Account Balance Enquiry Request")
		Payload(AccountBalancePayload)
		Result(AccountBalanceSuccessResponse)

		HTTP(func() {
			POST("Enquiry/AccountBalance/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 401
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
	Method("AccountFullStatement", func() {
		Description("Post an Account Full Statement Enquiry Request")
		Payload(AccountFullStatementPayload)
		Result(AccountFullStatementSuccessResponse)
		HTTP(func() {
			POST("Enquiry/FullStatement/Account/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)
		})
	})

	// 3. Account Mini Statement Enquiry API will enable you to enquire about your
	// own Co-operative Bank accounts' Mini statement for the specified account number.
	Method("AccountMiniStatement", func() {
		Description("Post an Account Mini Statement Enquiry Request")
		Payload(AccountMiniStatementPayload)
		Result(AccountMiniStatementSuccessResponse)
		HTTP(func() {
			POST("MiniStatement/Account/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

		})
	})

	// 4. Account Transactions Enquiry API will enable you to enquire about your
	// own Co-operative Bank accounts' latest transactions for the specified account
	// number and number of transactions to be returned.
	Method("AccountTransactions", func() {
		Description("Post an Account Transactions Enquiry Request")
		Payload(AccountTransactionsPayload)
		Result(AccountTransactionsSuccessResponse)
		HTTP(func() {
			POST("Enquiry/AccountTransactions/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)
		})
	})

	// 5. Account Validation Enquiry API will enable you to validate a Co-operative
	// Bank account number.
	Method("AccountValidation", func() {
		Description("Post an Account Validation Enquiry Request")
		Payload(AccountValidationPayload)
		Result(AccountValidationSuccessResponse)
		HTTP(func() {
			POST("Enquiry/Validation/Account/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)
		})
	})

	// 6. Exchange Rate Enquiry API will enable you to enquire about the current
	// SPOT exchange rate for the specified currencies.
	Method("ExchangeRate", func() {
		Description("Post an Exchange Rate Enquiry Request")
		Payload(ExchangeRatePayload)
		Result(ExchangeRateSuccessResponse)
		HTTP(func() {
			POST("Enquiry/ExchangeRate/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

		})
	})

	// 7. Internal Funds Transfer Account to Account API will enable you to transfer
	// funds from your own Co-operative Bank account to other Co-operative Bank account(s).
	Method("IFTAccountToAccount", func() {
		Description("Post an Internal Funds Transfer Account to Account Transaction")
		Payload(IFTAccountToAccountTXNRequest)
		Result(SuccessAcknowledgement)
		HTTP(func() {
			POST("FundsTransfer/Internal/A2A/2.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(AcknowledgementError400, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response(AcknowledgementError403, StatusForbidden)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

			// Status 409
			// RFC 7231, 6.5.8
			Response(AcknowledgementError409, StatusConflict)

		})
	})

	// 8. INSSimulation: This API is used to give instant notifications or alerts
	// on account activities(Debits,Credits) to the customer so that they can
	// reflect this in their accounting backend.
	Method("INSSimulation", func() {
		Description("Post a Debit/Credit Account Transaction Event Type Notification Simulation Request")
		Payload(INSTransactionSimulationRequest)
		Result(SuccessAcknowledgement)
		HTTP(func() {
			POST("Notifications/INS/Simulation/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(ErrorAcknowledgement, StatusUnauthorized)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

		})
	})

	// 9. PesaLink Send to Account Funds Transfer API will enable you to transfer funds
	// from your own Co-operative Bank account to Bank account(s) in IPSL participating banks.
	Method("PesaLinkSendToAccount", func() {
		Description("Post a PesaLink Funds Transfer Send to Account Transaction")
		Payload(PesaLinkSendToAccountTransactionRequest)
		Result(SuccessAcknowledgement)
		HTTP(func() {
			POST("FundsTransfer/External/A2A/PesaLink/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(AcknowledgementError400, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response(AcknowledgementError403, StatusForbidden)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

			// Status 409
			// RFC 7231, 6.5.8
			Response(AcknowledgementError409, StatusConflict)

		})
	})

	// 10. PesaLink Send to Phone Funds Transfer API will enable you to transfer funds
	// from your own Co-operative Bank account to a Phone Number(s) linked to a Bank account
	// in an IPSL participating bank.
	Method("PesaLinkSendToPhone", func() {
		Description("Post a PesaLink Funds Transfer Send to Phone Transaction")
		Payload(PesaLinkSendToPhoneTransactionRequest)
		Result(SuccessAcknowledgement)
		HTTP(func() {
			POST("FundsTransfer/External/A2M/PesaLink/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(AcknowledgementError400, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response(AcknowledgementError403, StatusForbidden)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

			// Status 409
			// RFC 7231, 6.5.8
			Response(AcknowledgementError409, StatusConflict)

		})
	})

	// 11. Send to M-Pesa Funds Transfer API will enable you to transfer funds from your own
	// Co-operative Bank account to an M-Pesa account recipient.
	Method("SendToMPesa", func() {
		Description("Post a Send To M-Pesa Funds Transfer Transaction")
		Payload(SendToMpesaTransactionRequest)
		Result(SuccessAcknowledgement)
		HTTP(func() {
			POST("FundsTransfer/External/A2M/Mpesa/v1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response(AcknowledgementError400, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response(AcknowledgementError403, StatusForbidden)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

			// Status 409
			// RFC 7231, 6.5.8
			Response(AcknowledgementError409, StatusConflict)

		})
	})

	// 12. Transaction Status Enquiry API will enable you to enquire about the status of a
	// previously requested transaction for the specified transaction message reference.
	Method("TransactionStatus", func() {
		Description("Post a Transaction Status Enquiry Request")
		Payload(FTTransactionStatusPayload)
		Result(SuccessResponse)
		HTTP(func() {
			POST("Enquiry/TransactionStatus/2.0.0")

			// Status 102
			// RFC 2518, 10.1
			Response(ProcessingResponse, StatusProcessing)

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK)

			// StatusMultiStatus
			// Status 207
			// RFC 4918, 11.1
			Response(MultiStatusResponse, StatusMultiStatus)

			// Status 400
			// RFC 7231, 6.5.1
			Response(ErrorResponse, StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response(MissingCredentials, StatusUnauthorized)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

		})
	})
})
