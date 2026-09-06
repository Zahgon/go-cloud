package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gocloud.dev/server"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := server.New(r, nil)
	log.Printf("Listening on port %s", port)
	log.Fatal(srv.ListenAndServe(fmt.Sprintf(":%s", port)))
}

func handle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
