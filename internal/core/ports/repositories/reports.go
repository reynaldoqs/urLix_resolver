package repositoryport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type ReportRespository interface {
	SaveR(report *domain.RechargeReport) error
	SaveA(report *domain.AdminMsgReport) error
}
