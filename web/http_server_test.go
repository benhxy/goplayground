package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandlerFriend(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"http://localhost:8080/joe@benhu.dev", "friend joe"},
		{"http://localhost:8080/kim@golang.org", "kim@golang.org"},
	}

	for _, test := range tests {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, test.input, nil)

		helloHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected StatusOK; actual %v", rec.Code)
		}

		responseBody := rec.Body.String()
		if !strings.Contains(responseBody, test.output) {
			t.Errorf("invalid response body: %v", responseBody)
		}
	}

}
