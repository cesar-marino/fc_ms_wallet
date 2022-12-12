package createaccount

import (
	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	"github.com/cesar-marino/fc_ms_wallet/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientID string
}

type CreateAccountOutputDto struct {
	ID string
}

type CreateAccountUseCase struct {
	accountGateway gateway.AccountGateway
	clientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway: accountGateway,
		clientGateway:  clientGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := uc.clientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	if err = uc.accountGateway.Save(account); err != nil {
		return nil, err
	}

	return &CreateAccountOutputDto{
		ID: account.ID,
	}, nil
}
