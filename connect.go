package connectapi

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	connect "github.com/wondenge/coop-go/gen/connect"
	"goa.design/goa/v3/security"
)

// connect service example implementation.
// The example methods log the requests and return zero values.
type connectsrvc struct {
	logger log.Logger
}

// NewConnect returns the connect service implementation.
func NewConnect(logger log.Logger) connect.Service {
	return &connectsrvc{logger}
}

// BasicAuth implements the authorization logic for service "connect" for the
// "basic" security scheme.
func (s *connectsrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// Post an Account Balance Enquiry Request
func (s *connectsrvc) AccountBalance(ctx context.Context, p *connect.AccountBalancePayload) (res *connect.AccountBalanceSuccessResponse, err error) {
	res = &connect.AccountBalanceSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountBalance"))
	return
}

// Post an Account Full Statement Enquiry Request
func (s *connectsrvc) AccountFullStatement(ctx context.Context, p *connect.AccountFullStatementPayload) (res *connect.AccountFullStatementSuccessResponse, err error) {
	res = &connect.AccountFullStatementSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountFullStatement"))
	return
}

// Post an Account Mini Statement Enquiry Request
func (s *connectsrvc) AccountMiniStatement(ctx context.Context, p *connect.AccountMiniStatementPayload) (res *connect.AccountMiniStatementSuccessResponse, err error) {
	res = &connect.AccountMiniStatementSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountMiniStatement"))
	return
}

// Post an Account Transactions Enquiry Request
func (s *connectsrvc) AccountTransactions(ctx context.Context, p *connect.AccountTransactionsPayload) (res *connect.AccountTransactionsSuccessResponse, err error) {
	res = &connect.AccountTransactionsSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountTransactions"))
	return
}

// Post an Account Validation Enquiry Request
func (s *connectsrvc) AccountValidation(ctx context.Context, p *connect.AccountValidationPayload) (res *connect.AccountValidationSuccessResponse, err error) {
	res = &connect.AccountValidationSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.AccountValidation"))
	return
}

// Post an Exchange Rate Enquiry Request
func (s *connectsrvc) ExchangeRate(ctx context.Context, p *connect.ExchangeRatePayload) (res *connect.ExchangeRateSuccessResponse, err error) {
	res = &connect.ExchangeRateSuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.ExchangeRate"))
	return
}

// Post an Internal Funds Transfer Account to Account Transaction
func (s *connectsrvc) IFTAccountToAccount(ctx context.Context, p *connect.IFTAccountToAccountTXNRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.IFTAccountToAccount"))
	return
}

// Post a Debit/Credit Account Transaction Event Type Notification Simulation
// Request
func (s *connectsrvc) INSSimulation(ctx context.Context, p *connect.INSTransactionSimulationRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.INSSimulation"))
	return
}

// Post a PesaLink Funds Transfer Send to Account Transaction
func (s *connectsrvc) PesaLinkSendToAccount(ctx context.Context, p *connect.PesaLinkSendToAccountTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToAccount"))
	return
}

// Post a PesaLink Funds Transfer Send to Phone Transaction
func (s *connectsrvc) PesaLinkSendToPhone(ctx context.Context, p *connect.PesaLinkSendToPhoneTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.PesaLinkSendToPhone"))
	return
}

// Post a Send To M-Pesa Funds Transfer Transaction
func (s *connectsrvc) SendToMPesa(ctx context.Context, p *connect.SendToMpesaTransactionRequest) (res *connect.SuccessAcknowledgement, err error) {
	res = &connect.SuccessAcknowledgement{}
	s.logger.Log("info", fmt.Sprintf("connect.SendToMPesa"))
	return
}

// Post a Transaction Status Enquiry Request
func (s *connectsrvc) TransactionStatus(ctx context.Context, p *connect.FTTransactionStatusPayload) (res *connect.SuccessResponse, err error) {
	res = &connect.SuccessResponse{}
	s.logger.Log("info", fmt.Sprintf("connect.TransactionStatus"))
	return
}

// Creates a valid JWT
func (s *connectsrvc) Token(ctx context.Context, p *connect.TokenPayload) (res *connect.Creds, err error) {
	res = &connect.Creds{}
	s.logger.Log("info", fmt.Sprintf("connect.token"))
	return
}
