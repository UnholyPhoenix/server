package calculator

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
)

type mockService struct{}

func (s mockService) Evaluate(expression string) (float64, error) {
	return 0.0, nil
}

func TestEndpoint(t *testing.T) {
	service := mockService{}
	endpoint := MakeEndpoint(service)

	tests := []struct {
		name        string
		requestExpr string
		result      float64
		err         error
	}{
		{
			"Expr not set",
			"",
			0.0,
			nil,
		},
		{
			"Expr set",
			"10 + 5",
			0.0,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Request{
				Expression: tt.requestExpr,
			}

			resp, err := endpoint(context.Background(), req)
			if err != nil {
				t.Fatalf("expect error to be nil, got %v", err)
			}

			response := resp.(Response)

			if expect, got := tt.result, response.Result; expect != got {
				t.Errorf("expected '%v', got '%v'", expect, got)
			}
		})
	}
}

func TestEndpointLogging(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogfmtLogger(buf)

	service := mockService{}
	endpoint := EndpointLoggingMiddleware(logger)(MakeEndpoint(service))

	req := Request{}

	_, err := endpoint(context.Background(), req)
	if err != nil {
		t.Fatalf("expect error to be nil, got %v", err)
	}

	if expect, got := "transport_error=null took=", buf.String(); !strings.HasPrefix(got, expect) {
		t.Errorf("expected '%v', got '%v'", expect, got)
	}
}
