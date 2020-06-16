package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Authentication
// An Access Token will be used to access, authorize and authenticate your interactions
// with the Co-op Bank’s Open Banking APIs.
// It will be generated using a Consumer Key and a Consumer Secret linked to your account.
// The Access Token will expire after a specified duration of time, after which regeneration
// of the Access Token will be required.
// To generate or re-generate an Access Token, your sign-in username and password will be required.
// It is very important to keep your API credentials safe, as they can be used to access
// your account, make changes to your account and carry out transactions.
// It is also important to note that once you generate a new set of API credentials you must
// update all your applications with the new API details for your applications to continue working.
// Note that the Sandbox and Live environments have unique Consumer Key and Consumer Secret which
// are NOT be interchangeable.

// The following headers will be required for each request:
// AuthenticationToken(String & Required)
// This is your auth key from the platform

// Generating Access Tokens
// The following cURL command shows how to generate an access token using the Password Grant type.
// curl -k -d "grant_type=password&username=Username&password=Password" \
//	    -H "Authorization: Basic Base64(consumer-key:consumer-secret)" \
//	     https://developer.co-opbank.co.ke:8243/token
//
// In a similar manner, you can generate an access token using the
// Client Credentials grant type with the following cURL command.
// curl -k -d "grant_type=client_credentials" \
//	    -H "Authorization: Basic Base64(consumer-key:consumer-secret)" \
//	     https://developer.co-opbank.co.ke:8243/token
// Token has a validity period of 3600 seconds.
// And the token has ( am_application_scope,default ) scopes.

type (
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
)

func GetAccessToken(ctx context.Context, ConsumerKey, ConsumerSecret string) (*TokenResponse, error) {

	buf := bytes.NewBuffer([]byte("grant_type=client_credentials"))

	req, err := http.NewRequestWithContext(ctx, "POST", "https://developer.co-opbank.co.ke:8243/token", buf)
	if err != nil {
		return &TokenResponse{}, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(ConsumerKey, ConsumerSecret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not load default HTTP client: %w", err)
	}

	//  We're done reading from response body, lets close it.
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			err := fmt.Errorf("could not close response body: %w", err)
			fmt.Println(err.Error())
		}
	}()

	// Read data from response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("oauth2: cannot fetch token: %v", err)
		fmt.Println(err.Error())
	}

	var token = TokenResponse{}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &token); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	return &token, nil
}
