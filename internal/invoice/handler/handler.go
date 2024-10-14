package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"super-payment-kunn/internal/invoice/dto"
	"super-payment-kunn/internal/invoice/service"
)

type InvoiceHandler struct {
	service service.InvoiceService
}

func NewInvoiceHandler(service service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
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
	invoices, err := h.service.GetInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(invoices)
}

func (h *InvoiceHandler) PostInvoice(w http.ResponseWriter, r *http.Request) {
	var req *dto.InvoiceRequestJson
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
	}

	err := h.service.PostInvoice(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
