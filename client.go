package connectapi

import (
	"github.com/go-kit/kit/log"
	"net/http"
	"time"
)

const (

	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = "http://developer.co-opbank.co.ke:8280"

	// APIBaseLive points to the live version of the API
	APIBaseLive = "https://developer.co-opbank.co.ke:8243"
)

type (
	APIClient struct {
		Client         *http.Client
		ConsumerKey    string
		ConsumerSecret string
		APIBase        string
		Token          *TokenResponse
		tokenExpiresAt time.Time
		logger         log.Logger // log the requests.
	}
)
