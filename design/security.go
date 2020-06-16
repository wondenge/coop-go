package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// OAuth2Auth defines a security scheme that uses OAuth2 tokens.
var OAuth2Auth = OAuth2Security("oauth2", func() {
	Description(`Secures endpoint by requiring a valid OAuth2 token.`)
})
