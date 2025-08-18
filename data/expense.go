package data

// Package data provides structures for handling financial data.

type Expense struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
}
