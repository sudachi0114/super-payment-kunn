package main

import (
	"log"
	"net/http"
	"super-payment-kunn/internal/db"
	"super-payment-kunn/internal/invoice/handler"
	"super-payment-kunn/internal/invoice/repositories"
	"super-payment-kunn/internal/invoice/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn := db.InitRDB()
	defer dbConn.Close()

	mux := http.NewServeMux()

	invoiceRepo := repositories.NewInvoiceRepositoryOnRDB(dbConn)
	invoiceService := service.NewInvoiceService(invoiceRepo)
	apiInvoiceHandler := handler.NewInvoiceHandler(invoiceService)

	mux.HandleFunc("/api/invoices", apiInvoiceHandler.HandleApiInvoice)

	log.Println("Starting web server on :8080")
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}
