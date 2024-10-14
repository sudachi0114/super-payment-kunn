package handler

import (
	"fmt"
	"net/http"
)

type InvoiceHandler struct {
}

func NewInvoiceHandler() *InvoiceHandler {
	return &InvoiceHandler{}
}

func (h *InvoiceHandler) HandleApiInvoice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetInvoices(w, r)
		return
	case http.MethodPost:
		h.PostInvoice(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
		return
	}
}

func (h *InvoiceHandler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET /api/invoices is not implemented")
}

func (h *InvoiceHandler) PostInvoice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST /api/invoices is not implemented")
}
