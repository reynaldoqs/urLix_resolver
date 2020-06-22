package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	serviceport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/services"
)

type adminController struct {
	aservice serviceport.AdminService
}

func NewAdminController(srv serviceport.AdminService) *adminController {
	return &adminController{
		aservice: srv,
	}
}

func (as *adminController) PostExecution(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var message domain.AdminMessage
	err := json.NewDecoder(req.Body).Decode(&message)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		//	json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err = as.aservice.Execute(&message)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func (as *adminController) GetUssdActions(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	ussds, err := as.aservice.GetActions()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(ussds)
}

func (as *adminController) PostUssdAction(res http.ResponseWriter, req *http.Request) {

	var ussd domain.UssdAction
	err := json.NewDecoder(req.Body).Decode(&ussd)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	err = as.aservice.AddAction(&ussd)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func (as *adminController) PatchUssdAction(res http.ResponseWriter, req *http.Request) {

	var ussd domain.UssdAction
	err := json.NewDecoder(req.Body).Decode(&ussd)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	err = as.aservice.UpdateAction(&ussd)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.WriteHeader(http.StatusNoContent)
}
