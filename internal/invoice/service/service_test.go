package service_test

import (
	"super-payment-kunn/internal/invoice/repositories"
	"super-payment-kunn/internal/invoice/service"
	"super-payment-kunn/models"
	"testing"
	"time"
)

func TestGetInvoices(t *testing.T) {
	mockRepo := repositories.NewInvoiceRepositoryOnMemory()
	mockRepo.Store(&models.Invoice{ID: 1, Amount: 100, DueDate: time.Now(), IssueDate: time.Now(), TotalAmount: 104.4, FeeRate: 0.04, TaxRate: 0.1, Status: models.StatusUnprocessed})

	service := service.NewInvoiceService(mockRepo)
	invoices, err := service.GetInvoices()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(invoices) != 1 {
		t.Errorf("expected invoice count is 2, but got: %d", len(invoices))
	}

	if invoices[0].ID != 1 {
		t.Errorf("expected ID: 1, bot got: %d", invoices[0].ID)
	}
}
