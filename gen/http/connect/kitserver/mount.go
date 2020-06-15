// Code generated by goa v3.1.3, DO NOT EDIT.
//
// connect go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/coop-go/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountAccountBalanceHandler configures the mux to serve the "connect" service
// "AccountBalance" endpoint.
func MountAccountBalanceHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Enquiry/AccountBalance/1.0.0", f)
}

// MountAccountFullStatementHandler configures the mux to serve the "connect"
// service "AccountFullStatement" endpoint.
func MountAccountFullStatementHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Enquiry/FullStatement/Account/1.0.0", f)
}

// MountAccountMiniStatementHandler configures the mux to serve the "connect"
// service "AccountMiniStatement" endpoint.
func MountAccountMiniStatementHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/MiniStatement/Account/1.0.0", f)
}

// MountAccountTransactionsHandler configures the mux to serve the "connect"
// service "AccountTransactions" endpoint.
func MountAccountTransactionsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Enquiry/AccountTransactions/1.0.0", f)
}

// MountAccountValidationHandler configures the mux to serve the "connect"
// service "AccountValidation" endpoint.
func MountAccountValidationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Enquiry/Validation/Account/1.0.0", f)
}

// MountExchangeRateHandler configures the mux to serve the "connect" service
// "ExchangeRate" endpoint.
func MountExchangeRateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Enquiry/ExchangeRate/1.0.0", f)
}

// MountIFTAccountToAccountHandler configures the mux to serve the "connect"
// service "IFTAccountToAccount" endpoint.
func MountIFTAccountToAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/FundsTransfer/Internal/A2A/2.0.0", f)
}

// MountINSSimulationHandler configures the mux to serve the "connect" service
// "INSSimulation" endpoint.
func MountINSSimulationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Notifications/INS/Simulation/1.0.0", f)
}

// MountPesaLinkSendToAccountHandler configures the mux to serve the "connect"
// service "PesaLinkSendToAccount" endpoint.
func MountPesaLinkSendToAccountHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/FundsTransfer/External/A2A/PesaLink/1.0.0", f)
}

// MountPesaLinkSendToPhoneHandler configures the mux to serve the "connect"
// service "PesaLinkSendToPhone" endpoint.
func MountPesaLinkSendToPhoneHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/FundsTransfer/External/A2M/PesaLink/1.0.0", f)
}

// MountSendToMPesaHandler configures the mux to serve the "connect" service
// "SendToMPesa" endpoint.
func MountSendToMPesaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/FundsTransfer/External/A2M/Mpesa/v1.0.0", f)
}

// MountTransactionStatusHandler configures the mux to serve the "connect"
// service "TransactionStatus" endpoint.
func MountTransactionStatusHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/Enquiry/TransactionStatus/2.0.0", f)
}
