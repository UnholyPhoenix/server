package calculator

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestServiceValidatorMiddleware(t *testing.T) {
	s := ValidateMiddleware()(New())
	kv := NewMemoryStorage()

	result, err := s.Evaluate("", kv)
	if expect, got, gotError := 0.00, result, err; expect != got {
		t.Errorf("expected '%v', got '%v', error '%v'", expect, got, gotError)
	}

	result, err = s.Evaluate("10 + 5", kv)
	if expect, got, gotError := 15.00, result, err; expect != got {
		t.Errorf("expected '%v', got '%v', error '%v'", expect, got, gotError)
	}
}

func TestServiceLoggerMiddleare(t *testing.T) {

	buf := new(bytes.Buffer)
	logger := log.NewLogfmtLogger(buf)

	s := ServiceLoggingMiddleware(logger)(New())
	kv := NewMemoryStorage()

	s.Evaluate("10 + 5", kv)

	if expect, got := "method=Evaluate expression=\"10 + 5\" result=15\n", buf.String(); expect != got {
		t.Errorf("expected '%v', got '%v'", expect, got)
	}
}
