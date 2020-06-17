package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

//                                                   Enquiry API
//
// 1. Account Balance Enquiry API: Enables you to enquire about your own
// Co-operative Bank accounts' balance as at now for the specified account number.
var AccountBalancePayload = Type("AccountBalancePayload", func() {
	Description("Account Balance Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Required("MessageReference", "AccountNumber")
})

var AccountBalanceSuccessResponse = ResultType("AccountBalanceSuccessResponse", func() {
	Description("Account Balance Success Response")
	TypeName("AccountBalanceSuccessResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("MessageReference", String, func() {
			Description("Your unique transaction request message identifier")
			MinLength(1)
			MaxLength(27)
			Example("40ca18c6765086089a1")
		})
		Attribute("MessageDateTime", String, func() {
			Description("Acknowledgement message creation timestamp")
			Format(FormatDateTime)
			Example("2017-12-04T09:27:00")
		})
		Attribute("MessageCode", String, func() {
			Description("Message Response Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("Success")
		})
		Attribute("AccountNumber", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("36001873000")
		})
		Attribute("AccountName", String, func() {
			Description("Account Name")
			Example("Your Account Name")

		})
		Attribute("Currency", String, func() {
			Description("Account currency in ISO Currency Code")
			Enum("KES", "USD", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
			Example("KES")
		})
		Attribute("ProductName", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("Savings Account")
		})
		Attribute("ClearedBalance", Float64, func() {
			Description("Cleared Balance Amount")
			Example(2195.5)
		})
		Attribute("BookedBalance", Float64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("BlockedBalance", Float64, func() {
			Description("Cleared Balance Amount")
			Example(1760.0)
		})
		Attribute("AvailableBalance", Float64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("ArrearsAmount", Float64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("BranchName", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("UKULIMA BRANCH")
		})
		Attribute("BranchSortCode", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("00011011")
		})
		Attribute("AverageBalance", Float64, func() {
			Description("Cleared Balance Amount")
			Example(75.83)
		})
		Attribute("UnclearedBalance", Float64, func() {
			Description("Cleared Balance Amount")
			Example(-2195.5)
		})
		Attribute("ODLimit", Float64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("CreditLimit", Float64, func() {
			Description("Cleared Balance Amount")
			Example(1000.0)
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})

	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("AccountNumber")
		Attribute("AccountName")
		Attribute("Currency")
		Attribute("ProductName")
		Attribute("ClearedBalance")
		Attribute("BookedBalance")
		Attribute("BlockedBalance")
		Attribute("AvailableBalance")
		Attribute("ArrearsAmount")
		Attribute("BranchName")
		Attribute("BranchSortCode")
		Attribute("AverageBalance")
		Attribute("UnclearedBalance")
		Attribute("ODLimit")
		Attribute("CreditLimit")
	})
})

// 2. AccountFullStatement
// 2. Account Full Statement Enquiry API:
// Enables you to enquire about your own Co-operative Bank accounts'
// full statement for the specified account number, start date and end date"
var AccountFullStatementPayload = Type("AccountFullStatementPayload", func() {
	Description("Account Full Statement Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Attribute("StartDate", String, func() {
		Description("Statement Start Date")
		Format(FormatDate)
		Example("2019-01-01")
	})
	Attribute("EndDate", String, func() {
		Description("Statement End Date")
		Format(FormatDate)
		Example("2019-07-01")
	})
	Required("MessageReference", "AccountNumber", "StartDate", "EndDate")
})

var AccountFullStatementSuccessResponse = ResultType("AccountFullStatementSuccessResponse", func() {
	Description("Account Full Statement Success Response")
	TypeName("AccountFullStatementSuccessResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("MessageReference", String, func() {
			Description("Your unique transaction request message identifier")
			MinLength(1)
			MaxLength(27)
			Example("40ca18c6765086089a1")
		})
		Attribute("MessageDateTime", String, func() {
			Description("Acknowledgement message creation timestamp")
			Format(FormatDateTime)
			Example("2017-12-04T09:27:00")
		})
		Attribute("MessageCode", String, func() {
			Description("Message Response Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Acknowledgement/Response Message Code Description")
			Example("Success")
		})
		Attribute("AccountNumber", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("36001873000")
		})
		Attribute("AccountName", String, func() {
			Description("Account Name")
			Example("Your Account Name")
		})
		Attribute("Transactions", ArrayOf(AccountTransaction), func() {
			MinLength(0)
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("AccountNumber")
		Attribute("AccountName")
		Attribute("Transactions")
	})
})

