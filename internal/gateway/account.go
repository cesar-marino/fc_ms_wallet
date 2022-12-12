package gateway

import "github.com/cesar-marino/fc_ms_wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
