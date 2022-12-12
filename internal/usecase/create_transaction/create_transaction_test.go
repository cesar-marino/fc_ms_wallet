package createtransaction_test

import (
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	createtransaction "github.com/cesar-marino/fc_ms_wallet/internal/usecase/create_transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionExecute(t *testing.T) {
	client1, _ := entity.NewClient("Client 1", "client1@test.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Client 2", "client2@test.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	accountGateway := new(AccountGatewayMock)
	accountGateway.On("FindById", account1.ID).Return(account1, nil)
	accountGateway.On("FindById", account2.ID).Return(account2, nil)

	transactionGateway := new(TransactionGatewayMock)
	transactionGateway.On("Create", mock.Anything).Return(nil)

	input := createtransaction.CreateTransactionInputDto{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100,
	}

	uc := createtransaction.NewCreateTransactionUseCase(transactionGateway, accountGateway)
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	accountGateway.AssertExpectations(t)
	transactionGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "FindById", 2)
	transactionGateway.AssertNumberOfCalls(t, "Create", 1)
}
