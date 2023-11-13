package employeehandler

import (
	"context"
	"encoding/json"
	"master/domain/contract/handlecontract"
	"master/domain/contract/servicecontract"
	"master/domain/query"
	"master/domain/request"
	"master/helper"
	"master/redist"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Handleremployee struct {
	se servicecontract.ServiceEmployee
}

func NewHandleemployee(se servicecontract.ServiceEmployee) handlecontract.HandleEmployee {
	return &Handleremployee{
		se: se,
	}
}

func (he *Handleremployee) AddEmployee(e echo.Context) error {
	requestaddemployee := request.RequestEmployee{}

	binderr := e.Bind(&requestaddemployee)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(nil, http.StatusBadRequest, true, binderr.Error()))
	}

	client := redist.RedisClient()

	defer redist.CloseRedisConnection(client)

	_, errexist := client.HExists(context.Background(), "employee", requestaddemployee.Email).Result()
	if errexist != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errexist.Error()))
	}

	data, errservice := he.se.AddEmployee(requestaddemployee)
	userDataJSON, err := json.Marshal(data)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(err.Error(), http.StatusInternalServerError, true, err.Error()))
	}
	errredist := client.HSet(context.Background(), "users", data.Email, userDataJSON).Err()
	if errredist != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errredist.Error()))
	}
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errservice.Error()))
	}
	respondata := query.ReqemployetoRespon(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false, "success register"))
}
