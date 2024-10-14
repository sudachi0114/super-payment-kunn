package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleApiInvoice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET /api/invoices is not implemented")
		return
	case http.MethodPost:
		fmt.Fprintln(w, "POST /api/invoices is not implemented")
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/invoices", handleApiInvoice)

	log.Println("Starting web server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}
