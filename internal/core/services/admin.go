package service

import (
	"fmt"

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
	// save on repo then get id
	order.IDMessage = "noId"
	fmt.Println(order)
	err = ad.cloudMsging.AdminNotify(farmer, order)
	if err != nil {
		err = errors.Wrap(err, "admin.Execute")
		return err
	}
	return err
}
