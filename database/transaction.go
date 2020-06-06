package database

import "github.com/gustavolimam/pismo-teste/types"

// InsertTransaction query to insert transaction data on database
func InsertTransaction(transaction types.Transaction) error {
	_, err := db.Exec("INSERT INTO `transaction` (`account_id`,`operation_type`,`amount`) VALUES (?,?,?);", transaction.AccountID, transaction.OperationTypeID, transaction.Amount)

	return err
}
