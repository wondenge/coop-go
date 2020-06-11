package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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
