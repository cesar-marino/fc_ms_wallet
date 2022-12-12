package createclient_test

import (
	"testing"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	createclient "github.com/cesar-marino/fc_ms_wallet/internal/usecase/create_client"
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

func TestCreateClientUseCaseExceute(t *testing.T) {
	gateway := new(ClientGatewayMock)
	gateway.On("Save", mock.Anything).Return(nil)

	uc := createclient.NewCreateClientUseCase(gateway)
	input := createclient.CreateClientInputDto{
		Name:  "César Marino",
		Email: "cesar@test.com",
	}

	output, err := uc.Execut(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.ID)
	assert.Equal(t, "César Marino", output.Name)
	assert.Equal(t, "cesar@test.com", output.Email)
	gateway.AssertExpectations(t)
	gateway.AssertNumberOfCalls(t, "Save", 1)
}
