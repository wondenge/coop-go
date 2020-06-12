package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var ErrorResponse = ResultType("ErrorResponse", func() {
	Description("Error Response")
	TypeName("ErrorResponse")
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
			Example("INVALID/MISSING PARAMETER")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})

	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
	})
})



var ErrorAcknowledgement = ResultType("ErrorAcknowledgement", func() {
	Description("Unauthorized Error Response")
	TypeName("ErrorAcknowledgement")
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
			Example("INVALID/MISSING PARAMETER")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})

	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
	})
})

var AcknowledgementError400 = ResultType("AcknowledgementError400", func() {
	Description("Bad Request Error Response")
	TypeName("AcknowledgementError400")
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
			Example("INVALID/MISSING PARAMETER")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})

	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
	})
})

var AcknowledgementError403 = ResultType("AcknowledgementError403", func() {
	Description("Forbidden Error Response")
	TypeName("AcknowledgementError403")
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
			Example("-9")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("CURRENCY INVALID/NOT ALLOWED")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})

	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
	})
})

var AcknowledgementError409 = ResultType("AcknowledgementError409", func() {
	Description("Conflict Error Response")
	TypeName("AcknowledgementError409")
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
			Example("-1")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("DUPLICATE MESSAGE REFERENCE")
		})
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription")
	})

	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
	})
})
