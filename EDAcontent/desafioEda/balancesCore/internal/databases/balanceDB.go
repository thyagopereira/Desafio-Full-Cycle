package databases

import (
	"database/sql"

	"github.com/thyagopereira/full-cycle/eda/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) Save(balance entity.Balance) error {
	queryStr := "INSERT INTO balances (id, account_id, amount, created_at) VALUES (?, ?, ?, ?)"
	stmt, err := b.DB.Prepare(queryStr)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(balance.Id, balance.AccountId, balance.Amount, balance.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
