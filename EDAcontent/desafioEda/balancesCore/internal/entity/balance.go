package entity

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	Id                   string
	AccountIdFrom        string
	AccountIdTo          string
	BalanceAccountIdFrom int
	BalanceAccountIdTo   int
	CreatedAt            time.Time
}

func NewBalance(accountIdFrom, accountIdTo string, balanceAccountIdTo, balanceAccountIdFrom int) *Balance {
	b := &Balance{
		Id:                   uuid.New().String(),
		AccountIdFrom:        accountIdFrom,
		AccountIdTo:          accountIdTo,
		BalanceAccountIdFrom: balanceAccountIdFrom,
		BalanceAccountIdTo:   balanceAccountIdTo,
		CreatedAt:            time.Now(),
	}

	return b
}
