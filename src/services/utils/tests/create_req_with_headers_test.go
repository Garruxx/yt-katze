package tests

import (
	"katze/src/services/utils"
	"testing"
)

func TestCreateReqWithHeaders(t *testing.T) {
	// Prueba con visitorID
	visitorID := "12345"
	body := []byte(`{"foo":"bar"}`)
	req, err := utils.CreateReqWithHeaders(body, "https://example.com", &visitorID)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	if req.Method != "POST" {
		t.Errorf("Expected POST method, got %v", req.Method)
	}
	if req.URL.String() != "https://example.com" {
		t.Errorf("Expected URL https://example.com, got %v", req.URL.String())
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %v", req.Header.Get("Content-Type"))
	}
	if req.Header.Get("X-Goog-Visitor-Id") != "12345" {
		t.Errorf("Expected X-Goog-Visitor-Id 12345, got %v", req.Header.Get("X-Goog-Visitor-Id"))
	}
	if req.Header.Get("Origin") != "https://music.youtube.com" {
		t.Errorf("Expected Origin https://music.youtube.com, got %v", req.Header.Get("Origin"))
	}
	if req.Header.Get("User-Agent") != "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36" {
		t.Errorf("Expected User-Agent Mozilla/5.0 ..., got %v", req.Header.Get("User-Agent"))
	}

	// Prueba sin visitorID
	body = []byte(`{"foo":"bar"}`)
	req, err = utils.CreateReqWithHeaders(body, "https://example.com", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	if req.Header.Get("X-Goog-Visitor-Id") != "" {
		t.Errorf("Expected X-Goog-Visitor-Id '', got %v", req.Header.Get("X-Goog-Visitor-Id"))
	}

	// Prueba con un cuerpo vacío
	body = []byte("")
	_, err = utils.CreateReqWithHeaders(body, "https://example.com", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
}
