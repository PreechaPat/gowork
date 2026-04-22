package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if len(response) != 7 {
		t.Errorf("expected 7 users, got %v", len(response))
	}
}

func TestGetUserHandler(t *testing.T) {
	// Note: r.PathValue requires Go 1.22+ and typically needs the mux to be involved
	// to populate the path values. For a unit test of the handler itself using PathValue,
	// we'd often use a mux in the test.
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/user/{name}", GetUserHandler)

	t.Run("Existing user", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/user/alice", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected 200, got %v", status)
		}
	})

	t.Run("Non-existing user", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/user/unknown", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("expected 404, got %v", status)
		}
	})
}
