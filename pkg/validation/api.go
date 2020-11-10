/*
 * AccountValidation
 *
 * Account Validation Enquiry API will enable you to validate a Co-operative Bank account number
 *
 * API version: 1.0.0
 *
 *
 */

package validation

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// AccountValidationService AccountValidationAPI service
type AccountValidationService service

/*
Account Post an Account Validation Enquiry Request
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountValidationRequest Account Validation Enquiry Request Body
@return AccountValidationSuccessResponse
*/
func (a *AccountValidationService) Account(ctx _context.Context, accountValidationRequest AccountValidationRequest) (AccountValidationSuccessResponse, *_nethttp.Response, error) {
	var (
		httpMethod   = _nethttp.MethodPost
		postBody     interface{}
		formFileName string
		fileName     string
		fileBytes    []byte
		returnValue  AccountValidationSuccessResponse
	)

	// create path and map variables
	path := a.client.cfg.BasePath + "/"
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
	// body params
	postBody = &accountValidationRequest
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
		if res.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, body, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, res, newErr
			}
			newErr.model = v
			return returnValue, res, newErr
		}
		if res.StatusCode == 401 {
			var v MissingCredentials
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
