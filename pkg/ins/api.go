/*
 * INSSimulation
 *
 * INS Simulation API enables you to simulate instant notifications reception on your specified notification endpoint
 *
 * API version: 1.0.0
 *
 *
 */

package ins

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// SendTransactionINSService SendTransactionINSAPI service
type SendTransactionINSService service

// INSTransactionOpts Optional parameters for the method 'INSTransaction'
type INSTransactionOpts struct {
	EndpointCredential optional.String
}

/*
INSTransaction Post a Debit/Credit Account Transaction Event Type Notification Simulation Request
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param notificationEndpoint Notification Endpoint
 * @param insTransactionSimulationRequest Debit/Credit Account Transaction Event Type Notification Simulation Request Body
 * @param optional nil or *INSTransactionOpts - Optional Parameters:
 * @param "EndpointCredential" (optional.String) -  Basic Auth Endpoint Credential
@return SuccessAcknowledgement
*/
func (a *SendTransactionINSService) INSTransaction(ctx _context.Context, notificationEndpoint string, insTransactionSimulationRequest InsTransactionSimulationRequest, localVarOptionals *INSTransactionOpts) (SuccessAcknowledgement, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  SuccessAcknowledgement
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/Transaction"
	headerParams := make(map[string]string)
	queryParams := _neturl.Values{}
	formParams := _neturl.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	headerParams["NotificationEndpoint"] = parameterToString(notificationEndpoint, "")
	if localVarOptionals != nil && localVarOptionals.EndpointCredential.IsSet() {
		headerParams["EndpointCredential"] = parameterToString(localVarOptionals.EndpointCredential.Value(), "")
	}
	// body params
	postBody = &insTransactionSimulationRequest
	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, formFileName, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	res, err := a.client.callAPI(r)
	if err != nil || res == nil {
		return returnValue, res, err
	}

	body, err := _ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return returnValue, res, err
	}

	if res.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  body,
			error: res.Status,
		}
		if res.StatusCode == 401 {
			var v ErrorAcknowledgement
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, res, newErr
			}
			newErr.model = v
			return returnValue, res, newErr
		}
		return returnValue, res, newErr
	}

	err = a.client.decode(&returnValue, body, res.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  body,
			error: err.Error(),
		}
		return returnValue, res, newErr
	}

	return returnValue, res, nil
}
