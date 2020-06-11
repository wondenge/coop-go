package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var SourceAccountCallbackRequest = Type("SourceAccountCallbackRequest", func() {
	Description("Source Account Callback Request")

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
	Attribute("ResponseCode", String, func() {
		Description("Posting leg response code")
		Example("0")
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})
	Required("AccountNumber", "Amount", "TransactionCurrency", "Narration", "ResponseCode", "ResponseDescription")
})

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

var PesaLinkSendToPhoneCallbackRequest = Type("PesaLinkSendToPhoneCallbackRequest", func() {
	Description("Pesalink send to phone callback request")

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
		Example("FULL SUCCESS")
	})
	Attribute("Source", SourceAccountCallbackRequest)
	Attribute("Destinations", ArrayOf(DestinationAccountCallbackRequest), func() {
		MinLength(1)
	})
	Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription", "Source", "Destinations")
})

var PesaLinkSendToAccountCallbackRequest = Type("PesaLinkSendToAccountCallbackRequest", func() {
	Description("Pesalink send to account callback request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("CallBackUrl", String, func() {
		Description("Your callback URL that will receive transaction processing results")
		Example("https://yourdomain.com/ft-callback")
	})
	Attribute("Source", SourceAccountTransactionRequest)
	Attribute("Destinations", ArrayOf(DestinationAccountTransactionRequest), func() {
		MinLength(1)
	})
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

var SendToMpesaCallbackRequest = Type("SendToMpesaCallbackRequest", func() {
	Description("Send to Mpesa callback request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("CallBackUrl", String, func() {
		Description("Your callback URL that will receive transaction processing results")
		Example("https://yourdomain.com/ft-callback")
	})
	Attribute("Source", SourceAccountTransactionRequest)
	Attribute("Destinations", ArrayOf(DestinationAccountTransactionRequest), func() {
		MinLength(1)
	})
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})
