package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Name string `json:"name"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    req.Name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged in")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Logged out")
}

func MeHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session_token")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"name": ""}`))
		return
	}

	m := make(map[string]string)
	m["name"] = cookie.Value
	json.NewEncoder(w).Encode(m)
}
