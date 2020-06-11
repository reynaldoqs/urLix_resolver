package service

import (
	"github.com/pkg/errors"
	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	repport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/repositories"
)

type reportss struct {
	db repport.ReportRespository
}

func NewReportsRepository(rep repport.ReportRespository) *reportss {
	return &reportss{
		db: rep,
	}
}

func (r *reportss) RechargeReport(report *domain.RechargeReport) error {

	err := r.db.SaveR(report)
	if err != nil {
		err = errors.Wrap(err, "reports.RechargeReport")
		return err
	}
	return err
}
func (r *reportss) AdminMsgReport(report *domain.AdminMsgReport) error {

	err := r.db.SaveA(report)
	if err != nil {
		err = errors.Wrap(err, "reports.AdminMsgReport")
		return err
	}
	return err
}
