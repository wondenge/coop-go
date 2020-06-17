package connectapi

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/wondenge/coop-go/gen/connect"
	"reflect"
	"testing"
	"time"
)

func TestAPIClient_SendToMPesa(t *testing.T) {
	type fields struct {
		Client         *http.Client
		ConsumerKey    string
		ConsumerSecret string
		APIBase        string
		Token          *TokenResponse
		tokenExpiresAt time.Time
		logger         log.Logger
	}
	type args struct {
		ctx context.Context
		p   *connect.SendToMpesaTransactionRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *connect.SuccessAcknowledgement
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &APIClient{
				Client:         tt.fields.Client,
				ConsumerKey:    tt.fields.ConsumerKey,
				ConsumerSecret: tt.fields.ConsumerSecret,
				APIBase:        tt.fields.APIBase,
				Token:          tt.fields.Token,
				tokenExpiresAt: tt.fields.tokenExpiresAt,
				logger:         tt.fields.logger,
			}
			gotRes, err := c.SendToMPesa(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendToMPesa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SendToMPesa() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}