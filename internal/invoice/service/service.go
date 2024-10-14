package service

import "super-payment-kunn/internal/invoice/repositories"

type InvoiceService interface {
	GetInvoices() []map[string]string
	PostInvoice(invoice map[string]string)
}

type invoiceService struct {
	repository repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) InvoiceService {
	return &invoiceService{repository: repo}
}

func (s *invoiceService) GetInvoices() []map[string]string {
	return s.repository.SelectAll()
}

func (s *invoiceService) PostInvoice(invoice map[string]string) {
	s.repository.Store(invoice)
}
