package repositoryport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type RechargesRespository interface {
	GetAll() ([]*domain.Recharge, error)
	Save(recharge *domain.Recharge) error
	//Update()
}
