package main

import (
	"fmt"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	fmt.Printf("Starting server on port: %s", port)
	server.ListenAndServe()
}
