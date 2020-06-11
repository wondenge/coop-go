package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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
