package service

import (
	"super-payment-kunn/internal/invoice/repositories"
	"super-payment-kunn/models"
)

type InvoiceService interface {
	GetInvoices() []*models.Invoice
	PostInvoice(invoice *models.Invoice)
}

type invoiceService struct {
	repository repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) InvoiceService {
	return &invoiceService{repository: repo}
}

func (s *invoiceService) GetInvoices() []*models.Invoice {
	return s.repository.SelectAll()
}

func (s *invoiceService) PostInvoice(invoice *models.Invoice) {
	s.repository.Store(invoice)
}
