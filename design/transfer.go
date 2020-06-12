package design

import (
	design2 "github.com/wondenge/coop-go/responses/design"
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
	Attribute("Source", design2.SourceAccountTXNRequest)
	Attribute("Destinations", ArrayOf(design2.DestinationAccountTXNRequest))
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

// ExternalFundsTransfer
// 9. PesaLinkSendToAccount
var PesaLinkSendToAccountTransactionRequest = Type("PesaLinkSendToAccountTransactionRequest", func() {
	Description("Pesalink send to account transaction request")

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
	Attribute("Source", design2.SourceAccountTransactionRequest)
	Attribute("Destinations", ArrayOf(design2.DestinationAccountTransactionRequest), func() {
		MinLength(1)
	})
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

// 10.PesaLinkSendToPhone
var PesaLinkSendToPhoneTransactionRequest = Type("PesaLinkSendToPhoneTransactionRequest", func() {
	Description("Pesalink send to phone transaction request")

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
	Attribute("Source", design2.SourceAccountTransactionRequest)
	Attribute("Destinations", ArrayOf(design2.DestinationAccountTransactionRequest), func() {
		MinLength(1)
	})
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

// 11.SendToMPesa
var SendToMpesaTransactionRequest = Type("SendToMpesaTransactionRequest", func() {
	Description("Send to Mpesa transaction request")

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
	Attribute("Source", design2.SourceAccountTransactionRequest)
	Attribute("Destinations", ArrayOf(design2.DestinationAccountTransactionRequest), func() {
		MinLength(1)
	})
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})
