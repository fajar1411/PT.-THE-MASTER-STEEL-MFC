package servicecontract

import "master/domain/request"

type ServiceCase interface {
	RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	Profile(id int) (data request.RequestUser, err error)
	// Filter(nama, norek string) (data []request.ReqProfile, err error)
}

type ServiceLogin interface {
	LoginUser(email string, password string) (string, request.RequestUser, error)
	LoginEmp(email string, password string) (string, request.RequestEmployee, error)
}
type ServiceEmployee interface {
	AddEmployee(newRequest request.RequestEmployee) (data request.RequestEmployee, err error)
}
