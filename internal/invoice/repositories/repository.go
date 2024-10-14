package repositories

type InvoiceRepository interface {
	SelectAll() []map[string]string
	Store(invoice map[string]string)
}

var dbOnMemory = make([]map[string]string, 0)

type invoiceRepository struct {
	db *[]map[string]string
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{db: &dbOnMemory}
}

func (r *invoiceRepository) SelectAll() []map[string]string {
	return *r.db
}

func (r *invoiceRepository) Store(invoice map[string]string) {
	*r.db = append(*r.db, invoice)
}
