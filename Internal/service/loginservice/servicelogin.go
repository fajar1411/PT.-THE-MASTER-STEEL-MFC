package loginservice

import (
	"errors"
	"master/bycripts"
	"master/domain/contract/repocontract"
	"master/domain/contract/servicecontract"
	"master/domain/request"
)

type ServiceLogin struct {
	rl repocontract.RepoLogin
	ru repocontract.RepoUser
}

func NewServiceLogin(rl repocontract.RepoLogin, ru repocontract.RepoUser) servicecontract.ServiceLogin {
	return &ServiceLogin{
		rl: rl,
		ru: ru,
	}
}

func (sl *ServiceLogin) LoginUser(email string, password string) (string, request.RequestUser, error) {
	if email == "" || password == "" {
		return "", request.RequestUser{}, errors.New("inputan tidak boleh kosong")
	}
	_, cekremail := sl.ru.EmaiuserExist(email)

	if cekremail != nil {
		return "", request.RequestUser{}, errors.New("email Not found")
	}
	token, datarepo, errrepo := sl.rl.LoginUser(email, password)
	checkpw := bycripts.CheckPassword(datarepo.Password, password)

	if checkpw != nil {
		return "", request.RequestUser{}, errors.New("password anda salah")
	}
	if errrepo != nil {
		return "", request.RequestUser{}, errors.New(errrepo.Error())
	}
	return token, datarepo, nil
}

// LoginEmp implements servicecontract.ServiceLogin.
func (sl *ServiceLogin) LoginEmp(email string, password string) (string, request.RequestEmployee, error) {
	if email == "" || password == "" {
		return "", request.RequestEmployee{}, errors.New("inputan tidak boleh kosong")
	}
	// _, cekremail := sl.ru.EmaiuserExist(email)

	// if cekremail != nil {
	// 	return "", request.RequestEmployee{}, errors.New("email Not found")
	// }
	token, datarepo, errrepo := sl.rl.LoginEmp(email, password)
	checkpw := bycripts.CheckPassword(datarepo.Password, password)

	if checkpw != nil {
		return "", request.RequestEmployee{}, errors.New("password anda salah")
	}
	if errrepo != nil {
		return "", request.RequestEmployee{}, errors.New(errrepo.Error())
	}
	return token, datarepo, nil
}
