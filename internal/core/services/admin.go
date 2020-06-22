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
	ussdRepo    repport.UssdRepository
}

func NewAdminService(
	fm msgport.CloudMessenger,
	fr repport.FarmersRepository,
	ur repport.UssdRepository) *admin {
	return &admin{
		cloudMsging: fm,
		farmerRepo:  fr,
		ussdRepo:    ur,
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

//USSD implementations

func (ad *admin) GetActions() ([]*domain.UssdAction, error) {
	return ad.ussdRepo.GetUssdActions()
}
func (ad *admin) AddAction(ussd *domain.UssdAction) error {
	return ad.ussdRepo.SaveUssd(ussd)
}
func (ad *admin) UpdateAction(ussd *domain.UssdAction) error {
	return ad.ussdRepo.UpdateUssd(ussd)
}
