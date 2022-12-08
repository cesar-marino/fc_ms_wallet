package entity_test

import (
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "teste@teste.com")
	account := entity.NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := entity.NewAccount(nil)

	assert.Nil(t, account)
}

func TestCreditAccunt(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "teste@teste.com")
	account := entity.NewAccount(client)
	account.Credit(100)

	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccunt(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "teste@teste.com")
	account := entity.NewAccount(client)
	account.Credit(100)
	account.Debit(50)

	assert.Equal(t, float64(50), account.Balance)
}
