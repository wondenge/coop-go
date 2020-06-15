package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("connect", func() {

	HTTP(func() {
		Path("/")
	})

	// 1. Account Balance Enquiry API enables you to enquire about
	// your own Co-operative Bank accounts' balance as at now for
	// the specified account number.
	Method("AccountBalance", func() {
		Description("Post an Account Balance Enquiry Request")
		Payload(AccountBalancePayload)
		Result(AccountBalanceSuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: -2=INVALID/MISSING PARAMETER, -8=ACCOUNT AUTHORIZATION FAILURE, -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: 900901=Invalid credentials, 900902=Missing credentials")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT")

		HTTP(func() {
			POST("Enquiry/AccountBalance/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Success Response-Synchronous: 0=Success")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

		})
	})

	// 2. Account Full Statement Enquiry API will enable you to enquire
	// about your own Co-operative Bank accounts' full statement for the
	// specified account number, start date and end date.
	Method("AccountFullStatement", func() {
		Description("Post an Account Full Statement Enquiry Request")
		Payload(AccountFullStatementPayload)
		Result(AccountFullStatementSuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: -2=INVALID/MISSING PARAMETER, -6=Date Fornat, -8=ACCOUNT AUTHORIZATION FAILURE, -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: 900901=Invalid credentials, 900902=Missing credentials")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT")

		HTTP(func() {
			POST("Enquiry/FullStatement/Account/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Success Response-Synchronous: 0=Success")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

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

	// 3. Account Mini Statement Enquiry API will enable you to enquire
	// about your own Co-operative Bank accounts' Mini statement for the
	// specified account number.
	Method("AccountMiniStatement", func() {
		Description("Post an Account Mini Statement Enquiry Request")
		Payload(AccountMiniStatementPayload)
		Result(AccountMiniStatementSuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: -2=INVALID/MISSING PARAMETER, -6=Date Fornat, -8=ACCOUNT AUTHORIZATION FAILURE, -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: 900901=Invalid credentials, 900902=Missing credentials")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT")

		HTTP(func() {
			POST("MiniStatement/Account/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Success Response-Synchronous: 0=Success")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

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

	// 4. Account Transactions Enquiry API will enable you to enquire
	// about your own Co-operative Bank accounts' latest transactions
	// for the specified account number and number of transactions to
	// be returned.
	Method("AccountTransactions", func() {
		Description("Post an Account Transactions Enquiry Request")
		Payload(AccountTransactionsPayload)
		Result(AccountTransactionsSuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: -2=INVALID/MISSING PARAMETER, -6=NO OF TRANSACTIONS IS LESS/GREATER THAN LIMIT ALLOWED, -8=ACCOUNT AUTHORIZATION FAILURE, -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: 900901=Invalid credentials, 900902=Missing credentials")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT")

		HTTP(func() {
			POST("Enquiry/AccountTransactions/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Success Response-Synchronous: 0=Success")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)
		})
	})

	// 5. Account Validation Enquiry API will enable you to validate
	// a Co-operative Bank account number.
	Method("AccountValidation", func() {
		Description("Post an Account Validation Enquiry Request")
		Payload(AccountValidationPayload)
		Result(AccountValidationSuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: -2=INVALID/MISSING PARAMETER, -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH, -15=INVALID ACCOUNT NUMBER")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: 900901=Invalid credentials, 900902=Missing credentials")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT")

		HTTP(func() {
			POST("Enquiry/Validation/Account/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Success Response-Synchronous: 0=VALID ACCOUNT NUMBER")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

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

	// 6. Exchange Rate Enquiry API will enable you to enquire about
	// the current SPOT exchange rate for the specified currencies.
	Method("ExchangeRate", func() {
		Description("Post an Exchange Rate Enquiry Request")
		Payload(ExchangeRatePayload)
		Result(ExchangeRateSuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: -2=INVALID/MISSING PARAMETER, -11=MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: 900901=Invalid credentials, 900902=Missing credentials")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous: -4=Record Not Found")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: -4=REQUEST TIMED OUT")

		HTTP(func() {
			POST("Enquiry/ExchangeRate/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Success Response-Synchronous: 0=Success")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

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

	// 7. Internal Funds Transfer Account to Account API will enable
	// you to transfer funds from your own Co-operative Bank account
	// to other Co-operative Bank account(s).
	Method("IFTAccountToAccount", func() {
		Description("Post an Internal Funds Transfer Account to Account Transaction")
		Payload(IFTAccountToAccountTXNRequest)
		Result(SuccessAcknowledgement)

		// 400
		Error("bad_request", AcknowledgementError400, "Bad Request Error Response-Synchronous: <br>-2 - INVALID/MISSING PARAMETER, <br>-3 - THE AMOUNT DEBITED OR CREDITED MUST BE GREATER THAT 0/POSITIVE INTEGER, <br>-5 - DEBIT AND CREDIT(S) AMOUNTS NOT BALANCING, <br>-6 - TRANSACTION AMOUNT LESS/GREATER THAN MINIMUM/MAXIMUM LIMIT ALLOWED, <br>-11 - MESSAGE REFERENCE/REFERENCE NUMBER LONGER THAN ALLOWED LENGTH, <br>-12 - DUPLICATE REFERENCE/IDENTICAL REFERENCE AND MESSAGEREFERENCE")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: <br>900901 - Invalid credentials, <br>900902 - Missing credentials, <br>-8 - DEBIT ACCOUNT AUTHORIZATION FAILURE")

		// 403
		Error("forbidden", AcknowledgementError403, "Forbidden Error Response-Synchronous: <br>-9 - CURRENCY INVALID/NOT ALLOWED, <br>-16 - DAILY LIMIT EXHAUSTED")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous: <br>Wrong API Resource URI")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: <br>-4 - REQUEST TIMED OUT")

		// 409
		Error("conflict", AcknowledgementError409, "Conflict Error Response-Synchronous: <br>-1 - DUPLICATE MESSAGE REFERENCE")

		HTTP(func() {
			POST("FundsTransfer/Internal/A2A/2.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Acknowledgement Response-Synchronous: <br>0 - REQUEST ACCEPTED FOR PROCESSING")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response("forbidden", StatusForbidden)

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
			Response("conflict", StatusConflict)

		})
	})

	// 8. INSSimulation: This API is used to give instant notifications
	// or alerts on account activities(Debits,Credits) to the customer
	// so that they can reflect this in their accounting backend.
	Method("INSSimulation", func() {
		Description("Post a Debit/Credit Account Transaction Event Type Notification Simulation Request")
		Payload(INSTransactionSimulationRequest)
		Result(SuccessAcknowledgement)

		// 400
		Error("bad_request", ErrorResult, "Bad Request Error Response-Synchronous")

		// 401
		Error("unauthorized", ErrorAcknowledgement, "Unauthorized Error Response-Synchronous")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous")

		HTTP(func() {
			POST("Notifications/INS/Simulation/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Acknowledgement Response-Synchronous")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 405
			// RFC 7231, 6.5.5
			Response("not_allowed", StatusMethodNotAllowed)

			// Status 408
			// RFC 7231, 6.5.7
			Response("timeout", StatusRequestTimeout)

		})
	})

	// 9. PesaLink Send to Account Funds Transfer API will enable you
	// to transfer funds from your own Co-operative Bank account to
	// Bank account(s) in IPSL participating banks.
	Method("PesaLinkSendToAccount", func() {
		Description("Post a PesaLink Funds Transfer Send to Account Transaction")
		Payload(PesaLinkSendToAccountTransactionRequest)
		Result(SuccessAcknowledgement)

		// 400
		Error("bad_request", AcknowledgementError400, "Bad Request Error Response-Synchronous: <br>-2 - INVALID/MISSING PARAMETER, <br>-3 - THE AMOUNT DEBITED OR CREDITED MUST BE GREATER THAT 0/POSITIVE INTEGER, <br>-5 - DEBIT AND CREDIT(S) AMOUNTS NOT BALANCING, <br>-6 - TRANSACTION AMOUNT LESS/GREATER THAN MINIMUM/MAXIMUM LIMIT ALLOWED, <br>-11 - MESSAGE REFERENCE/REFERENCE NUMBER LONGER THAN ALLOWED LENGTH, <br>-12 - DUPLICATE REFERENCE/IDENTICAL REFERENCE AND MESSAGEREFERENCE")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: <br>900901 - Invalid credentials, <br>900902 - Missing credentials, <br>-8 - DEBIT ACCOUNT AUTHORIZATION FAILURE")

		// 403
		Error("forbidden", AcknowledgementError403, "Forbidden Error Response-Synchronous: <br>-9 - CURRENCY INVALID/NOT ALLOWED, <br>-10 - BANK CODE INVALID/NOT PESALINK MEMBER BANK, <br>-16 - DAILY LIMIT EXHAUSTED")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous: <br>Wrong API Resource URI")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: <br>-4 - REQUEST TIMED OUT")

		// 409
		Error("conflict", AcknowledgementError409, "Conflict Error Response-Synchronous: <br>-1 - DUPLICATE MESSAGE REFERENCE")

		HTTP(func() {
			POST("FundsTransfer/External/A2A/PesaLink/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Acknowledgement Response-Synchronous: <br>0 - REQUEST ACCEPTED FOR PROCESSING")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response("forbidden", StatusForbidden)

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
			Response("conflict", StatusConflict)

		})
	})

	// 10. PesaLink Send to Phone Funds Transfer API will enable you
	// to transfer funds from your own Co-operative Bank account to
	// a Phone Number(s) linked to a Bank account in an IPSL participating bank.
	Method("PesaLinkSendToPhone", func() {
		Description("Post a PesaLink Funds Transfer Send to Phone Transaction")
		Payload(PesaLinkSendToPhoneTransactionRequest)
		Result(SuccessAcknowledgement)

		// 400
		Error("bad_request", AcknowledgementError400, "Bad Request Error Response-Synchronous: <br>-2 - INVALID/MISSING PARAMETER, <br>-3 - THE AMOUNT DEBITED OR CREDITED MUST BE GREATER THAT 0/POSITIVE INTEGER, <br>-5 - DEBIT AND CREDIT(S) AMOUNTS NOT BALANCING, <br>-6 - TRANSACTION AMOUNT LESS/GREATER THAN MINIMUM/MAXIMUM LIMIT ALLOWED, <br>-11 - MESSAGE REFERENCE/REFERENCE NUMBER LONGER THAN ALLOWED LENGTH, <br>-12 - DUPLICATE REFERENCE/IDENTICAL REFERENCE AND MESSAGEREFERENCE, <br>-20 - INVALID PHONE NUMBER")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: <br>900901 - Invalid credentials, <br>900902 - Missing credentials, <br>-8 - DEBIT ACCOUNT AUTHORIZATION FAILURE")

		// 403
		Error("forbidden", AcknowledgementError403, "Forbidden Error Response-Synchronous: <br>-9 - CURRENCY INVALID/NOT ALLOWED, <br>-10 - BANK CODE INVALID/NOT PESALINK MEMBER BANK, <br>-16 - DAILY LIMIT EXHAUSTED")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous: <br>Wrong API Resource URI")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: <br>-4 - REQUEST TIMED OUT")

		// 409
		Error("conflict", AcknowledgementError409, "Conflict Error Response-Synchronous: <br>-1 - DUPLICATE MESSAGE REFERENCE")

		HTTP(func() {
			POST("FundsTransfer/External/A2M/PesaLink/1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Acknowledgement Response-Synchronous: <br>0 - REQUEST ACCEPTED FOR PROCESSING")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response("forbidden", StatusForbidden)

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
			Response("conflict", StatusConflict)

		})
	})

	// 11. Send to M-Pesa Funds Transfer API will enable you to transfer
	// funds from your own Co-operative Bank account to an M-Pesa account recipient.
	Method("SendToMPesa", func() {
		Description("Post a Send To M-Pesa Funds Transfer Transaction")
		Payload(SendToMpesaTransactionRequest)
		Result(SuccessAcknowledgement)

		// 400
		Error("bad_request", AcknowledgementError400, "Bad Request Error Response-Synchronous: <br>-2: INVALID/MISSING PARAMETER, <br>-3: THE AMOUNT DEBITED OR CREDITED MUST BE GREATER THAT 0/POSITIVE INTEGER, <br>-5: DEBIT AND CREDIT(S) AMOUNTS NOT BALANCING, <br>-6: TRANSACTION AMOUNT LESS/GREATER THAN MINIMUM/MAXIMUM LIMIT ALLOWED, <br>-11: MESSAGE REFERENCE/REFERENCE NUMBER LONGER THAN ALLOWED LENGTH, <br>-12: DUPLICATE REFERENCE/IDENTICAL REFERENCE AND MESSAGEREFERENCE, <br>-19: INVALID M-PESA NUMBER")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: <br>900901: Invalid credentials, <br>900902: Missing credentials, <br>-8: DEBIT ACCOUNT AUTHORIZATION FAILURE")

		// 403
		Error("forbidden", AcknowledgementError403, "Forbidden Error Response-Synchronous: <br>-9: CURRENCY INVALID/NOT ALLOWED, <br>-16: DAILY LIMIT EXHAUSTED")

		// 404
		Error("not_found", ErrorResult, "Not Found Error Response-Synchronous: <br>Wrong API Resource URI")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: <br>-4: REQUEST TIMED OUT")

		// 409
		Error("conflict", AcknowledgementError409, "Conflict Error Response-Synchronous: <br>-1: DUPLICATE MESSAGE REFERENCE")

		HTTP(func() {
			POST("FundsTransfer/External/A2M/Mpesa/v1.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Acknowledgement Response-Synchronous: <br>0: REQUEST ACCEPTED FOR PROCESSING")
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 403
			// RFC 7231, 6.5.3
			Response("forbidden", StatusForbidden)

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
			Response("conflict", StatusConflict)

		})
	})

	// 12. Transaction Status Enquiry API will enable you to enquire
	// about the status of a previously requested transaction for the
	// specified transaction message reference.
	Method("TransactionStatus", func() {
		Description("Post a Transaction Status Enquiry Request")
		Payload(FTTransactionStatusPayload)
		Result(SuccessResponse)

		// 400
		Error("bad_request", ErrorResponse, "Bad Request Error Response-Synchronous: <br>-2 - INVALID/MISSING PARAMETER, <br>-11 - MESSAGE REFERENCE LONGER THAN ALLOWED LENGTH")

		// 401
		Error("unauthorized", MissingCredentials, "Unauthorized Error Response-Synchronous: <br>900901 - Invalid credentials, <br>900902 - Missing credentials")

		// 404
		Error("not_found", NotFoundErrorResponse, "Not Found Error Response-Synchronous: <br>-13 - MESSAGE REFERENCE DOES NOT EXIST, <br>Wrong API Resource URI")

		// 405
		Error("not_allowed", ErrorResult, "Method Not Allowed Error Response-Synchronous")

		// 408
		Error("timeout", ErrorResult, "Request Timeout Error Response-Synchronous: <br>-4 - REQUEST TIMED OUT")

		HTTP(func() {
			POST("Enquiry/TransactionStatus/2.0.0")

			// Status 200
			// RFC 7231, 6.3.1
			Response(StatusOK, func() {
				Description("OK Successfully Processed Transaction Response-Synchronous: <br>0 - FULL SUCCESS, <br>Some JSON keys may not be present in the response depending on the transaction type")
			})

			// StatusMultiStatus
			// Status 207
			// RFC 4918, 11.1
			Response(StatusMultiStatus, func() {
				Description("Multi-Status Partial Success/Fully Failed Processed Transaction Response-Synchronous: <br>1 - PARTIAL SUCCESS, <br>2 - FULL FAILURE, <br>Some JSON keys may not be present in the response depending on the transaction type")
				Body(MultiStatusResponse)
			})

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

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

	Method("token", func() {
		Description("Creates a valid JWT")

		// The token endpoint is secured via basic auth
		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Username("username", String, "consumer-key for Username", func() {
				Example("user")
			})
			Password("password", String, "consumer-secret for Password", func() {
				Example("password")
			})
			Required("username", "password")
		})

		Result(Creds)

		HTTP(func() {
			POST("/token")

			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})
})
