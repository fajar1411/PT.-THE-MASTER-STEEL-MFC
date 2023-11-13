package userservice

import (
	"errors"
	"master/bycripts"
	"master/domain/contract/repocontract"
	"master/domain/contract/servicecontract"
	"master/domain/request"
	"master/validasi"

	"github.com/go-playground/validator"
)

type ServicesCase struct {
	ru       repocontract.RepoUser
	validate *validator.Validate
}

func NewServiceUser(ru repocontract.RepoUser) servicecontract.ServiceCase {
	return &ServicesCase{
		ru:       ru,
		validate: validator.New(),
	}
}

func (sc *ServicesCase) RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error) {
	validerr := sc.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestUser{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	haspw := bycripts.Bcript(newRequest.Password)
	newRequest.Password = haspw

	datarepo, errrepo := sc.ru.RegisterUser(newRequest)

	if errrepo != nil {
		return request.RequestUser{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}

func (sc *ServicesCase) AllUser() (data []request.RequestUser, err error) {
	datarepo, errrepo := sc.ru.AllUser()

	if errrepo != nil {
		return []request.RequestUser{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}

func (sc *ServicesCase) Profile(id int) (data request.RequestUser, err error) {
	data, err = sc.ru.Profile(id)

	if err != nil {
		return data, err
	}
	return data, nil
}

// // Filter implements servicecontract.ServiceCase.
// func (sc *ServicesCase) Filter(nama string, norek string) (data []request.ReqProfile, err error) {
// 	datarepo, errrepo := sc.ru.Filter(nama, norek)

// 	if err != nil {
// 		return data, errrepo
// 	}
// 	return datarepo, nil
// }
