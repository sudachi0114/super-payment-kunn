package dto

type InvoiceRequestJson struct {
	Amount  int    `json:"amount"`
	DueDate string `json:"due_date"`
}
