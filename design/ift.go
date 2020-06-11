package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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
	Attribute("Source", SourceAccountTXNRequest)
	Attribute("Destinations", ArrayOf(DestinationAccountTXNRequest))
	Required("MessageReference", "CallBackUrl", "Source", "Destinations")
})

var IFTAccountToAccountCallbackRequest = Type("IFTAccountToAccountCallbackRequest", func() {
	Description("IFT Account To Account Callback Request")

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
