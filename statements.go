package connectapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/coop-go/gen/connect"
	"io/ioutil"
	"net/http"
)

// 2. Account Full Statement Enquiry API will enable you to enquire about your own
// Co-operative Bank accounts' full statement for the specified account number,
// start date and end date.
//
// Post an Account Full Statement Enquiry Request
func (c *APIClient) AccountFullStatement(ctx context.Context, p *connect.AccountFullStatementPayload) (res *connect.AccountFullStatementSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, "/Enquiry/FullStatement/Account/1.0.0"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+c.Token.AccessToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not load default HTTP client: %w", err)
	}

	if err := c.logger.Log("info", fmt.Sprintf("connect.AccountFullStatement")); err != nil {
		err := fmt.Errorf("could not log to stdout: %w", err)
		fmt.Println(err.Error())
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

	res = &connect.AccountFullStatementSuccessResponse{}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	return res, err
}

// 3. Account Mini Statement Enquiry API will enable you to enquire about your
// own Co-operative Bank accounts' Mini statement for the specified account number.
//
// Post an Account Mini Statement Enquiry Request
func (c *APIClient) AccountMiniStatement(ctx context.Context, p *connect.AccountMiniStatementPayload) (res *connect.AccountMiniStatementSuccessResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, "/MiniStatement/Account/1.0.0"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+c.Token.AccessToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not load default HTTP client: %w", err)
	}

	if err := c.logger.Log("info", fmt.Sprintf("connect.AccountMiniStatement")); err != nil {
		err := fmt.Errorf("could not log to stdout: %w", err)
		fmt.Println(err.Error())
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

	res = &connect.AccountMiniStatementSuccessResponse{}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	return res, err

}
