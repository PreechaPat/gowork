package main

import (
	"log"
	"net/http"
	"os"
	"worker/handler"
)

func NewServer(logger *log.Logger) *http.ServeMux {
	return http.NewServeMux()
}

func addRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"up"}`))
	})

	mux.HandleFunc("GET /echo", handler.EchoHandler)
}

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := NewServer(logger)

	addRoutes(
		mux,
	)

	logger.Print("start server at 8080")

	http.ListenAndServe(":8080", mux)
}
