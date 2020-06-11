package repositoryport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type ReportRespository interface {
	Save(report *domain.RechargeReport) error
}
