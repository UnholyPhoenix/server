package calculator

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Request definition
type Request struct {
	Expression string
}

// Response definition
type Response struct {
	Result     float64
	Error      error
	Expression string
}

// MakeEndpoint creates endpoint for greeter
func MakeEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)

		result, err := s.Evaluate(req.Expression)
		return Response{
			Result:     result,
			Error:      err,
			Expression: req.Expression,
		}, nil
	}
}
