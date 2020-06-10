package service

import (
	"github.com/pkg/errors"
	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	msgport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/cloudmessaging"
	repport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/repositories"
)

type admin struct {
	cloudMsging msgport.CloudMessenger
	farmerRepo  repport.FarmersRepository
}

func NewAdminService(
	fm msgport.CloudMessenger,
	fr repport.FarmersRepository) *admin {
	return &admin{
		cloudMsging: fm,
		farmerRepo:  fr,
	}
}

func (ad *admin) Execute(order *domain.AdminMessage) error {
	farmer, err := ad.farmerRepo.GetByNumber(order.FarmerNumber)
	if err != nil {
		err = errors.Wrap(err, "admin.Execute")
		return err
	}

	message := domain.AdminMessage{
		ExecCodes:    []string{"*#62#", "*10*6*10#"},
		FarmerNumber: farmer.PhoneNumber,
		IDMessage:    "00",
	}

	err = ad.cloudMsging.AdminNotify(farmer, &message)
	if err != nil {
		err = errors.Wrap(err, "admin.Execute")
		return err
	}
	return err
}
