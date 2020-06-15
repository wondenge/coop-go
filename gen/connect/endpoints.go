// Code generated by goa v3.1.3, DO NOT EDIT.
//
// connect endpoints
//
// Command:
// $ goa gen github.com/wondenge/coop-go/design

package connect

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "connect" service endpoints.
type Endpoints struct {
	AccountBalance        endpoint.Endpoint
	AccountFullStatement  endpoint.Endpoint
	AccountMiniStatement  endpoint.Endpoint
	AccountTransactions   endpoint.Endpoint
	AccountValidation     endpoint.Endpoint
	ExchangeRate          endpoint.Endpoint
	IFTAccountToAccount   endpoint.Endpoint
	INSSimulation         endpoint.Endpoint
	PesaLinkSendToAccount endpoint.Endpoint
	PesaLinkSendToPhone   endpoint.Endpoint
	SendToMPesa           endpoint.Endpoint
	TransactionStatus     endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "connect" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		AccountBalance:        NewAccountBalanceEndpoint(s),
		AccountFullStatement:  NewAccountFullStatementEndpoint(s),
		AccountMiniStatement:  NewAccountMiniStatementEndpoint(s),
		AccountTransactions:   NewAccountTransactionsEndpoint(s),
		AccountValidation:     NewAccountValidationEndpoint(s),
		ExchangeRate:          NewExchangeRateEndpoint(s),
		IFTAccountToAccount:   NewIFTAccountToAccountEndpoint(s),
		INSSimulation:         NewINSSimulationEndpoint(s),
		PesaLinkSendToAccount: NewPesaLinkSendToAccountEndpoint(s),
		PesaLinkSendToPhone:   NewPesaLinkSendToPhoneEndpoint(s),
		SendToMPesa:           NewSendToMPesaEndpoint(s),
		TransactionStatus:     NewTransactionStatusEndpoint(s),
	}
}

// Use applies the given middleware to all the "connect" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.AccountBalance = m(e.AccountBalance)
	e.AccountFullStatement = m(e.AccountFullStatement)
	e.AccountMiniStatement = m(e.AccountMiniStatement)
	e.AccountTransactions = m(e.AccountTransactions)
	e.AccountValidation = m(e.AccountValidation)
	e.ExchangeRate = m(e.ExchangeRate)
	e.IFTAccountToAccount = m(e.IFTAccountToAccount)
	e.INSSimulation = m(e.INSSimulation)
	e.PesaLinkSendToAccount = m(e.PesaLinkSendToAccount)
	e.PesaLinkSendToPhone = m(e.PesaLinkSendToPhone)
	e.SendToMPesa = m(e.SendToMPesa)
	e.TransactionStatus = m(e.TransactionStatus)
}

// NewAccountBalanceEndpoint returns an endpoint function that calls the method
// "AccountBalance" of service "connect".
func NewAccountBalanceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AccountBalancePayload)
		res, err := s.AccountBalance(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccountBalanceSuccessResponse(res, "default")
		return vres, nil
	}
}

// NewAccountFullStatementEndpoint returns an endpoint function that calls the
// method "AccountFullStatement" of service "connect".
func NewAccountFullStatementEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AccountFullStatementPayload)
		res, err := s.AccountFullStatement(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccountFullStatementSuccessResponse(res, "default")
		return vres, nil
	}
}

// NewAccountMiniStatementEndpoint returns an endpoint function that calls the
// method "AccountMiniStatement" of service "connect".
func NewAccountMiniStatementEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AccountMiniStatementPayload)
		res, err := s.AccountMiniStatement(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccountMiniStatementSuccessResponse(res, "default")
		return vres, nil
	}
}

// NewAccountTransactionsEndpoint returns an endpoint function that calls the
// method "AccountTransactions" of service "connect".
func NewAccountTransactionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AccountTransactionsPayload)
		res, err := s.AccountTransactions(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccountTransactionsSuccessResponse(res, "default")
		return vres, nil
	}
}

// NewAccountValidationEndpoint returns an endpoint function that calls the
// method "AccountValidation" of service "connect".
func NewAccountValidationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AccountValidationPayload)
		res, err := s.AccountValidation(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccountValidationSuccessResponse(res, "default")
		return vres, nil
	}
}

// NewExchangeRateEndpoint returns an endpoint function that calls the method
// "ExchangeRate" of service "connect".
func NewExchangeRateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ExchangeRatePayload)
		res, err := s.ExchangeRate(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedExchangeRateSuccessResponse(res, "default")
		return vres, nil
	}
}

// NewIFTAccountToAccountEndpoint returns an endpoint function that calls the
// method "IFTAccountToAccount" of service "connect".
func NewIFTAccountToAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IFTAccountToAccountTXNRequest)
		res, err := s.IFTAccountToAccount(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessAcknowledgement(res, "default")
		return vres, nil
	}
}

// NewINSSimulationEndpoint returns an endpoint function that calls the method
// "INSSimulation" of service "connect".
func NewINSSimulationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*INSTransactionSimulationRequest)
		res, err := s.INSSimulation(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessAcknowledgement(res, "default")
		return vres, nil
	}
}

// NewPesaLinkSendToAccountEndpoint returns an endpoint function that calls the
// method "PesaLinkSendToAccount" of service "connect".
func NewPesaLinkSendToAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PesaLinkSendToAccountTransactionRequest)
		res, err := s.PesaLinkSendToAccount(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessAcknowledgement(res, "default")
		return vres, nil
	}
}

// NewPesaLinkSendToPhoneEndpoint returns an endpoint function that calls the
// method "PesaLinkSendToPhone" of service "connect".
func NewPesaLinkSendToPhoneEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PesaLinkSendToPhoneTransactionRequest)
		res, err := s.PesaLinkSendToPhone(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessAcknowledgement(res, "default")
		return vres, nil
	}
}

// NewSendToMPesaEndpoint returns an endpoint function that calls the method
// "SendToMPesa" of service "connect".
func NewSendToMPesaEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SendToMpesaTransactionRequest)
		res, err := s.SendToMPesa(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessAcknowledgement(res, "default")
		return vres, nil
	}
}

// NewTransactionStatusEndpoint returns an endpoint function that calls the
// method "TransactionStatus" of service "connect".
func NewTransactionStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FTTransactionStatusPayload)
		res, err := s.TransactionStatus(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccessResponse(res, "default")
		return vres, nil
	}
}