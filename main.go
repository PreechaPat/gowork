package main

import (
	"flag"
	"fmt"
	"gowork/handler"
	mw "gowork/middleware"
	"log"
	"net/http"
	"os"
)

// Just without file server
func NewDevServer(logger *log.Logger) *http.ServeMux {

	mux := http.NewServeMux()
	addRoutes(logger, mux)
	return mux
}

func NewProdServer(logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	addRoutes(logger, mux)

	fileServer := http.FileServer(http.Dir("./dist"))
	mux.Handle("/", fileServer)

	return mux
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}

func addRoutes(logger *log.Logger, mux *http.ServeMux) {

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Hello is called")
		w.Write([]byte("Hello world"))
	})

	mux.Handle("GET /api/health", mw.NewChain(mw.Logger(logger)).Then(http.HandlerFunc(HealthHandler)))

	mux.HandleFunc("GET /api/auth", handler.AuthHandler)
	mux.HandleFunc("GET /api/echo", handler.EchoHandler)
	mux.HandleFunc("GET /api/users", handler.ListUsersHandler)
	mux.HandleFunc("GET /api/user/{name}", handler.GetUserHandler)

}

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := NewProdServer(logger)

	logger.Printf("start server at %s", *port)

	http.ListenAndServe(fmt.Sprintf(":%s", *port), mux)
}
