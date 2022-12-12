package gateway

import "github.com/cesar-marino/fc_ms_wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
