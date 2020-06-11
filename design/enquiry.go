package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// TransactionStatus
// Transaction Status Enquiry API will enable you to enquire about
// the status of a previously requested transaction for the specified
// transaction message reference.
var FTTransactionStatusRequest = Type("FTTransactionStatusRequest", func() {
	Description("FT Transaction Status Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})

	Required("MessageReference")
})
