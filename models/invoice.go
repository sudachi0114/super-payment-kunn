package models

import "time"

type InvoiceRequestJson struct {
	Amount  int `json:"amount"`
	DueDate int `json:"due_date"`
}

type InvoiceStatus string

const (
	StatusUnprocessed InvoiceStatus = "未処理"
	StatusProcessing  InvoiceStatus = "処理中"
	StatusPaid        InvoiceStatus = "支払い済み"
	StatusError       InvoiceStatus = "エラー"
)

// 請求書データ (企業・取引先に紐づく)
type Invoice struct {
	ID          int
	Amount      int           // 支払金額
	DueDate     time.Time     // 支払期日
	IssueDate   time.Time     // 発行日
	TotalAmount float64       // 請求金額
	FeeRate     float32       // 手数料率 -> 手数料
	TaxRate     float32       // 消費税率 → 消費税
	Status      InvoiceStatus // ステータス (未処理、処理中、支払い済み、エラー)
}
