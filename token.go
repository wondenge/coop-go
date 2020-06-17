package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

// GetAccessToken returns struct of TokenResponse
// No need to call SetAccessToken to apply new access token for current Client
// Endpoint: POST https://developer.co-opbank.co.ke:8243/token
func (c *APIClient) GetAccessToken(ctx context.Context) (*TokenResponse, error) {

	buf := bytes.NewBuffer([]byte("grant_type=client_credentials"))

	req, err := http.NewRequestWithContext(ctx, "POST", "https://developer.co-opbank.co.ke:8243/token", buf)
	if err != nil {
		return &TokenResponse{}, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Makes a request to the API using key:secret basic auth
	req.SetBasicAuth(c.ConsumerKey, c.ConsumerSecret)

	resp, err := c.Client.Do(req)
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

	// Set Token fur current Client
	if token.AccessToken != "" {
		c.Token = &token
		c.tokenExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	return &token, nil
}
