package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thyagopereira/full-cycle/eda/internal/entity"
)

func TestCreateBalance(t *testing.T) {
	b := entity.NewBalance("123", 10)
	assert.NotNil(t, b)
	assert.Equal(t, b.AccountId, "123")
	assert.Equal(t, b.Amount, 10)
}

func TestCreateBalanceNegativeAmount(t *testing.T) {

	assert.Panics(t, func() {
		entity.NewBalance("1234", -1)
	})
}
