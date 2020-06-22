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
	farmerMsging msgport.CloudMessenger
	farmerRepo   repport.FarmersRepository
	rechargeRepo repport.RechargesRespository
	ussdRepo     repport.UssdRepository
}

func NewService(
	fm msgport.CloudMessenger,
	fr repport.FarmersRepository,
	rr repport.RechargesRespository,
	ur repport.UssdRepository,
) *service {
	return &service{
		farmerMsging: fm,
		farmerRepo:   fr,
		rechargeRepo: rr,
		ussdRepo:     ur,
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
		err = errors.Wrap(err, "recharge.Create")
		return err
	}

	farmers, err := s.farmerRepo.GetAll()
	if err != nil {
		err = errors.Wrap(err, "recharge.Create")
		return err
	}

	farmer, err := getAviableFarmer(farmers, "ENTEL")
	if err != nil {
		err = errors.Wrap(err, "recharge.Create")
		return err
	}

	//set actios
	var actions []string

	action, err := s.ussdRepo.GetByAction("recarga_tarjeta")
	if err != nil {
		fmt.Println(err)
		err = errors.Wrap(err, "recharge.Create")
		return err
	}
	fmt.Println(action)
	err = action.Replace("#", recharge.CardNumber)
	if err != nil {
		fmt.Println(err)
		err = errors.Wrap(err, "recharge.Create")
		return err
	}
	fmt.Println(action)
	actions = append(actions, action.GetUSSD())

	fmt.Println(actions)
	dataToFarmer := domain.RechargeMessage{
		ExecCodes:     actions,
		TargetCompany: recharge.Company,
		IDRecharge:    recharge.ID,
		Mount:         recharge.Mount,
	}

	dataToFarmer.FarmerNumber = farmer.PhoneNumber

	err = s.farmerMsging.RechargeNotify(farmer, &dataToFarmer)
	if err != nil {
		err = errors.Wrap(err, "recharge.Create")
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
