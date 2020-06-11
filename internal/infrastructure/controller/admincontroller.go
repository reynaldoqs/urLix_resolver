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

func (as *adminController) Execute(res http.ResponseWriter, req *http.Request) {

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
