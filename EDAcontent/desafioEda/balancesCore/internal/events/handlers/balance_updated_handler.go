package handler

import (
	"fmt"
	"sync"

	"github.com/thyagopereira/full-cycle/eda/internal/databases"
	"github.com/thyagopereira/full-cycle/eda/internal/entity"
	"github.com/thyagopereira/full-cycle/eda/pkg/events"
)

type UpdateBalanceHandler struct {
	balanceRepository  databases.BalanceDB
	accountsRepository databases.AccountsDB
}

func NewUpdateBalanceKafkaHandler(balanceDB databases.BalanceDB, accountsDB databases.AccountsDB) *UpdateBalanceHandler {
	return &UpdateBalanceHandler{
		balanceRepository:  balanceDB,
		accountsRepository: accountsDB,
	}
}

func (h *UpdateBalanceHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Before the payload cast")

	payload := message.GetPayload()
	atributes := payload.(map[string]interface{})

	accountIdFrom := atributes["account_id_from"].(string)
	accountIdTo := atributes["account_id_to"].(string)
	balanceIdFrom := int(atributes["balance_account_id_to"].(float64))
	balanceIdTo := int(atributes["balance_account_id_from"].(float64))

	balance := entity.NewBalance(accountIdFrom,
		accountIdTo, balanceIdTo, balanceIdFrom)

	err := h.balanceRepository.Save(*balance)
	if err != nil {
		fmt.Println("Err on save balance")
		panic(err)
	}
	err = h.accountsRepository.UpdateBalance(accountIdFrom, balanceIdFrom)
	if err != nil {
		panic(err)
	}

	err = h.accountsRepository.UpdateBalance(accountIdTo, balanceIdTo)
	if err != nil {
		panic("Wrong on update balance")
	}

}
