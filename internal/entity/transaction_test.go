package entity_test

import (
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := entity.NewClient("Client 1", "client1@test.com")
	account1 := entity.NewAccount(client1)
	client2, _ := entity.NewClient("Client 2", "client2@test.com")
	account2 := entity.NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	client1, _ := entity.NewClient("Client 1", "client1@test.com")
	account1 := entity.NewAccount(client1)
	client2, _ := entity.NewClient("Client 2", "client2@test.com")
	account2 := entity.NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account1, account2, 2000)

	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}
