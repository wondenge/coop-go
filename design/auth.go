package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Authentication
// An Access Token will be used to access, authorize and authenticate
// your interactions with the Co-op Bank’s Open Banking APIs.
// It will be generated using a Consumer Key and a Consumer Secret linked
// to your account.
// The Access Token will expire after a specified duration of time, after
// which regeneration of the Access Token will be required.
// To generate or re-generate an Access Token, your sign-in username and
// password will be required. It is very important to keep your
// API credentials safe, as they can be used to access your account, make
// changes to your account and carry out transactions. It is also important
// to note that once you generate a new set of API credentials you must
// update all your applications with the new API details for your
// applications to continue working. Note that the Sandbox and Live
// environments have unique Consumer Key and Consumer Secret which
// are NOT be interchangeable.

// The following headers will be required for each request:
// AuthenticationToken(String & Required)
// This is your auth key from the platform

// BasicAuth defines a security scheme using basic authentication.
// The scheme protects the "token" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during token")
	Scope("api:read", "Read-only access")
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = Type("Creds", func() {
	Attribute("jwt", String, "JWT token", func() {
		Example("a833ce4a-5903-3c5c-a9e9-7ae042e3080b")
	})
	Required("jwt")
})
