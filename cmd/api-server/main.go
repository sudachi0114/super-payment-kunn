package main

import (
	"log"
	"net/http"
	"super-payment-kunn/internal/invoice/handler"
	"super-payment-kunn/internal/invoice/repositories"
	"super-payment-kunn/internal/invoice/service"
)

func main() {
	mux := http.NewServeMux()

	invoiceRepo := repositories.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepo)
	apiInvoiceHandler := handler.NewInvoiceHandler(invoiceService)

	mux.HandleFunc("/api/invoices", apiInvoiceHandler.HandleApiInvoice)

	log.Println("Starting web server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}
