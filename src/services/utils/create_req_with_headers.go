package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func CreateReqWithHeaders(body []byte, url string, visitorID *string) (
	*http.Request, error,
) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		err := fmt.Errorf("error creating request: %v", err)
		return nil, err
	}
	if visitorID != nil {
		req.Header.Set("X-Goog-Visitor-Id", *visitorID)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://music.youtube.com")
	req.Header.Set(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/"+
			"537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
	)

	return req, nil
}
