package database

import (
	"github.com/gustavolimam/pismo-teste/types"
)

// InsertAccount query to insert account data on database
func InsertAccount(docNumber string) (int64, error) {
	res, err := db.Exec("INSERT INTO `account` (`document_number`) VALUES (?); ", docNumber)

	id, _ := res.LastInsertId()

	return id, err
}

// GetAccount query to return data for a specific account
func GetAccount(accountID string) (*types.Account, error) {

	account := new(types.Account)

	err := db.QueryRow("SELECT `account_id`, `document_number` FROM `account` WHERE `account_id`=?", accountID).
		Scan(&account.AccountID, &account.DocumentNumber)

	return account, err
}
