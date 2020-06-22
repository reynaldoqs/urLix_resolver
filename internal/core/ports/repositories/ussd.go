package repositoryport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type UssdRepository interface {
	GetUssdActions() ([]*domain.UssdAction, error)
	GetByAction(action string) (*domain.UssdAction, error)
	SaveUssd(ussd *domain.UssdAction) error
	UpdateUssd(ussd *domain.UssdAction) error
}
