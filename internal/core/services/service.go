package service

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	msgport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/cloudmessaging"
	repport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/repositories"
)

type service struct {
	farmerMsging msgport.FarmerMessenger
	farmerRepo   repport.FarmersRepository
	rechargeRepo repport.RechargesRespository
}

func NewService(
	fm msgport.FarmerMessenger,
	fr repport.FarmersRepository,
	rr repport.RechargesRespository) *service {
	return &service{
		farmerMsging: fm,
		farmerRepo:   fr,
		rechargeRepo: rr,
	}
}

func (s *service) Validate(recharge *domain.Recharge) error {
	v := validator.New()
	err := v.Struct(recharge)
	if err != nil {
		var errors string
		for _, err := range err.(validator.ValidationErrors) {
			errors += fmt.Sprintf("error: %v isn't aceptable value for %v \n", err.Value(), err.Field())
		}

		return fmt.Errorf(errors)
	}
	return err
}

func (s *service) Create(recharge *domain.Recharge) error {
	recharge.CreatedAt = time.Now()

	err := s.rechargeRepo.Save(recharge)
	if err != nil {
		err = errors.Wrap(err, "service.Create")
		return err
	}

	farmers, err := s.farmerRepo.GetAll()
	if err != nil {
		err = errors.Wrap(err, "service.Create")
		return err
	}

	dataToFarmer := domain.FarmerCloudMessage{
		ExecCodes:  []string{"*#62#", "*10*6*10#"},
		Company:    recharge.Company,
		IDRecharge: recharge.ID,
		Mount:      recharge.Mount,
	}

	farmer, err := getAviableFarmer(farmers, "ENTEL")
	if err != nil {
		err = errors.Wrap(err, "service.Create")
		return err
	}

	dataToFarmer.FarmerNumber = farmer.PhoneNumber

	err = s.farmerMsging.Notify(farmer, &dataToFarmer)
	if err != nil {
		err = errors.Wrap(err, "service.Create")
		return err
	}
	return err
}

func (s *service) List() ([]*domain.Recharge, error) {
	return s.rechargeRepo.GetAll()
}

//utils
//improve criteria
func getAviableFarmer(farmers []*domain.Farmer, criteria string) (*domain.Farmer, error) {
	for _, farmer := range farmers {
		if farmer.Company == criteria {
			fmt.Println("we found the farmer", farmer)
			return farmer, nil
		}
	}
	return nil, fmt.Errorf("Not found farmer with that criteri")
}
