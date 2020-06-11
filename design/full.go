package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Account Full Statement Enquiry API will enable you to enquire about
// your own Co-operative Bank accounts' full statement for the specified
// account number, start date and end date"
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
