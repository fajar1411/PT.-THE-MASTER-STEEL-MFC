package loginhandler

import (
	"context"
	"master/domain/contract/handlecontract"
	"master/domain/contract/servicecontract"
	"master/domain/query"
	"master/domain/request"
	"master/helper"
	"master/redist"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type HandlerLogin struct {
	sl servicecontract.ServiceLogin
	handlecontract.HandleLogin
}

func NewHandlLogin(sl servicecontract.ServiceLogin) *HandlerLogin {
	return &HandlerLogin{
		sl: sl,
	}
}

func (hl *HandlerLogin) LoginUser(e echo.Context) error {
	reques := request.RequestUser{}

	errbind := e.Bind(&reques)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(nil, http.StatusBadRequest, true, errbind.Error()))
	}
	client := redist.RedisClient()

	defer redist.CloseRedisConnection(client)
	token, dataservice, errservice := hl.sl.LoginUser(reques.Email, reques.Password)
	err := client.Set(context.Background(), "login_user", token, redist.Expiration).Err()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, err.Error()))
	}
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errservice.Error()))
	}
	e.Response().Header().Set("Authorization", token)
	respon := query.ReqtoResponLogin(dataservice, token)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false, "Succes Login"))
}

func (hl *HandlerLogin) LogoutUser(e echo.Context) error {
	token := e.Request().Header.Get("Authorization")

	if token == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(nil, http.StatusUnauthorized, true, "invalid token"))
	}

	client := redist.RedisClient()
	defer redist.CloseRedisConnection(client)

	err := client.Del(context.Background(), "login_user", token).Err()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, err.Error()))
	}

	return e.JSON(http.StatusOK, helper.GetResponse(nil, http.StatusOK, false, "logout success"))
}

// LoginEmp implements handlecontract.HandleLogin.
func (hl *HandlerLogin) LoginEmp(e echo.Context) error {
	request := request.RequestEmployee{}

	errbind := e.Bind(&request)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(nil, http.StatusBadRequest, true, errbind.Error()))
	}
	client := redist.RedisClient()

	defer redist.CloseRedisConnection(client)
	// fmt.Print("ini handle emp", request.Email)
	token, dataservice, errservice := hl.sl.LoginEmp(request.Email, request.Password)

	err := client.Set(context.Background(), "login_emp", token, redist.Expiration).Err()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, err.Error()))
	}
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(nil, http.StatusInternalServerError, true, errservice.Error()))
	}
	e.Response().Header().Set("Authorization", token)
	respon := query.ReqtoResponlogin(dataservice, token)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false, "Succes Login"))
}