// 3. AccountMiniStatement
var AccountMiniStatementPayload = Type("AccountMiniStatementPayload", func() {
	Description("Account Mini Statement Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Required("MessageReference", "AccountNumber")
})

var AccountMiniStatementSuccessResponse = ResultType("AccountMiniStatementSuccessResponse", func() {
	Description("Account Mini Statement Success Response")
	TypeName("AccountMiniStatementSuccessResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("MessageReference", String, func() {
			Description("Your unique transaction request message identifier")
			MinLength(1)
			MaxLength(27)
			Example("40ca18c6765086089a1")
		})
		Attribute("MessageDateTime", String, func() {
			Description("Acknowledgement message creation timestamp")
			Format(FormatDateTime)
			Example("2017-12-04T09:27:00")
		})
		Attribute("MessageCode", String, func() {
			Description("Message Response Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Acknowledgement/Response Message Code Description")
			Example("Success")
		})
		Attribute("AccountNumber", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("36001873000")
		})
		Attribute("AccountName", String, func() {
			Description("Account Name")
			Example("Your Account Name")
		})
		Attribute("Transactions", ArrayOf(AccountTransaction), func() {
			MinLength(0)
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("AccountNumber")
		Attribute("AccountName")
		Attribute("Transactions")
	})
})

// 4. AccountTransactions

// 4. Account Transactions Enquiry API:
// Enables you to enquire about your own Co-operative Bank accounts'
// latest transactions for the specified account number and number
// of transactions to be returned
var AccountTransactionsPayload = Type("AccountTransactionsPayload", func() {
	Description("Account Transactions Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Attribute("NoOfTransactions", Int, func() {
		Description("No Of Latest Transactions To Be Returned")
		Example(1)
	})
	Required("MessageReference", "AccountNumber", "NoOfTransactions")
})

var AccountTransactionsSuccessResponse = ResultType("AccountTransactionsSuccessResponse", func() {
	Description("Account Transactions Success Response")
	TypeName("AccountTransactionsSuccessResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("MessageReference", String, func() {
			Description("Your unique transaction request message identifier")
			MinLength(1)
			MaxLength(27)
			Example("40ca18c6765086089a1")
		})
		Attribute("MessageDateTime", String, func() {
			Description("Acknowledgement message creation timestamp")
			Format(FormatDateTime)
			Example("2017-12-04T09:27:00")
		})
		Attribute("MessageCode", String, func() {
			Description("Message Response Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Acknowledgement/Response Message Code Description")
			Example("Success")
		})
		Attribute("AccountNumber", String, func() {
			Description("Posting account number")
			MinLength(14)
			MaxLength(14)
			Example("36001873000")
		})
		Attribute("AccountName", String, func() {
			Description("Account Name")
			Example("Your Account Name")
		})
		Attribute("NoOfTransactions", Int, func() {
			Description("No Of Latest Transactions To Be Returned")
			Example(1)
		})
		Attribute("TotalCredits", Float64, func() {
			Description("Total Credits Amount")
			Example(200.0)
		})
		Attribute("TotalDebits", Float64, func() {
			Description("Total Debits Amount")
			Example(0.0)
		})
		Attribute("Transactions", ArrayOf(AccountTransaction), func() {
			MinLength(0)
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("AccountNumber")
		Attribute("AccountName")
		Attribute("NoOfTransactions")
		Attribute("TotalCredits")
		Attribute("TotalDebits")
		Attribute("Transactions")
	})
})

// 5. AccountValidation
// 5. Account Validation Enquiry API:
// Enables you to validate a Co-operative Bank account number.
var AccountValidationPayload = Type("AccountValidationPayload", func() {
	Description("Account Validation Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Required("MessageReference", "AccountNumber")
})

var AccountValidationSuccessResponse = ResultType("AccountValidationSuccessResponse", func() {
	Description("Account Validation Success Response")
	TypeName("AccountValidationSuccessResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("MessageReference", String, func() {
			Description("Your unique transaction request message identifier")
			MinLength(1)
			MaxLength(27)
			Example("40ca18c6765086089a1")
		})
		Attribute("MessageDateTime", String, func() {
			Description("Acknowledgement message creation timestamp")
			Format(FormatDateTime)
			Example("2017-12-04T09:27:00")
		})
		Attribute("MessageCode", String, func() {
			Description("Message Response Code")
			Example("-2")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("VALID ACCOUNT NUMBER")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
	})
})

