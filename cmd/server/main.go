package main

import (
	"log"

	"github.com/karanp3/proglog/internal/server"
)

func main() {
	addr := ":8080"
	srv := server.NewHTTPServer(addr)
	log.Printf("Starting server on port%s", addr)
	log.Fatal(srv.ListenAndServe())
}
