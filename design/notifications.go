package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Notifications API
// 8. INSSimulation
// This API is used to give instant notifications or alerts on
// account activities(Debits,Credits) to the customer so that
// they can reflect this in their accounting backend.

var INSTransactionSimulationRequest = Type("INSTransactionSimulationRequest", func() {
	Description("INS Transaction Simulation Request")

	AccessToken("access_token", String, func() {
		Example("1c9f6c4b-625c-3255-ba1a-026df12ab648")
	})
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
	Attribute("ServiceName", String, func() {
		Description("Notification Service Name identifier")
	})
	Attribute("NotificationCode", String, func() {
		Description("Notification Code identifier")
	})
	Attribute("PaymentRef", String, func() {
		Description("Transaction Payment Reference")
		Example("SFI427E9136D7D3F21C2C89")
	})
	Attribute("AccountNumber", String, func() {
		Description("Posting account number")
		MinLength(14)
		MaxLength(14)
		Example("54321987654321")
	})
	Attribute("Amount", Float64, func() {
		Description("Transaction Amount")
		Example(120777.45)
	})
	Attribute("TransactionDate", String, func() {
		Description("Posting date of the Transaction")
		Format(FormatDateTime)
		Example("20190301165420")
	})
	Attribute("EventType", String, func() {
		Description("The event of the transaction")
		Example("DEBIT")
	})
	Attribute("Currency", String, func() {
		Description("Transaction Posting account currency in ISO Currency Code")
		Enum("KES", "USD", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR")
	})
	Attribute("ExchangeRate", UInt64, func() {
		Description("Exchange Rate")
		Example(1)
	})
	Attribute("Narration", String, func() {
		Description("Transaction Posting account narration")
		Example("Supplier Payments")
	})
	Attribute("CustMemo", CustMemo)
	Attribute("ValueDate", String, func() {
		Description("Transaction Posting Value Date")
		Format(FormatDate)
		Example("20190301")
	})
	Attribute("EntryDate", String, func() {
		Description("Transaction Posting Entry Date")
		Format(FormatDate)
		Example("20190301")
	})
	Attribute("TransactionId", String, func() {
		Description("Co-operative Bank's processed transaction number")
		Example("1169716b65891lI6")
	})
	Required("MessageReference", "MessageDateTime", "PaymentRef", "AccountNumber", "Amount", "TransactionDate",
		"EventType", "Currency", "ExchangeRate", "Narration", "CustMemo", "ValueDate", "EntryDate", "TransactionId")
})

var CustMemo = Type("CustMemo", func() {
	Description("Cust Memo")

	Attribute("CustMemoLine1", String, func() {
		Description("CustMemo CustMemoLine1")
		Example("728210595 ABD01")
	})
	Attribute("CustMemoLine2", String, func() {
		Description("CustMemo CustMemoLine2")
	})
	Attribute("CustMemoLine3", String, func() {
		Description("CustMemo CustMemoLine2")
	})
})
