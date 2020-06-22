package serviceport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type AdminService interface {
	Execute(order *domain.AdminMessage) error
	GetActions() ([]*domain.UssdAction, error)
	AddAction(ussd *domain.UssdAction) error
	UpdateAction(ussd *domain.UssdAction) error
}
