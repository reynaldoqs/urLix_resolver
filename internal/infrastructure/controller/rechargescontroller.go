package controller

import (
	"encoding/json"
	"net/http"

	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	serviceport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/services"
)

type rechargesController struct {
	rservice serviceport.RechargesService
}

func NewRechargesController(srv serviceport.RechargesService) *rechargesController {
	return &rechargesController{
		rservice: srv,
	}
}

func (rc *rechargesController) AddRecharge(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var recharge domain.Recharge
	err := json.NewDecoder(req.Body).Decode(&recharge)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		//	json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err = rc.rservice.Validate(&recharge)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err)
		return
	}
	err = rc.rservice.Create(&recharge)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func (rc *rechargesController) GetRecharges(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	recharges, err := rc.rservice.List()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(recharges)
}
