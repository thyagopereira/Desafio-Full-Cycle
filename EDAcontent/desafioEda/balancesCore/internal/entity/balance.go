package entity

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	Id        string
	AccountId string
	Amount    int
	CreatedAt time.Time
}

func NewBalance(accountId string, amount int) *Balance {
	b := &Balance{
		Id:        uuid.New().String(),
		AccountId: accountId,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	if !b.isValid() {
		panic("Amount < 0")
	}

	return b
}

func (b *Balance) isValid() bool {
	return b.Amount > 0
}
