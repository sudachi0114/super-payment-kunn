package models

import "time"

// 請求書データ (企業・取引先に紐づく)
type Invoice struct {
	ID      int       `json:"id"`
	Amount  int       `json:"amount"`   // 支払金額
	DueDate time.Time `json:"due_date"` // 支払期日
	// 発行日
	// 請求金額
	// 手数料
	// 手数料率
	// 消費税
	// 消費税率
	// ステータス (未処理、処理中、支払い済み、エラー)
}
