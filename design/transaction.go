package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Account Transactions Enquiry API will enable you to enquire about your
// own Co-operative Bank accounts' latest transactions for the specified
// account number and number of transactions to be returned
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


