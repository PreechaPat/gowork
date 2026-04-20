package handler

import (
	"fmt"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Query().Get("message")

	if message == "" {
		http.Error(w, "Missing 'message' parameter", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s\n", message)
}
