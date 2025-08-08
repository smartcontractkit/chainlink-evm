package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type contextKey string

const responseLimitCtxKey contextKey = "responseLimitCtxKey"

// LimitedTransport wraps an http.RoundTripper and limits the size of the response body. Limit is set via context using WithResponseSizeLimit
type LimitedTransport struct {
	// RoundTripper is the underlying http.RoundTripper to use for the actual request.
	// This will typically be http.DefaultTransport or a custom *http.Transport.
	RoundTripper http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface for LimitedTransport.
func (t *LimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Perform the actual HTTP request using the underlying RoundTripper.
	resp, err := t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// If the response body is not nil, wrap it with an io.limitReader.
	// This will ensure that only up to MaxResponseSize bytes can be read.
	respLimit := GetResponseSizeLimit(req.Context())
	if resp.Body != nil && respLimit > 0 {
		resp.Body = limitReader(resp.Body, int64(respLimit))
	}

	return resp, nil
}

// WithResponseSizeLimit - sets a limit on the size of the response body for HTTP requests made with the LimitedTransport.
func WithResponseSizeLimit(ctx context.Context, limit uint32) context.Context {
	if limit > 0 {
		return context.WithValue(ctx, responseLimitCtxKey, limit)
	}

	return ctx
}

func GetResponseSizeLimit(ctx context.Context) uint32 {
	limit, ok := ctx.Value(responseLimitCtxKey).(uint32)
	if !ok {
		return 0
	}
	return limit
}

var errResponseTooLarge = errors.New("response is too large")

// limitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
// The underlying implementation is a *limitedReader.
func limitReader(r io.ReadCloser, n int64) *limitedReader {
	return &limitedReader{R: r, N: n, Limit: n}
}

// A limitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
type limitedReader struct {
	R     io.ReadCloser // underlying reader
	N     int64         // max bytes remaining
	Limit int64         // original limit for error reporting
}

func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, fmt.Errorf("reached read limit of %d bytes: %w", l.Limit, errResponseTooLarge)
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func (l *limitedReader) Close() error {
	return l.R.Close()
}
