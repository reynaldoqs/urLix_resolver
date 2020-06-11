package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	serviceport "github.com/reynaldoqs/urLix_resolver/internal/core/ports/services"
)

type reportsController struct {
	rservice serviceport.ReportsService
}

func NewReportsController(srv serviceport.ReportsService) *reportsController {
	return &reportsController{
		rservice: srv,
	}
}

func (rs *reportsController) PostRechargeReport(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var report domain.RechargeReport
	err := json.NewDecoder(req.Body).Decode(&report)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	err = rs.rservice.RechargeReport(&report)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.WriteHeader(http.StatusNoContent)
}

func (rs *reportsController) PostAdminReport(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var report domain.AdminMsgReport
	err := json.NewDecoder(req.Body).Decode(&report)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	err = rs.rservice.AdminMsgReport(&report)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.WriteHeader(http.StatusNoContent)
}
