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
func MakeEndpoint(s Service, kv *KV) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)

		if req.Expression != "" {
			result, err := s.Evaluate(req.Expression, kv)

			if err == nil {
				// Add to memory storage
				kv.Add(result)
			}

			// return resoponse
			return Response{
				Result:     result,
				Error:      err,
				Expression: req.Expression,
			}, nil
		}

		// return resoponse
		return Response{
			Result:     0,
			Error:      nil,
			Expression: "",
		}, nil

	}
}
