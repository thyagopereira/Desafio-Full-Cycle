package usecases

import (
	"time"

	"github.com/thyagopereira/full-cycle/eda/internal/databases"
)

type BalanceUseCase interface {
	GetBalanceFromAccoundId(id string) (*BalanceFromAccountIdDTO, error)
}

type BalanceFromAccountIdDTO struct {
	Balance   int
	AccountId string
	Timestamp time.Time
}

type BalanceUseCaseImpl struct {
	DB databases.AccountsDB
}

func NewBalanceUseCases(db databases.AccountsDB) *BalanceUseCaseImpl {
	return &BalanceUseCaseImpl{
		DB: db,
	}
}

func (b *BalanceUseCaseImpl) GetBalanceFromAccoundId(accountId string) (*BalanceFromAccountIdDTO, error) {
	result, err := b.DB.FindById(accountId)
	if err != nil {
		return nil, err
	}

	return &BalanceFromAccountIdDTO{
		Balance:   result.Balance,
		AccountId: result.Id,
		Timestamp: time.Now(),
	}, nil
}
