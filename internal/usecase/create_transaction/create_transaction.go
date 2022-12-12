package createtransaction

import (
	"github.com/cesar-marino/fc_ms_wallet/internal/entity"
	"github.com/cesar-marino/fc_ms_wallet/internal/gateway"
)

type CreateTransactionInputDto struct {
	AccountFromID string
	AccountToID   string
	Amount        float64
}

type CreateTransactionOutputDto struct {
	ID string
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: transactionGateway,
		accountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	acountFrom, err := uc.accountGateway.FindById(input.AccountFromID)
	if err != nil {
		return nil, err
	}

	acountTo, err := uc.accountGateway.FindById(input.AccountToID)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(acountFrom, acountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.transactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDto{
		ID: transaction.ID,
	}, nil
}
