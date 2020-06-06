package database

import "github.com/gustavolimam/pismo-teste/types"

// InsertTransaction query to insert transaction data on database
func InsertTransaction(transaction types.Transaction) (int64, error) {
	res, err := db.Exec("INSERT INTO `transaction` (`account_id`,`operation_type`,`amount`) VALUES (?,?,?);", transaction.AccountID, transaction.OperationTypeID, transaction.Amount)

	id, _ := res.LastInsertId()

	return id, err
}
