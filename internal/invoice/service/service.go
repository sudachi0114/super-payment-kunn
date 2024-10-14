package service

import (
	"super-payment-kunn/internal/invoice/repositories"
	"super-payment-kunn/models"
	"time"
)

// TODO: RDS を利用するようになったら AUTO_INCREMENT にする
var ID = 0

type InvoiceService interface {
	GetInvoices() []*models.Invoice
	PostInvoice(invoice *models.InvoiceRequestJson) (*models.Invoice, error)
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

func (s *invoiceService) PostInvoice(req *models.InvoiceRequestJson) (*models.Invoice, error) {
	issueDate := time.Now()

	// FIXME: 変更があるかもしれないので、外から差し込めるような形式にしたい
	const feeRate = 0.04
	const taxRate = 0.1

	totalAmount := float64(req.Amount) + (float64(req.Amount) * feeRate * (1.0 + taxRate))

	invoice := &models.Invoice{
		ID:          ID,
		Amount:      req.Amount,
		DueDate:     time.Unix(int64(req.DueDate), 0),
		IssueDate:   issueDate,
		TotalAmount: totalAmount,
		FeeRate:     feeRate,
		TaxRate:     taxRate,
		Status:      models.StatusUnprocessed,
	}

	err := s.repository.Store(invoice)
	if err != nil {
		return nil, err
	}

	ID++
	return invoice, nil
}
