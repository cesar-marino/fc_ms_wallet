package createaccount_test

import (
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	createaccount "github.com/cesar-marino/fc_ms_wallet/internal/usecase/create_account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
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

func TestCreateAccountUseCaseExecute(t *testing.T) {
	client, _ := entity.NewClient("César Marino", "cesar@test.com")

	clientGateway := new(ClientGatewayMock)
	clientGateway.On("Get", client.ID).Return(client, nil)

	accountGateway := new(AccountGatewayMock)
	accountGateway.On("Save", mock.Anything).Return(nil)

	uc := createaccount.NewCreateAccountUseCase(accountGateway, clientGateway)
	input := createaccount.CreateAccountInputDto{
		ClientID: client.ID,
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientGateway.AssertExpectations(t)
	accountGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Get", 1)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}
