package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var SuccessAcknowledgement = ResultType("SuccessAcknowledgement", func() {
	Description("Success Acknowledgement")
	TypeName("SuccessAcknowledgement")
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
			Description("Acknowledgement/Response Message Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("FULL SUCCESS")
			Example("REQUEST ACCEPTED FOR PROCESSING")
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

var SuccessResponse = ResultType("SuccessResponse", func() {
	Description("Success Response")
	TypeName("SuccessResponse")
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
			Description("Acknowledgement/Response Message Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("FULL SUCCESS")
		})
		Attribute("Source", SourceAccount)
		Attribute("Destinations", ArrayOf(DestinationAccount))
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription", "Source", "Destinations")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("Source")
		Attribute("Destinations")
	})
})

var MultiStatusResponse = ResultType("MultiStatusResponse", func() {
	Description("MultiStatus Response")
	TypeName("MultiStatusResponse")
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
			Description("Acknowledgement/Response Message Code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Message Code description")
			Example("FULL SUCCESS")
		})
		Attribute("Source", SourceAccount)
		Attribute("Destinations", ArrayOf(DestinationAccount))
		Required("MessageReference", "MessageDateTime", "MessageCode", "MessageDescription", "Source", "Destinations")
	})
	View("default", func() {
		Attribute("MessageReference")
		Attribute("MessageDateTime")
		Attribute("MessageCode")
		Attribute("MessageDescription")
		Attribute("Source")
		Attribute("Destinations")
	})
})

var ProcessingResponse = ResultType("ProcessingResponse", func() {
	Description("Processing Response")
	TypeName("ProcessingResponse")
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
			Description("Transaction request message code")
			Example("0")
		})
		Attribute("MessageDescription", String, func() {
			Description("Transaction request message code description")
			Example("FULL SUCCESS")
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

var NotFound = ResultType("NotFound", func() {
	Description("Not Found")
	TypeName("NotFound")
	ContentType("application/json")

	Attributes(func() {
		Attribute("timestamp", String, func() {
			Example("2019-03-28T12:42:55.260+0000")
		})
		Attribute("status", String, func() {
			Example("404")
		})
		Attribute("error", String, func() {
			Example("Not Found")
		})
		Attribute("message", String, func() {
			Example("No Message Available")
		})
		Attribute("path", String, func() {
			Example("/*")
		})
	})
})
