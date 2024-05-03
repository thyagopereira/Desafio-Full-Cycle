package databases

import (
	"database/sql"
	"fmt"

	"github.com/thyagopereira/full-cycle/eda/internal/entity"
)

type AccountsDB struct {
	DB *sql.DB
}

func NewAccountsDB(db *sql.DB) *AccountsDB {
	return &AccountsDB{
		DB: db,
	}
}

func (a *AccountsDB) FindById(accountId string) (*entity.Account, error) {
	var account entity.Account

	queryString := `SELECT a.id, a.balance FROM accounts a WHERE a.id = ?`
	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountId)
	row.Scan(
		&account.Id,
		&account.Balance,
	)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountsDB) UpdateBalance(accountId string, balance int) error {
	queryString := "UPDATE accounts SET balance = ? WHERE id = ?"
	stmt, err := a.DB.Prepare(queryString)

	fmt.Println("querystring as prepared", err)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance, accountId)
	if err != nil {
		return err
	}

	return nil
}
