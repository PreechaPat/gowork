package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMeHandler_NoCookie(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/me", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MeHandler)

	// This should not panic
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	expected := `{"name": ""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMeHandler_WithCookie(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/me", nil)
	if err != nil {
		t.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:  "session_token",
		Value: "testuser",
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"name":"testuser"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}