// 6. ExchangeRate
// Exchange Rate Enquiry API will enable you to enquire about the
// current SPOT exchange rate for the specified currencies
var ExchangeRatePayload = Type("ExchangeRatePayload", func() {
	Description("Exchange Rate Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("FromCurrencyCode", String, func() {
		Description("From Currency Code in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("KES")
	})
	Attribute("ToCurrencyCode", String, func() {
		Description("To Currency Code in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("USD")
	})
	Required("MessageReference", "FromCurrencyCode", "ToCurrencyCode")
})

var ExchangeRateSuccessResponse = ResultType("ExchangeRateSuccessResponse", func() {
	Description("Exchange Rate Success Response")
	TypeName("ExchangeRateSuccessResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("MessageReference", String, func() {
			Description("Your unique transaction request message identifier")
			MinLength(1)
			MaxLength(27)
			Example("40ca18c6765086089a1")
		})
		Attribute("MessageDateTime", String, func() {
			Description("Acknowledgement message creation timestamp")
			Format(FormatDateTime)
			Example("2017-12-04T09:27:00")
		})
		Attribute("MessageCode", String, func() {
			Description("Message Response Code")
			Example("-2")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("VALID ACCOUNT NUMBER")
		})
		Attribute("FromCurrencyCode", String, func() {
			Description("From Currency Code in ISO Currency Code")
			Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
			Example("KES")
		})
		Attribute("ToCurrencyCode", String, func() {
			Description("To Currency Code in ISO Currency Code")
			Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
			Example("USD")
		})
		Attribute("RateType", String, func() {
			Description("Exchange rate type")
			Example("SPOT")
		})
		Attribute("Rate", Float64, func() {
			Description("Exchange rate amount")
			Example(103.5)
		})
		Attribute("Tolerance", Float64, func() {
			Description("Exchange rate tolerance")
			Example(15.0)
		})
		Attribute("MultiplyDivide", func() {
			Description("Exchange rate is a multiply or a divide")
			Enum("M", "D")
			Example("D")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("FromCurrencyCode")
		Attribute("ToCurrencyCode")
		Attribute("RateType")
		Attribute("Rate")
		Attribute("Tolerance")
		Attribute("MultiplyDivide")
	})
})

// 12. TransactionStatus
// TransactionStatus
// Transaction Status Enquiry API will enable you to enquire about
// the status of a previously requested transaction for the specified
// transaction message reference.
var FTTransactionStatusPayload = Type("FTTransactionStatusPayload", func() {
	Description("FT Transaction Status Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})

	Required("MessageReference")
})

var AccountTransaction = Type("AccountTransaction", func() {
	Description("Account Transaction")

	Attribute("TransactionId", String, func() {
		Description("Unique Transaction Posting Identifier")
		Example("116bdbebcca41aXF")
	})
	Attribute("TransactionDate", String, func() {
		Description("Transaction Date")
		Format(FormatDateTime)
		Example("2019-04-29T10:05:41.178+03:00")
	})
	Attribute("ValueDate", String, func() {
		Description("Transaction Value Date")
		Format(FormatDateTime)
		Example("2019-04-29T10:05:40.751+03:00")
	})
	Attribute("Narration", String, func() {
		Description("Transaction Narration")
		Example("Electricity Payment")
	})
	Attribute("TransactionType", String, func() {
		Description("Transaction Type, C for Credit or D for Debit")
		Enum("C", "D")
		Example("C")
	})
	Attribute("ServicePoint", String, func() {
		Description("Transaction Service Point")
		Example("ATM")
	})
	Attribute("TransactionReference", String, func() {
		Description("Unique Transaction Reference")
		Example("911909902484902484")
	})
	Attribute("CreditAmount", Float64, func() {
		Description("Transaction Credit Amount")
		Example(200.0)
	})
	Attribute("DebitAmount", Float64, func() {
		Description("Transaction Debit Amount")
		Example(0.0)
	})
	Attribute("RunningClearedBalance", Float64, func() {
		Description("Account Running Cleared Balance")
		Example(1215.7)
	})
	Attribute("RunningBookBalance", Float64, func() {
		Description("Account Running Book Balance")
		Example(1215.7)
	})
	Attribute("DebitLimit", Float64, func() {
		Description("Account Debit Limit")
		Example(0.0)
	})
	Attribute("LimitExpiryDate", String, func() {
		Description("Account Debit Limit Expiry Date")
		Example("2019-04-29T10:05:41.178+03:00")
	})
})
