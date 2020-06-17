package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Funds Transfer API

// InternalFundsTransfer
// 7. IFTAccountToAccount

// Internal Funds Transfer Account to Account API will enable you
// to transfer funds from your own Co-operative Bank account
// to other Co-operative Bank account(s).
var IFTAccountToAccountTXNRequest = Type("IFTAccountToAccountTXNRequest", func() {
	Description("IFT Account To Account Transaction Request")

	AccessToken("access_token", String, func() {
		Example("1c9f6c4b-625c-3255-ba1a-026df12ab648")
	})
	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("CallBackUrl", String, func() {
		Description("Your callback URL that will receive transaction processing results")
		Example("https://yourdomain.com/ftresponse")
	})
	Attribute("Source", SourceAccountTXNRequest)
	Attribute("Destinations", DestinationsTXNRequest)
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

// ExternalFundsTransfer
// 9. PesaLinkSendToAccount
var PesaLinkSendToAccountTransactionRequest = Type("PesaLinkSendToAccountTransactionRequest", func() {
	Description("Pesalink send to account transaction request")

	AccessToken("access_token", String, func() {
		Example("1c9f6c4b-625c-3255-ba1a-026df12ab648")
	})
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
	Attribute("Destinations", DestinationsTransactionRequest)
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

// 10.PesaLinkSendToPhone
var PesaLinkSendToPhoneTransactionRequest = Type("PesaLinkSendToPhoneTransactionRequest", func() {
	Description("Pesalink send to phone transaction request")

	AccessToken("access_token", String, func() {
		Example("1c9f6c4b-625c-3255-ba1a-026df12ab648")
	})
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
	Attribute("Destinations", DestinationsTransactionRequest)
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

// 11.SendToMPesa
var SendToMpesaTransactionRequest = Type("SendToMpesaTransactionRequest", func() {
	Description("Send to Mpesa transaction request")

	AccessToken("access_token", String, func() {
		Example("1c9f6c4b-625c-3255-ba1a-026df12ab648")
	})
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
	Attribute("Destinations", DestinationsTransactionRequest)
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

var SourceAccount = Type("SourceAccount", func() {
	Description("Source Account")

	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("12345678912345")
	})
	Attribute("Amount", Float64, func() {
		Description("Transaction Amount")
		Example(0)
	})
	Attribute("TransactionCurrency")
	Attribute("Narration", String, func() {
		Description("Posting account transaction narration")
		Example("Supplier Payment")
	})
	Attribute("ResponseCode", String, func() {
		Description("Posting leg response code")
		Example("0")
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})
	Required("AccountNumber", "Amount", "TransactionCurrency", "Narration")
})

var SourceAccountTransactionRequest = Type("SourceAccountTransactionRequest", func() {
	Description("Source Account Transaction Request")

	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Attribute("Amount", Float64, func() {
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
	})
	Attribute("ResponseDescription", String, func() {
		Description("Posting leg response description")
		Example("Success")
	})
	Required("AccountNumber", "Amount", "TransactionCurrency", "Narration")
})

var SourceAccountTXNRequest = Type("SourceAccountTXNRequest", func() {
	Description("Source Account Transaction Request")

	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("36001873000")
	})
	Attribute("Amount", Float64, func() {
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
	})
	Required("AccountNumber", "Amount", "TransactionCurrency", "Narration")
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
	Attribute("Amount", Float64, func() {
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

var DestinationsTransactionRequest = Type("DestinationsTransactionRequest", func() {
	Description("Destinations Transaction Request")

	Attribute("DestinationAccountTransactionRequest", ArrayOf(DestinationAccountTransactionRequest), func() {
		MinLength(1)
	})
})

var DestinationAccountTransactionRequest = Type("DestinationAccountTransactionRequest", func() {
	Description("Destination Account Transaction Request")

	Attribute("ReferenceNumber", String, func() {
		Description("Unique posting reference for the transaction leg")
		MinLength(1)
		MaxLength(30)
		Example("40ca18c6765086089a1_1")
	})
	Attribute("MobileNumber", String, func() {

		// "07xxxxxxxx", "2547xxxxxxxx", "+2547xxxxxxxx"
		Description("Recipient phone number linked to a bank account in an IPSL participating bank")
		MinLength(10)
		MaxLength(13)
		Example("07xxxxxxxx")
	})
	Attribute("Amount", Float64, func() {
		Description("Transaction Amount")
		Minimum(0.01)
		Maximum(999999.99)
		Example(777)
	})
	Attribute("Narration", String, func() {
		Description("Posting account transaction narration")
		MinLength(1)
		MaxLength(25)
		Example("Stationary Payment")
	})
	Required("ReferenceNumber", "MobileNumber", "Amount", "Narration")
})

var DestinationsTXNRequest = Type("DestinationsTXNRequest", func() {
	Description("Destinations TXN Request")

	Attribute("DestinationsTXNRequest", ArrayOf(DestinationAccountTXNRequest), func() {
		MinLength(1)
	})
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
	Attribute("Amount", Float64, func() {
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
	})
	Required("ReferenceNumber", "AccountNumber", "Amount", "TransactionCurrency", "Narration")
})
