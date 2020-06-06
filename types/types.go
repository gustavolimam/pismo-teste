package types

import "time"

type Account struct {
	AccountID      int    `json:"account_id" db:"account_id"`
	DocumentNumber string `json:"document_number" db:"document_number"`
}

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float32   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}
