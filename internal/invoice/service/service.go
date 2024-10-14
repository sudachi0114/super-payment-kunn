package service

import (
	"errors"
	"super-payment-kunn/internal/invoice/dto"
	"super-payment-kunn/internal/invoice/repositories"
	"super-payment-kunn/models"
	"time"
)

// TODO: RDS を利用するようになったら AUTO_INCREMENT にする
var ID = 0

type InvoiceService interface {
	GetInvoices() ([]*models.Invoice, error)
	PostInvoice(invoice *dto.InvoiceRequestJson) error
}

type invoiceService struct {
	repository repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) InvoiceService {
	return &invoiceService{repository: repo}
}

func (s *invoiceService) GetInvoices() ([]*models.Invoice, error) {
	return s.repository.SelectAll()
}

func (s *invoiceService) PostInvoice(req *dto.InvoiceRequestJson) error {
	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		return errors.New("Invalid due date format")
	}
	issueDate := time.Now()

	// FIXME: 変更があるかもしれないので、外から差し込めるような形式にしたい
	const feeRate = 0.04
	const taxRate = 0.1

	totalAmount := float64(req.Amount) + (float64(req.Amount) * feeRate * (1.0 + taxRate))

	// FIXME: 本当はドメインで変換するべきかもしれない
	invoice := &models.Invoice{
		ID:          ID,
		Amount:      req.Amount,
		DueDate:     dueDate,
		IssueDate:   issueDate,
		TotalAmount: totalAmount,
		FeeRate:     feeRate,
		TaxRate:     taxRate,
		Status:      models.StatusUnprocessed,
	}

	err = s.repository.Store(invoice)
	if err != nil {
		return err
	}

	ID++
	return nil
}
