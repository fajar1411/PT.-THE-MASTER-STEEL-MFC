package userhandler

import (
	"context"
	"encoding/json"
	"master/domain/contract/handlecontract"
	"master/domain/contract/servicecontract"
	"master/domain/query"
	"master/domain/request"
	"master/helper"
	"master/middlewares"
	"master/redist"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type HandlerUser struct {
	um servicecontract.ServiceCase
}

func NewHandleUser(um servicecontract.ServiceCase) handlecontract.HandleUser {
	return &HandlerUser{
		um: um,
	}
}

func (hu *HandlerUser) RegisterUser(e echo.Context) error {
	requestRegister := request.RequestUser{}

	binderr := e.Bind(&requestRegister)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(nil, http.StatusBadRequest, true, binderr.Error()))
	}

	client := redist.RedisClient()

	defer redist.CloseRedisConnection(client)

	_, errexist := client.HExists(context.Background(), "users", requestRegister.Email).Result()
	if errexist != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errexist.Error()))
	}

	data, errservice := hu.um.RegisterUser(requestRegister)
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
	respondata := query.ReqtoRespon(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false, "success register"))
}

// AllUser implements handlecontract.HandleUser.
func (hc *HandlerUser) AllUser(e echo.Context) error {
	dataservice, errservice := hc.um.AllUser()
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errservice.Error()))
	}
	respondata := query.ListrequserToRes(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false, "succes melihat semua data user"))

}

func (hu *HandlerUser) Profile(e echo.Context) error {
	id := middlewares.ExtractTokenId(e)

	if id <= 0 {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(nil, http.StatusUnauthorized, true, "id tidak ada"))

	}
	dataservice, errservice := hu.um.Profile(id)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, "gagal melihat data"))
	}
	respon := query.ReqtoRespon(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false, "sukses melihat profile"))
}

// // Filter implements handlecontract.HandleUser.
// func (hu *HandlerUser) Filter(e echo.Context) error {
// 	nama := e.QueryParam("nama")
// 	norek := e.QueryParam("norek")

// 	dataservice, errservice := hu.um.Filter(nama, norek)

// 	if errservice != nil {
// 		return e.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, http.StatusInternalServerError, true))
// 	}
// 	respon := query.FilterequserToRes(dataservice)
// 	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, true))

// }
