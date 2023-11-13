package employservice

import (
	"errors"
	"master/bycripts"
	"master/domain/contract/repocontract"
	"master/domain/contract/servicecontract"
	"master/domain/request"
	"master/validasi"

	"github.com/go-playground/validator"
)

type ServicesEmployee struct {
	ru       repocontract.RepoUser
	re       repocontract.RepoEmployee
	validate *validator.Validate
}

func NewServiceemployee(re repocontract.RepoEmployee, ru repocontract.RepoUser) servicecontract.ServiceEmployee {
	return &ServicesEmployee{
		re:       re,
		ru:       ru,
		validate: validator.New(),
	}
}

// AddEmployee implements servicecontract.ServiceEmployee.
func (se *ServicesEmployee) AddEmployee(newRequest request.RequestEmployee) (data request.RequestEmployee, err error) {
	validerr := se.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestEmployee{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	haspw := bycripts.Bcript(newRequest.Password)
	newRequest.Password = haspw

	datarepo, errrepo := se.re.AddEmployee(newRequest)

	if errrepo != nil {
		return request.RequestEmployee{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}
