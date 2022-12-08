package entity_test

import (
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := entity.NewClient("César Marino", "cesar@test.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "César Marino", client.Name)
	assert.Equal(t, "cesar@test.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := entity.NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "cesar@test.com")
	err := client.Update("Outro Client", "test@test.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Outro Client", client.Name)
	assert.Equal(t, "test@test.com", client.Email)
}

func TestUpdateClientWhitInvalidArgs(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "cesar@test.com")
	err := client.Update("", "test@test.com")

	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "cesar@test.com")
	account := entity.NewAccount(client)

	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
