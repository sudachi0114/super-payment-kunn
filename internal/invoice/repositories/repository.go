package repositories

import (
	"errors"
	"super-payment-kunn/models"
)

type InvoiceRepository interface {
	SelectAll() []*models.Invoice
	Store(invoice *models.Invoice) error
}

var dbOnMemory = make(map[int]*models.Invoice, 0)

type invoiceRepository struct {
	db map[int]*models.Invoice
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{db: dbOnMemory}
}

func (r *invoiceRepository) SelectAll() []*models.Invoice {
	invoices := make([]*models.Invoice, 0, len(r.db))
	for _, invoice := range r.db {
		invoices = append(invoices, invoice)
	}
	return invoices
}

func (r *invoiceRepository) Store(invoice *models.Invoice) error {
	if _, exists := r.db[invoice.ID]; exists {
		return errors.New("invoice ID is already exists")
	}
	r.db[invoice.ID] = invoice
	return nil
}
