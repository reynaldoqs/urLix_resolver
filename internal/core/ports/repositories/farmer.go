package repositoryport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type FarmersRepository interface {
	GetAll() ([]*domain.Farmer, error)
	GetByNumber(num int) (*domain.Farmer, error)
}
