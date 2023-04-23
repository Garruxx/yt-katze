package tests

import (
	"io"
	"katze/src/services/models"
	"katze/src/services/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	// Create a new test server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Check that the request method is POST
		assert.Equal(t, http.MethodPost, req.Method)

		// Check that the request URL is correct
		assert.Equal(t, "https://music.youtube.com/test-url-path?test-query=1", req.URL.String())

		// Check that the request headers are correct
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, "https://music.youtube.com", req.Header.Get("Origin"))
		assert.Equal(t, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36", req.Header.Get("User-Agent"))

		// Check that the request body is correct
		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %v", err)
		}
		assert.Equal(t, `{"context":{"client":{"clientName":"WEB_REMIX","clientVersion":"1.20230306.01.00"}},"params":"test-params","query":"test-query","browseId":"test-browse-id","videoId":"videoIdTest"}`, string(body))

		// Send a response
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"key":"value"}`))
	}))
	defer server.Close()

	// Create a new request with test data
	request := models.Request{
		UrlPath:       "/test-url-path",
		UrlQueries:    "test-query=1",
		Params:        "test-params",
		Query:         "test-query",
		BrowseID:      "test-browse-id",
		VideoID:       "videoIdTest",
		GoogVisitorID: nil,
	}

	// Call the function being tested
	_, err := utils.Request(request)
	assert.NoError(t, err)
}
