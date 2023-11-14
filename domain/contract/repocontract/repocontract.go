package repocontract

import "master/domain/request"

type RepoUser interface {
	RegisterUser(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	EmaiuserExist(email string) (data request.RequestUser, err error)
	Profile(id int) (data request.RequestUser, err error)
	// Filter(nama, norek string) (data []request.ReqProfile, err error)
}
type RepoLogin interface {
	LoginUser(email string, password string) (string, request.RequestUser, error)
	LoginEmp(email string, password string) (string, request.RequestEmployee, error)
}
type RepoEmployee interface {
	AddEmployee(newRequest request.RequestEmployee) (data request.RequestEmployee, err error)
	EmailExist(email string) (data request.RequestEmployee, err error)
	GetSalary(id int) (data request.RequestSalary, err error)
	NipExist(nip string) (data request.RequestEmployee, err error)
}
type RepoSalary interface {
	AddSalary(newRequest request.RequestSalary) (data request.RequestSalary, err error)
}
