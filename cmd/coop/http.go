package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	connect "github.com/wondenge/coop-go/gen/connect"
	health "github.com/wondenge/coop-go/gen/health"
	connectkitsvr "github.com/wondenge/coop-go/gen/http/connect/kitserver"
	connectsvr "github.com/wondenge/coop-go/gen/http/connect/server"
	healthkitsvr "github.com/wondenge/coop-go/gen/http/health/kitserver"
	healthsvr "github.com/wondenge/coop-go/gen/http/health/server"
	swaggerkitsvr "github.com/wondenge/coop-go/gen/http/swagger/kitserver"
	swaggersvr "github.com/wondenge/coop-go/gen/http/swagger/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, connectEndpoints *connect.Endpoints, healthEndpoints *health.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		connectAccountBalanceHandler        *kithttp.Server
		connectAccountFullStatementHandler  *kithttp.Server
		connectAccountMiniStatementHandler  *kithttp.Server
		connectAccountTransactionsHandler   *kithttp.Server
		connectAccountValidationHandler     *kithttp.Server
		connectExchangeRateHandler          *kithttp.Server
		connectIFTAccountToAccountHandler   *kithttp.Server
		connectINSSimulationHandler         *kithttp.Server
		connectPesaLinkSendToAccountHandler *kithttp.Server
		connectPesaLinkSendToPhoneHandler   *kithttp.Server
		connectSendToMPesaHandler           *kithttp.Server
		connectTransactionStatusHandler     *kithttp.Server
		connectServer                       *connectsvr.Server
		healthShowHandler                   *kithttp.Server
		healthServer                        *healthsvr.Server
		swaggerServer                       *swaggersvr.Server
	)
	{
		eh := errorHandler(logger)
		connectAccountBalanceHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.AccountBalance),
			connectkitsvr.DecodeAccountBalanceRequest(mux, dec),
			connectkitsvr.EncodeAccountBalanceResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeAccountBalanceError(enc, nil)),
		)
		connectAccountFullStatementHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.AccountFullStatement),
			connectkitsvr.DecodeAccountFullStatementRequest(mux, dec),
			connectkitsvr.EncodeAccountFullStatementResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeAccountFullStatementError(enc, nil)),
		)
		connectAccountMiniStatementHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.AccountMiniStatement),
			connectkitsvr.DecodeAccountMiniStatementRequest(mux, dec),
			connectkitsvr.EncodeAccountMiniStatementResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeAccountMiniStatementError(enc, nil)),
		)
		connectAccountTransactionsHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.AccountTransactions),
			connectkitsvr.DecodeAccountTransactionsRequest(mux, dec),
			connectkitsvr.EncodeAccountTransactionsResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeAccountTransactionsError(enc, nil)),
		)
		connectAccountValidationHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.AccountValidation),
			connectkitsvr.DecodeAccountValidationRequest(mux, dec),
			connectkitsvr.EncodeAccountValidationResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeAccountValidationError(enc, nil)),
		)
		connectExchangeRateHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.ExchangeRate),
			connectkitsvr.DecodeExchangeRateRequest(mux, dec),
			connectkitsvr.EncodeExchangeRateResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeExchangeRateError(enc, nil)),
		)
		connectIFTAccountToAccountHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.IFTAccountToAccount),
			connectkitsvr.DecodeIFTAccountToAccountRequest(mux, dec),
			connectkitsvr.EncodeIFTAccountToAccountResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeIFTAccountToAccountError(enc, nil)),
		)
		connectINSSimulationHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.INSSimulation),
			connectkitsvr.DecodeINSSimulationRequest(mux, dec),
			connectkitsvr.EncodeINSSimulationResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeINSSimulationError(enc, nil)),
		)
		connectPesaLinkSendToAccountHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.PesaLinkSendToAccount),
			connectkitsvr.DecodePesaLinkSendToAccountRequest(mux, dec),
			connectkitsvr.EncodePesaLinkSendToAccountResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodePesaLinkSendToAccountError(enc, nil)),
		)
		connectPesaLinkSendToPhoneHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.PesaLinkSendToPhone),
			connectkitsvr.DecodePesaLinkSendToPhoneRequest(mux, dec),
			connectkitsvr.EncodePesaLinkSendToPhoneResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodePesaLinkSendToPhoneError(enc, nil)),
		)
		connectSendToMPesaHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.SendToMPesa),
			connectkitsvr.DecodeSendToMPesaRequest(mux, dec),
			connectkitsvr.EncodeSendToMPesaResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeSendToMPesaError(enc, nil)),
		)
		connectTransactionStatusHandler = kithttp.NewServer(
			endpoint.Endpoint(connectEndpoints.TransactionStatus),
			connectkitsvr.DecodeTransactionStatusRequest(mux, dec),
			connectkitsvr.EncodeTransactionStatusResponse(enc),
			kithttp.ServerErrorEncoder(connectkitsvr.EncodeTransactionStatusError(enc, nil)),
		)
		connectServer = connectsvr.New(connectEndpoints, mux, dec, enc, eh, nil)
		healthShowHandler = kithttp.NewServer(
			endpoint.Endpoint(healthEndpoints.Show),
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			healthkitsvr.EncodeShowResponse(enc),
		)
		healthServer = healthsvr.New(healthEndpoints, mux, dec, enc, eh, nil)
		swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	connectkitsvr.MountAccountBalanceHandler(mux, connectAccountBalanceHandler)
	connectkitsvr.MountAccountFullStatementHandler(mux, connectAccountFullStatementHandler)
	connectkitsvr.MountAccountMiniStatementHandler(mux, connectAccountMiniStatementHandler)
	connectkitsvr.MountAccountTransactionsHandler(mux, connectAccountTransactionsHandler)
	connectkitsvr.MountAccountValidationHandler(mux, connectAccountValidationHandler)
	connectkitsvr.MountExchangeRateHandler(mux, connectExchangeRateHandler)
	connectkitsvr.MountIFTAccountToAccountHandler(mux, connectIFTAccountToAccountHandler)
	connectkitsvr.MountINSSimulationHandler(mux, connectINSSimulationHandler)
	connectkitsvr.MountPesaLinkSendToAccountHandler(mux, connectPesaLinkSendToAccountHandler)
	connectkitsvr.MountPesaLinkSendToPhoneHandler(mux, connectPesaLinkSendToPhoneHandler)
	connectkitsvr.MountSendToMPesaHandler(mux, connectSendToMPesaHandler)
	connectkitsvr.MountTransactionStatusHandler(mux, connectTransactionStatusHandler)
	healthkitsvr.MountShowHandler(mux, healthShowHandler)
	swaggerkitsvr.MountGenHTTPOpenapiJSON(mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(logger)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range connectServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range healthServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range swaggerServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
