package repositories

import (
	"database/sql"
	"errors"
	"log"
	"super-payment-kunn/models"
)

type InvoiceRepository interface {
	SelectAll() ([]*models.Invoice, error)
	Store(invoice *models.Invoice) error
}

var dbOnMemory = make(map[int]*models.Invoice, 0)

type invoiceRepositoryOnMemory struct {
	db map[int]*models.Invoice
}

func NewInvoiceRepositoryOnMemory() InvoiceRepository {
	return &invoiceRepositoryOnMemory{db: dbOnMemory}
}

func (r *invoiceRepositoryOnMemory) SelectAll() ([]*models.Invoice, error) {
	invoices := make([]*models.Invoice, 0, len(r.db))
	for _, invoice := range r.db {
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (r *invoiceRepositoryOnMemory) Store(invoice *models.Invoice) error {
	if _, exists := r.db[invoice.ID]; exists {
		return errors.New("invoice ID is already exists")
	}
	r.db[invoice.ID] = invoice
	return nil
}

type invoiceRepositoryOnRDB struct {
	db *sql.DB
}

func NewInvoiceRepositoryOnRDB(db *sql.DB) InvoiceRepository {
	return &invoiceRepositoryOnRDB{db: db}
}

func (r *invoiceRepositoryOnRDB) SelectAll() ([]*models.Invoice, error) {
	query := "SELECT id, amount, due_date, issue_date, total_amount, fee_rate, tax_rate, status FROM invoice"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatalf("Error retriving invoice: %v", err)
		return nil, err
	}
	defer rows.Close()

	var invoices []*models.Invoice
	for rows.Next() {
		var invoice models.Invoice
		if err := rows.Scan(&invoice.ID, &invoice.Amount, &invoice.DueDate, &invoice.IssueDate, &invoice.TotalAmount, &invoice.FeeRate, &invoice.TaxRate, &invoice.Status); err != nil {
			log.Fatalf("Error scaning invoice: %v", err)
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func (r *invoiceRepositoryOnRDB) Store(invoice *models.Invoice) error {
	query := "INSERT INTO invoice (amount, due_date, issue_date, total_amount, fee_rate, tax_rate, status) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, invoice.Amount, invoice.DueDate, invoice.IssueDate, invoice.TotalAmount, invoice.FeeRate, invoice.TaxRate, invoice.Status)
	if err != nil {
		log.Fatalf("Error Saving invoice: %v", err)
	}
	return nil
}
