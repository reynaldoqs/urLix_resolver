package serviceport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type RechargesService interface {
	Validate(recharge *domain.Recharge) error
	Create(recharge *domain.Recharge) error
	List() ([]*domain.Recharge, error)
}
