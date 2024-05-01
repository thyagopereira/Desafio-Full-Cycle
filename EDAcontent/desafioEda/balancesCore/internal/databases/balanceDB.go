package databases

import (
	"database/sql"
	"fmt"

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
	queryStr := "INSERT INTO balances (id, account_id_from, account_id_to, balance_account_id_from, balance_account_id_to, created_at) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := b.DB.Prepare(queryStr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(balance.Id, balance.AccountIdFrom, balance.AccountIdTo, balance.BalanceAccountIdFrom, balance.BalanceAccountIdTo, balance.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
