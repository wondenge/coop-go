package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("connect", func() {
	Title("Co-op Bank Connect API")
	Description("HTTP Service for accessing Co-op Bank Connect API.")
	Version("1.0")
	TermsOfService("https://github.com/wondenge/coop-go/blob/master/LICENCE")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/coop-go/blob/master/LICENCE")
	})
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/coop-go/blob/master/README.md")
	})
	Server("connect", func() {
		Description("connect hosts Co-op Bank Connect API Services.")
		Services()
		Host("local", func() {
			Description("localhost")
			URI("http://localhost:3000")
		})
		Host("development", func() {
			Description("Development Hosts")
			URI("http://developer.co-opbank.co.ke:8280")
		})
		Host("production", func() {
			Description("Production Host")
			URI("https://developer.co-opbank.co.ke:8243")
		})
	})
})
