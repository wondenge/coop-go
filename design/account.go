package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Account Balance Enquiry API will enable you to enquire about your own
// Co-operative Bank accounts' balance as at now for the specified account number.
var AccountBalanceRequest = Type("AccountBalanceRequest", func() {
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
		Attribute("ClearedBalance", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(2195.5)
		})
		Attribute("BookedBalance", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("BlockedBalance", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(1760.0)
		})
		Attribute("AvailableBalance", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("ArrearsAmount", UInt64, func() {
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
		Attribute("AverageBalance", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(75.83)
		})
		Attribute("UnclearedBalance", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(-2195.5)
		})
		Attribute("ODLimit", UInt64, func() {
			Description("Cleared Balance Amount")
			Example(0.0)
		})
		Attribute("CreditLimit", UInt64, func() {
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
