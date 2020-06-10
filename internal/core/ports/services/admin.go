package serviceport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type AdminService interface {
	Execute(order *domain.AdminMessage) error
}
