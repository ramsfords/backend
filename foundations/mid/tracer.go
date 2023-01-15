package mid

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/echo/v5"
	"go.opentelemetry.io/otel/trace"
)

// ctxKey represents the type of value for the context key.
type ctxKey string

// key is how request values are stored/retrieved.
const key ctxKey = "1"

// Values represent state for each request.
type Values struct {
	TraceID    string    `json:"trace_id"`
	Now        time.Time `json:"now"`
	StatusCode int       `json:"status_code"`
}

func (va Values) String() string {
	return fmt.Sprintf("trace_id:%s,now:%v,status_code:%d", va.TraceID, va.Now.Local().String(), va.StatusCode)
}

// GetValues returns the values from the context.
func GetValues(ctx context.Context) (*Values, error) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return nil, errors.New("web value missing from context")
	}
	return v, nil
}

// GetTraceID returns the trace id from the context.
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceID
}

// SetStatusCode sets the status code back into the context.
func SetStatusCode(ctx context.Context, statusCode int) error {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return errors.New("web value missing from context")
	}
	v.StatusCode = statusCode
	return nil
}

func TraceInjector() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Pull the context from request
		c := ctx.Request().Context()
		// Capture the parent request span from the context.
		span := trace.SpanFromContext(c)
		// Set the context with the required values to
		// process the request.
		v := Values{
			TraceID: span.SpanContext().TraceID().String(),
			Now:     time.Now(),
		}
		ctx.Set("trace_id", &v)
		return nil
	}
}
