package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Starting web server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}
