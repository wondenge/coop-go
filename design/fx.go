package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Exchange Rate Enquiry API will enable you to enquire about the
// current SPOT exchange rate for the specified currencies
var ExchangeRateRequest = Type("ExchangeRateRequest", func() {
	Description("Exchange Rate Request")

	Attribute("MessageReference", String, func() {
		Description("Your unique transaction request message identifier")
		MinLength(1)
		MaxLength(27)
		Example("40ca18c6765086089a1")
	})
	Attribute("FromCurrencyCode", String, func() {
		Description("From Currency Code in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("KES")
	})
	Attribute("ToCurrencyCode", String, func() {
		Description("To Currency Code in ISO Currency Code")
		Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
		Example("USD")
	})
	Required("MessageReference", "FromCurrencyCode", "ToCurrencyCode")
})

var ExchangeRateSuccessResponse = ResultType("ExchangeRateSuccessResponse", func() {
	Description("Exchange Rate Success Response")
	TypeName("ExchangeRateSuccessResponse")
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
			Example("-2")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("VALID ACCOUNT NUMBER")
		})
		Attribute("FromCurrencyCode", String, func() {
			Description("From Currency Code in ISO Currency Code")
			Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
			Example("KES")
		})
		Attribute("ToCurrencyCode", String, func() {
			Description("To Currency Code in ISO Currency Code")
			Enum("USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
			Example("USD")
		})
		Attribute("RateType", String, func() {
			Description("Exchange rate type")
			Example("SPOT")
		})
		Attribute("Rate", UInt64, func() {
			Description("Exchange rate amount")
			Example(103.5)
		})
		Attribute("Tolerance", UInt64, func() {
			Description("Exchange rate tolerance")
			Example(15.0)
		})
		Attribute("MultiplyDivide", func() {
			Description("Exchange rate is a multiply or a divide")
			Enum("M", "D")
			Example("D")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("FromCurrencyCode")
		Attribute("ToCurrencyCode")
		Attribute("RateType")
		Attribute("Rate")
		Attribute("Tolerance")
		Attribute("MultiplyDivide")
	})
})
