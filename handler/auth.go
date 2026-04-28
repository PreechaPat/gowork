package handler

import (
	"fmt"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No cookie found!", http.StatusUnauthorized)
		}
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Found the cookie! Value is: %s", cookie.Value)
}
