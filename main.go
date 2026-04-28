package main

import (
	"log"
	"net/http"
	"os"
	"worker/handler"
)

// Just without file server
func NewDevServer(logger *log.Logger) *http.ServeMux {

	mux := http.NewServeMux()
	addRoutes(mux)
	return mux
}

func NewProdServer(logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	addRoutes(mux)

	fileServer := http.FileServer(http.Dir("./dist"))
	mux.Handle("/", fileServer)

	return mux
}

func addRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Hello is called")
		w.Write([]byte("Hello world"))
	})

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"up"}`))
	})

	mux.HandleFunc("GET /api/auth", handler.AuthHandler)
	mux.HandleFunc("GET /api/echo", handler.EchoHandler)
	mux.HandleFunc("GET /api/users", handler.ListUsersHandler)
	mux.HandleFunc("GET /api/user/{name}", handler.GetUserHandler)

}

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := NewProdServer(logger)

	logger.Print("start server at 8080")

	http.ListenAndServe(":8080", mux)
}
