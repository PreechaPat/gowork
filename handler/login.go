package handler

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:        "auth_token",
		Value:       "nothing",
		Quoted:      false,
		Path:        "/",
		Domain:      "",
		MaxAge:      3600,
		Secure:      false,
		HttpOnly:    true,
		SameSite:    http.SameSiteStrictMode,
		Partitioned: false,
		Raw:         "",
		Unparsed:    []string{},
	}

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Logged in")
}
