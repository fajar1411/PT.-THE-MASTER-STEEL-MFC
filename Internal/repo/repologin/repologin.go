package repologin

import (
	"errors"
	"fmt"
	"master/domain/contract/repocontract"
	"master/domain/model"
	"master/domain/query"
	"master/domain/request"
	"master/middlewares"

	"gorm.io/gorm"
)

type Repologin struct {
	db *gorm.DB
}

func NewRepoLogin(db *gorm.DB) repocontract.RepoLogin {
	return &Repologin{
		db: db,
	}
}

func (rl *Repologin) LoginUser(email string, password string) (string, request.RequestUser, error) {
	userdata := model.User{}

	tx := rl.db.Where("email = ?", email).First(&userdata)
	if tx.Error != nil {
		return "", request.RequestUser{}, tx.Error
	}
	createtoken, errtoken := middlewares.CreateToken(int(userdata.ID), userdata.Email, "")

	if errtoken != nil {
		return "", request.RequestUser{}, errors.New("gagal membuat token")
	}

	datamodeltoreq := query.ModelusertoReq(&userdata)
	return createtoken, datamodeltoreq, nil
}

// LoginEmp implements repocontract.RepoLogin.
func (rl *Repologin) LoginEmp(email string, password string) (string, request.RequestEmployee, error) {
	empdata := model.Employee{}
	fmt.Print("ini", email)

	tx := rl.db.Where("email = ?", email).First(&empdata)
	if tx.Error != nil {
		return "", request.RequestEmployee{}, tx.Error
	}
	createtoken, errtoken := middlewares.CreateToken(int(empdata.ID), empdata.Email, empdata.Role)

	if errtoken != nil {
		return "", request.RequestEmployee{}, errors.New("gagal membuat token")
	}

	datamodeltoreq := query.ModelempToreq(&empdata)
	return createtoken, datamodeltoreq, nil
}
