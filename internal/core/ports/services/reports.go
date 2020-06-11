package serviceport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type ReportsService interface {
	RechargeReport(report *domain.RechargeReport) error
}
