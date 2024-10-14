package main

import (
	"log"
	"net/http"
	"super-payment-kunn/internal/invoice/handler"
)

func main() {
	mux := http.NewServeMux()

	apiInvoiceHandler := handler.NewInvoiceHandler()

	mux.HandleFunc("/api/invoices", apiInvoiceHandler.HandleApiInvoice)

	log.Println("Starting web server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}
