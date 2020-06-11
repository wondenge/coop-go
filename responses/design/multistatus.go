package design

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


