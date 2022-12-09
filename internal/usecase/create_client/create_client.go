package createclient

import (
	"time"

	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	"github.com/cesar-marino/fc_ms_wallet/internal/gateway"
)

type CreateClientInputDto struct {
	Name  string
	Email string
}

type CreateClientOutputDto struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	client gateway.ClientGateway
}

func NewCreateClientUseCase(client gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		client: client,
	}
}

func (uc *CreateClientUseCase) Execut(input CreateClientInputDto) (*CreateClientOutputDto, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = uc.client.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDto{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
