package calculator

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// ValidateMiddleware is a validator middleware for service
func ValidateMiddleware() Middleware {
	return func(next Service) Service {
		return validateMiddleware{next}
	}
}

type validateMiddleware struct {
	next Service
}

func (mw validateMiddleware) Evaluate(expression string, kv *KV) (float64, error) {
	if expression == "" {
		return 0, nil
	}

	return mw.next.Evaluate(expression, kv)
}

// ServiceLoggingMiddleware is a validator middleware for service
func ServiceLoggingMiddleware(log log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{log, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) Evaluate(expression string, kv *KV) (float64, error) {
	result, err := mw.next.Evaluate(expression, kv)
	mw.logger.Log("method", "Evaluate", "expression", expression, "result", result)
	return result, err
}

// EndpointLoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func EndpointLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}
