package testutils

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type HTTPServer struct {
	t        *testing.T
	lggr     logger.SugaredLogger
	server   *httptest.Server
	callback JSONRPCHandler
	chainID  *big.Int
}

// NewHTTPServer starts an HTTP server which invokes callback for each message received.
// If chainID is set, then eth_chainId calls will be automatically handled.
func NewHTTPServer(t *testing.T, chainID *big.Int, callback JSONRPCHandler) *HTTPServer {
	ts := &HTTPServer{
		t:        t,
		lggr:     logger.Sugared(logger.Test(t)),
		chainID:  chainID,
		callback: callback,
	}
	ts.server = httptest.NewServer(http.HandlerFunc(ts.ServeHTTP))
	t.Cleanup(ts.Close)
	return ts
}

func (ts *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ts.lggr.Errorf("Error reading body: %v", err)
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	ts.lggr.Debugf("Received message: %s", string(data))

	req := gjson.ParseBytes(data)

	if req.IsArray() { // Handle batch request
		ts.lggr.Debug("Received batch request")
		var responses []string
		for i, reqElem := range req.Array() {
			var response string
			response, _, err = handleRequest(ts.lggr, ts.chainID, ts.callback, reqElem)
			if err != nil {
				http.Error(w, fmt.Errorf("failed to handle elem %d of batch request: %w", i, err).Error(), http.StatusInternalServerError)
				return
			}
			responses = append(responses, response)
		}

		ts.writeJSON(w, fmt.Sprintf("[%s]", strings.Join(responses, ",")))
		return
	}
	// Handle single request
	response, asyncResponse, err := handleRequest(ts.lggr, ts.chainID, ts.callback, req)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to handle request: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	if asyncResponse != "" {
		panic("async response not supported in HTTP server")
	}

	ts.writeJSON(w, response)
}

func (ts *HTTPServer) writeJSON(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(data))
	if err != nil {
		err = fmt.Errorf("failed to write response: %w", err)
		ts.lggr.Errorf("Error writing response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ts *HTTPServer) URL() *url.URL {
	u, err := url.Parse(ts.server.URL)
	require.NoError(ts.t, err, "Failed to parse url")
	return u
}

func (ts *HTTPServer) Close() {
	ts.server.Close()
}
