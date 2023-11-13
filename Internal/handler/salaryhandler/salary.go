package salaryhandler

import (
	"master/domain/contract/handlecontract"
	"master/domain/contract/servicecontract"
	"master/domain/query"
	"master/domain/request"
	"master/helper"
	"master/middlewares"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Handlersalary struct {
	ss servicecontract.ServiceSalary
}

func NewHandleSalary(ss servicecontract.ServiceSalary) handlecontract.HandleSalary {
	return &Handlersalary{
		ss: ss,
	}
}

func (hs *Handlersalary) AddSalary(e echo.Context) error {

	role := middlewares.ExtractTokenRole(e)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(nil, http.StatusUnauthorized, true, "anda bukan admin maupun pegawai kami"))
	}
	requestaddsalary := request.RequestSalary{}

	nip := e.QueryParam("nip")

	binderr := e.Bind(&requestaddsalary)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(nil, http.StatusBadRequest, true, binderr.Error()))
	}
	data, errservice := hs.ss.AddSalary(nip, requestaddsalary)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errservice.Error()))
	}
	respondata := query.ReqsalarytoRespon(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false, "success addsalary"))
}
