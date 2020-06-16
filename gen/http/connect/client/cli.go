// Code generated by goa v3.1.3, DO NOT EDIT.
//
// connect HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wondenge/coop-go/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	connect "github.com/wondenge/coop-go/gen/connect"
	goa "goa.design/goa/v3/pkg"
)

// BuildAccountBalancePayload builds the payload for the connect AccountBalance
// endpoint from CLI flags.
func BuildAccountBalancePayload(connectAccountBalanceBody string, connectAccountBalanceAccessToken string) (*connect.AccountBalancePayload, error) {
	var err error
	var body AccountBalanceRequestBody
	{
		err = json.Unmarshal([]byte(connectAccountBalanceBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"AccountNumber\": \"36001873000\",\n      \"MessageReference\": \"40ca18c6765086089a1\"\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if utf8.RuneCountInString(body.AccountNumber) < 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
		}
		if utf8.RuneCountInString(body.AccountNumber) > 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectAccountBalanceAccessToken
	}
	v := &connect.AccountBalancePayload{
		MessageReference: body.MessageReference,
		AccountNumber:    body.AccountNumber,
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildAccountFullStatementPayload builds the payload for the connect
// AccountFullStatement endpoint from CLI flags.
func BuildAccountFullStatementPayload(connectAccountFullStatementBody string, connectAccountFullStatementAccessToken string) (*connect.AccountFullStatementPayload, error) {
	var err error
	var body AccountFullStatementRequestBody
	{
		err = json.Unmarshal([]byte(connectAccountFullStatementBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"AccountNumber\": \"36001873000\",\n      \"EndDate\": \"2019-07-01\",\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"StartDate\": \"2019-01-01\"\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if utf8.RuneCountInString(body.AccountNumber) < 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
		}
		if utf8.RuneCountInString(body.AccountNumber) > 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.StartDate", body.StartDate, goa.FormatDate))

		err = goa.MergeErrors(err, goa.ValidateFormat("body.EndDate", body.EndDate, goa.FormatDate))

		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectAccountFullStatementAccessToken
	}
	v := &connect.AccountFullStatementPayload{
		MessageReference: body.MessageReference,
		AccountNumber:    body.AccountNumber,
		StartDate:        body.StartDate,
		EndDate:          body.EndDate,
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildAccountMiniStatementPayload builds the payload for the connect
// AccountMiniStatement endpoint from CLI flags.
func BuildAccountMiniStatementPayload(connectAccountMiniStatementBody string, connectAccountMiniStatementAccessToken string) (*connect.AccountMiniStatementPayload, error) {
	var err error
	var body AccountMiniStatementRequestBody
	{
		err = json.Unmarshal([]byte(connectAccountMiniStatementBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"AccountNumber\": \"36001873000\",\n      \"MessageReference\": \"40ca18c6765086089a1\"\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if utf8.RuneCountInString(body.AccountNumber) < 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
		}
		if utf8.RuneCountInString(body.AccountNumber) > 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectAccountMiniStatementAccessToken
	}
	v := &connect.AccountMiniStatementPayload{
		MessageReference: body.MessageReference,
		AccountNumber:    body.AccountNumber,
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildAccountTransactionsPayload builds the payload for the connect
// AccountTransactions endpoint from CLI flags.
func BuildAccountTransactionsPayload(connectAccountTransactionsBody string, connectAccountTransactionsAccessToken string) (*connect.AccountTransactionsPayload, error) {
	var err error
	var body AccountTransactionsRequestBody
	{
		err = json.Unmarshal([]byte(connectAccountTransactionsBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"AccountNumber\": \"36001873000\",\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"NoOfTransactions\": 1\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if utf8.RuneCountInString(body.AccountNumber) < 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
		}
		if utf8.RuneCountInString(body.AccountNumber) > 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectAccountTransactionsAccessToken
	}
	v := &connect.AccountTransactionsPayload{
		MessageReference: body.MessageReference,
		AccountNumber:    body.AccountNumber,
		NoOfTransactions: body.NoOfTransactions,
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildAccountValidationPayload builds the payload for the connect
// AccountValidation endpoint from CLI flags.
func BuildAccountValidationPayload(connectAccountValidationBody string, connectAccountValidationAccessToken string) (*connect.AccountValidationPayload, error) {
	var err error
	var body AccountValidationRequestBody
	{
		err = json.Unmarshal([]byte(connectAccountValidationBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"AccountNumber\": \"36001873000\",\n      \"MessageReference\": \"40ca18c6765086089a1\"\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if utf8.RuneCountInString(body.AccountNumber) < 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
		}
		if utf8.RuneCountInString(body.AccountNumber) > 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectAccountValidationAccessToken
	}
	v := &connect.AccountValidationPayload{
		MessageReference: body.MessageReference,
		AccountNumber:    body.AccountNumber,
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildExchangeRatePayload builds the payload for the connect ExchangeRate
// endpoint from CLI flags.
func BuildExchangeRatePayload(connectExchangeRateBody string, connectExchangeRateAccessToken string) (*connect.ExchangeRatePayload, error) {
	var err error
	var body ExchangeRateRequestBody
	{
		err = json.Unmarshal([]byte(connectExchangeRateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"FromCurrencyCode\": \"KES\",\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"ToCurrencyCode\": \"USD\"\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if !(body.FromCurrencyCode == "USD" || body.FromCurrencyCode == "KES" || body.FromCurrencyCode == "EUR" || body.FromCurrencyCode == "GBP" || body.FromCurrencyCode == "AUD" || body.FromCurrencyCode == "CHF" || body.FromCurrencyCode == "CAD" || body.FromCurrencyCode == "ZAR") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.FromCurrencyCode", body.FromCurrencyCode, []interface{}{"USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR"}))
		}
		if !(body.ToCurrencyCode == "USD" || body.ToCurrencyCode == "KES" || body.ToCurrencyCode == "EUR" || body.ToCurrencyCode == "GBP" || body.ToCurrencyCode == "AUD" || body.ToCurrencyCode == "CHF" || body.ToCurrencyCode == "CAD" || body.ToCurrencyCode == "ZAR") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.ToCurrencyCode", body.ToCurrencyCode, []interface{}{"USD", "KES", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectExchangeRateAccessToken
	}
	v := &connect.ExchangeRatePayload{
		MessageReference: body.MessageReference,
		FromCurrencyCode: body.FromCurrencyCode,
		ToCurrencyCode:   body.ToCurrencyCode,
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildIFTAccountToAccountPayload builds the payload for the connect
// IFTAccountToAccount endpoint from CLI flags.
func BuildIFTAccountToAccountPayload(connectIFTAccountToAccountBody string, connectIFTAccountToAccountAccessToken string) (*connect.IFTAccountToAccountTXNRequest, error) {
	var err error
	var body IFTAccountToAccountRequestBody
	{
		err = json.Unmarshal([]byte(connectIFTAccountToAccountBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"CallBackUrl\": \"https://yourdomain.com/ftresponse\",\n      \"Destinations\": {\n         \"DestinationsTXNRequest\": [\n            {\n               \"AccountNumber\": \"54321987654321\",\n               \"Amount\": 777,\n               \"Narration\": \"Supplier Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\",\n               \"TransactionCurrency\": \"KES\"\n            }\n         ]\n      },\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"Source\": {\n         \"AccountNumber\": \"36001873000\",\n         \"Amount\": 777,\n         \"Narration\": \"Supplier Payment\",\n         \"TransactionCurrency\": \"KES\"\n      }\n   }'")
		}
		if body.Source == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Source", "body"))
		}
		if body.Destinations == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Destinations", "body"))
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if body.Source != nil {
			if err2 := ValidateSourceAccountTXNRequestRequestBody(body.Source); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Destinations != nil {
			if err2 := ValidateDestinationsTXNRequestRequestBody(body.Destinations); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectIFTAccountToAccountAccessToken
	}
	v := &connect.IFTAccountToAccountTXNRequest{
		MessageReference: body.MessageReference,
		CallBackURL:      body.CallBackURL,
	}
	if body.Source != nil {
		v.Source = marshalSourceAccountTXNRequestRequestBodyToConnectSourceAccountTXNRequest(body.Source)
	}
	if body.Destinations != nil {
		v.Destinations = marshalDestinationsTXNRequestRequestBodyToConnectDestinationsTXNRequest(body.Destinations)
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildINSSimulationPayload builds the payload for the connect INSSimulation
// endpoint from CLI flags.
func BuildINSSimulationPayload(connectINSSimulationBody string, connectINSSimulationAccessToken string) (*connect.INSTransactionSimulationRequest, error) {
	var err error
	var body INSSimulationRequestBody
	{
		err = json.Unmarshal([]byte(connectINSSimulationBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"AccountNumber\": \"54321987654321\",\n      \"Amount\": 120777.45,\n      \"Currency\": \"ZAR\",\n      \"CustMemo\": {\n         \"CustMemoLine1\": \"728210595 ABD01\",\n         \"CustMemoLine2\": \"Quae aperiam.\",\n         \"CustMemoLine3\": \"Dolores in officia et itaque et ut.\"\n      },\n      \"EntryDate\": \"20190301\",\n      \"EventType\": \"DEBIT\",\n      \"ExchangeRate\": 1,\n      \"MessageDateTime\": \"2017-12-04T09:27:00\",\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"Narration\": \"Supplier Payments\",\n      \"NotificationCode\": \"Itaque vel qui incidunt repellat.\",\n      \"PaymentRef\": \"SFI427E9136D7D3F21C2C89\",\n      \"ServiceName\": \"Aliquid maiores.\",\n      \"TransactionDate\": \"20190301165420\",\n      \"TransactionId\": \"1169716b65891lI6\",\n      \"ValueDate\": \"20190301\"\n   }'")
		}
		if body.CustMemo == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("CustMemo", "body"))
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.MessageDateTime", body.MessageDateTime, goa.FormatDateTime))

		if utf8.RuneCountInString(body.AccountNumber) < 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, true))
		}
		if utf8.RuneCountInString(body.AccountNumber) > 14 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.AccountNumber", body.AccountNumber, utf8.RuneCountInString(body.AccountNumber), 14, false))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.TransactionDate", body.TransactionDate, goa.FormatDateTime))

		if !(body.Currency == "KES" || body.Currency == "USD" || body.Currency == "EUR" || body.Currency == "GBP" || body.Currency == "AUD" || body.Currency == "CHF" || body.Currency == "CAD" || body.Currency == "ZAR") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.Currency", body.Currency, []interface{}{"KES", "USD", "EUR", "GBP", "AUD", "CHF", "CAD", "ZAR"}))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.ValueDate", body.ValueDate, goa.FormatDate))

		err = goa.MergeErrors(err, goa.ValidateFormat("body.EntryDate", body.EntryDate, goa.FormatDate))

		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectINSSimulationAccessToken
	}
	v := &connect.INSTransactionSimulationRequest{
		MessageReference: body.MessageReference,
		MessageDateTime:  body.MessageDateTime,
		ServiceName:      body.ServiceName,
		NotificationCode: body.NotificationCode,
		PaymentRef:       body.PaymentRef,
		AccountNumber:    body.AccountNumber,
		Amount:           body.Amount,
		TransactionDate:  body.TransactionDate,
		EventType:        body.EventType,
		Currency:         body.Currency,
		ExchangeRate:     body.ExchangeRate,
		Narration:        body.Narration,
		ValueDate:        body.ValueDate,
		EntryDate:        body.EntryDate,
		TransactionID:    body.TransactionID,
	}
	if body.CustMemo != nil {
		v.CustMemo = marshalCustMemoRequestBodyToConnectCustMemo(body.CustMemo)
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildPesaLinkSendToAccountPayload builds the payload for the connect
// PesaLinkSendToAccount endpoint from CLI flags.
func BuildPesaLinkSendToAccountPayload(connectPesaLinkSendToAccountBody string, connectPesaLinkSendToAccountAccessToken string) (*connect.PesaLinkSendToAccountTransactionRequest, error) {
	var err error
	var body PesaLinkSendToAccountRequestBody
	{
		err = json.Unmarshal([]byte(connectPesaLinkSendToAccountBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"CallBackUrl\": \"https://yourdomain.com/ft-callback\",\n      \"Destinations\": {\n         \"DestinationAccountTransactionRequest\": [\n            {\n               \"Amount\": 777,\n               \"MobileNumber\": \"07xxxxxxxx\",\n               \"Narration\": \"Stationary Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\"\n            },\n            {\n               \"Amount\": 777,\n               \"MobileNumber\": \"07xxxxxxxx\",\n               \"Narration\": \"Stationary Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\"\n            }\n         ]\n      },\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"Source\": {\n         \"AccountNumber\": \"36001873000\",\n         \"Amount\": 777,\n         \"Narration\": \"Supplier Payment\",\n         \"ResponseDescription\": \"Success\",\n         \"TransactionCurrency\": \"KES\"\n      }\n   }'")
		}
		if body.Source == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Source", "body"))
		}
		if body.Destinations == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Destinations", "body"))
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if body.Source != nil {
			if err2 := ValidateSourceAccountTransactionRequestRequestBody(body.Source); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Destinations != nil {
			if err2 := ValidateDestinationsTransactionRequestRequestBody(body.Destinations); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectPesaLinkSendToAccountAccessToken
	}
	v := &connect.PesaLinkSendToAccountTransactionRequest{
		MessageReference: body.MessageReference,
		CallBackURL:      body.CallBackURL,
	}
	if body.Source != nil {
		v.Source = marshalSourceAccountTransactionRequestRequestBodyToConnectSourceAccountTransactionRequest(body.Source)
	}
	if body.Destinations != nil {
		v.Destinations = marshalDestinationsTransactionRequestRequestBodyToConnectDestinationsTransactionRequest(body.Destinations)
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildPesaLinkSendToPhonePayload builds the payload for the connect
// PesaLinkSendToPhone endpoint from CLI flags.
func BuildPesaLinkSendToPhonePayload(connectPesaLinkSendToPhoneBody string, connectPesaLinkSendToPhoneAccessToken string) (*connect.PesaLinkSendToPhoneTransactionRequest, error) {
	var err error
	var body PesaLinkSendToPhoneRequestBody
	{
		err = json.Unmarshal([]byte(connectPesaLinkSendToPhoneBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"CallBackUrl\": \"https://yourdomain.com/ft-callback\",\n      \"Destinations\": {\n         \"DestinationAccountTransactionRequest\": [\n            {\n               \"Amount\": 777,\n               \"MobileNumber\": \"07xxxxxxxx\",\n               \"Narration\": \"Stationary Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\"\n            },\n            {\n               \"Amount\": 777,\n               \"MobileNumber\": \"07xxxxxxxx\",\n               \"Narration\": \"Stationary Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\"\n            }\n         ]\n      },\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"Source\": {\n         \"AccountNumber\": \"36001873000\",\n         \"Amount\": 777,\n         \"Narration\": \"Supplier Payment\",\n         \"ResponseDescription\": \"Success\",\n         \"TransactionCurrency\": \"KES\"\n      }\n   }'")
		}
		if body.Source == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Source", "body"))
		}
		if body.Destinations == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Destinations", "body"))
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if body.Source != nil {
			if err2 := ValidateSourceAccountTransactionRequestRequestBody(body.Source); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Destinations != nil {
			if err2 := ValidateDestinationsTransactionRequestRequestBody(body.Destinations); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectPesaLinkSendToPhoneAccessToken
	}
	v := &connect.PesaLinkSendToPhoneTransactionRequest{
		MessageReference: body.MessageReference,
		CallBackURL:      body.CallBackURL,
	}
	if body.Source != nil {
		v.Source = marshalSourceAccountTransactionRequestRequestBodyToConnectSourceAccountTransactionRequest(body.Source)
	}
	if body.Destinations != nil {
		v.Destinations = marshalDestinationsTransactionRequestRequestBodyToConnectDestinationsTransactionRequest(body.Destinations)
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildSendToMPesaPayload builds the payload for the connect SendToMPesa
// endpoint from CLI flags.
func BuildSendToMPesaPayload(connectSendToMPesaBody string, connectSendToMPesaAccessToken string) (*connect.SendToMpesaTransactionRequest, error) {
	var err error
	var body SendToMPesaRequestBody
	{
		err = json.Unmarshal([]byte(connectSendToMPesaBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"CallBackUrl\": \"https://yourdomain.com/ft-callback\",\n      \"Destinations\": {\n         \"DestinationAccountTransactionRequest\": [\n            {\n               \"Amount\": 777,\n               \"MobileNumber\": \"07xxxxxxxx\",\n               \"Narration\": \"Stationary Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\"\n            },\n            {\n               \"Amount\": 777,\n               \"MobileNumber\": \"07xxxxxxxx\",\n               \"Narration\": \"Stationary Payment\",\n               \"ReferenceNumber\": \"40ca18c6765086089a1_1\"\n            }\n         ]\n      },\n      \"MessageReference\": \"40ca18c6765086089a1\",\n      \"Source\": {\n         \"AccountNumber\": \"36001873000\",\n         \"Amount\": 777,\n         \"Narration\": \"Supplier Payment\",\n         \"ResponseDescription\": \"Success\",\n         \"TransactionCurrency\": \"KES\"\n      }\n   }'")
		}
		if body.Source == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Source", "body"))
		}
		if body.Destinations == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("Destinations", "body"))
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if body.Source != nil {
			if err2 := ValidateSourceAccountTransactionRequestRequestBody(body.Source); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Destinations != nil {
			if err2 := ValidateDestinationsTransactionRequestRequestBody(body.Destinations); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectSendToMPesaAccessToken
	}
	v := &connect.SendToMpesaTransactionRequest{
		MessageReference: body.MessageReference,
		CallBackURL:      body.CallBackURL,
	}
	if body.Source != nil {
		v.Source = marshalSourceAccountTransactionRequestRequestBodyToConnectSourceAccountTransactionRequest(body.Source)
	}
	if body.Destinations != nil {
		v.Destinations = marshalDestinationsTransactionRequestRequestBodyToConnectDestinationsTransactionRequest(body.Destinations)
	}
	v.AccessToken = &accessToken

	return v, nil
}

// BuildTransactionStatusPayload builds the payload for the connect
// TransactionStatus endpoint from CLI flags.
func BuildTransactionStatusPayload(connectTransactionStatusBody string, connectTransactionStatusAccessToken string) (*connect.FTTransactionStatusPayload, error) {
	var err error
	var body TransactionStatusRequestBody
	{
		err = json.Unmarshal([]byte(connectTransactionStatusBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"MessageReference\": \"40ca18c6765086089a1\"\n   }'")
		}
		if utf8.RuneCountInString(body.MessageReference) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 1, true))
		}
		if utf8.RuneCountInString(body.MessageReference) > 27 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.MessageReference", body.MessageReference, utf8.RuneCountInString(body.MessageReference), 27, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var accessToken string
	{
		accessToken = connectTransactionStatusAccessToken
	}
	v := &connect.FTTransactionStatusPayload{
		MessageReference: body.MessageReference,
	}
	v.AccessToken = &accessToken

	return v, nil
}
