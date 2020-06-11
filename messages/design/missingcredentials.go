package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var MissingCredentials = ResultType("MissingCredentials", func() {
	Description("Missing Credentials")

	Attribute("Fault", func() {
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
})
