package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// 1. Account Balance Enquiry API:
// Enables you to enquire about your own Co-operative Bank accounts'
// balance as at now for the specified account number.
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

// 2. Account Full Statement Enquiry API:
// Enables you to enquire about your own Co-operative Bank accounts'
// full statement for the specified account number, start date and end date"
var AccountFullStatementRequest = Type("AccountFullStatementRequest", func() {
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

var AccountMiniStatementRequest = Type("AccountMiniStatementRequest", func() {
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

// 4. Account Transactions Enquiry API:
// Enables you to enquire about your own Co-operative Bank accounts'
// latest transactions for the specified account number and number
// of transactions to be returned
var AccountTransactionsRequest = Type("AccountTransactionsRequest", func() {
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
		Attribute("TotalCredits", UInt64, func() {
			Description("Total Credits Amount")
			Example(200.0)
		})
		Attribute("TotalDebits", UInt64, func() {
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

// 5. Account Validation Enquiry API:
// Enables you to validate a Co-operative Bank account number.
var AccountValidationRequest = Type("AccountValidationRequest", func() {
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
	Attribute("CreditAmount", UInt64, func() {
		Description("Transaction Credit Amount")
		Example(200.0)
	})
	Attribute("DebitAmount", UInt64, func() {
		Description("Transaction Debit Amount")
		Example(0.0)
	})
	Attribute("RunningClearedBalance", UInt64, func() {
		Description("Account Running Cleared Balance")
		Example(1215.7)
	})
	Attribute("RunningBookBalance", UInt64, func() {
		Description("Account Running Book Balance")
		Example(1215.7)
	})
	Attribute("DebitLimit", UInt64, func() {
		Description("Account Debit Limit")
		Example(0.0)
	})
	Attribute("LimitExpiryDate", String, func() {
		Description("Account Debit Limit Expiry Date")
		Example("2019-04-29T10:05:41.178+03:00")
	})
})
