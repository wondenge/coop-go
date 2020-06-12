package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var DestinationAccountCallbackRequest = Type("DestinationAccountCallbackRequest", func() {
	Description("Destination Account Callback Request")

	Attribute("ReferenceNumber", String, func() {
		Description("Unique posting reference for the transaction leg")
		MinLength(1)
		MaxLength(30)
		Example("40ca18c6765086089a1_1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
		Example("54321987654321")
	})
	Attribute("Amount", UInt64, func() {
		Description("Transaction Amount")
		Minimum(0.01)
		Maximum(999999.99)
		Example(777)
	})
	Attribute("TransactionCurrency", String, func() {
		Description("Posting account currency in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("KES")
	})
	Attribute("Narration", String, func() {
		Description("Posting account transaction narration")
		MinLength(1)
		MaxLength(25)
		Example("Electricity Payment")
	})
	Attribute("TransactionID", String, func() {
		Description("Co-operative Bank's processed transaction number")
		Example("1169716b65891lI6")
	})
	Attribute("ResponseCode", String, func() {
		Description("Posting leg response code")
		Example("0")
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})
	Required("ReferenceNumber", "AccountNumber", "Amount", "TransactionCurrency", "Narration", "ResponseCode", "ResponseDescription")
})

var DestinationAccount = Type("DestinationAccount", func() {
	Description("Destination Account")

	Attribute("ReferenceNumber", String, func() {
		Description("Unique posting reference for the transaction leg")
		MinLength(1)
		MaxLength(30)
		Example("40ca18c6765086089a1_1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("12345678912345")
	})
	Attribute("MobileNumber", String, func() {
		Description("Recipient phone number linked to a bank account in an IPSL participating bank")
		Enum("07xxxxxxxx", "2547xxxxxxxx", "+2547xxxxxxxx")
		MinLength(10)
		MaxLength(13)
	})
	Attribute("PhoneNumber", String, func() {
		Description("Recipient phone number linked to a bank account in an IPSL participating bank")
		Enum("07xxxxxxxx", "2547xxxxxxxx", "+2547xxxxxxxx")
		MinLength(10)
		MaxLength(13)
	})
	Attribute("BankCode", String, func() {
		Description("Posting account bank code")
		Example("011")
	})
	Attribute("Amount", UInt64, func() {
		Description("Transaction Amount")
	})
	Attribute("TransactionCurrency", String, func() {
		Description("Posting account currency in ISO Currency Code")
		Enum("KES", "USD", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("KES")
	})
	Attribute("Narration", String, func() {
		Description("Posting account transaction narration")
		Example("Supplier Payment")
	})
	Attribute("TransactionID", String, func() {
		Description("Co-operative Bank's processed transaction number")
		Example("1169716b65891lI6")
	})
	Attribute("ResponseCode", String, func() {
		Description("Posting leg response code")
		Example("0")
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})
	Required("ReferenceNumber", "AccountNumber", "Amount", "TransactionCurrency", "Narration")
})

var DestinationAccountTransactionRequest = Type("DestinationAccountTransactionRequest", func() {
	Description("Destination Account Transaction Request")
	Extend(DestinationAccountTXNRequest)
	Required("ReferenceNumber", "AccountNumber", "Amount", "TransactionCurrency", "Narration")
})

var DestinationAccountTXNRequest = Type("DestinationAccountTXNRequest", func() {
	Description("Destination Account Transaction Request")

	Attribute("ReferenceNumber", String, func() {
		Description("Unique posting reference for the transaction leg")
		MinLength(1)
		MaxLength(30)
		Example("40ca18c6765086089a1_1")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
		Example("54321987654321")
	})
	Attribute("Amount", UInt64, func() {
		Description("Transaction Amount")
		Minimum(0.01)
		Maximum(999999.99)
		Example(777)
	})
	Attribute("TransactionCurrency", String, func() {
		Description("Posting account currency in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("KES")
	})
	Attribute("Narration", String, func() {
		Description("Posting account transaction narration")
		MinLength(1)
		MaxLength(25)
		Example("Supplier Payment")
		Example("Electricity Payment")
	})
	Required("ReferenceNumber", "AccountNumber", "Amount", "TransactionCurrency", "Narration")
})
