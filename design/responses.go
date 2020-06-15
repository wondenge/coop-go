package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Status 102
var ProcessingResponse = ResultType("ProcessingResponse", func() {
	Description("Processing In Progress Response")
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
			Example("3")
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

// Status 200
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

// Status 200
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

// Status 207
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

var MissingCredentials = ResultType("MissingCredentials", func() {
	Description("Missing Credentials")
	TypeName("MissingCredentials")
	ContentType("application/json")

	Attributes(func() {
		Attribute("Fault", MissingCredentialFault)
	})
	View("default", func() {
		Attribute("Fault")
	})
})

var MissingCredentialFault = Type("MissingCredentialFault", func() {
	Attribute("code", String, func() {
		Example("900902")
	})
	Attribute("message", String, func() {
		Example("Missing Credentials")
	})
	Attribute("description", String, func() {

		// Make sure your API invocation call has a
		// header: "Authorization: Bearer ACCESS_TOKEN"
		Example("Required OAuth credentials not provided.")
	})
})

var NotFoundErrorResponse = ResultType("NotFoundErrorResponse", func() {
	Description("Not Found Error Response")
	TypeName("NotFoundErrorResponse")
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
			Example("-13")
		})
		Attribute("MessageDescription", String, func() {
			Description("Transaction request message code description")
			Example("MESSAGE REFERENCE DOES NOT EXIST")
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
