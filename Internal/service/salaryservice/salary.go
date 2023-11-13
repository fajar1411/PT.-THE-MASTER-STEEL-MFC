package Salaryservice

import (
	"master/domain/contract/repocontract"
	"master/domain/contract/servicecontract"
	"master/domain/request"
	"time"

	"github.com/go-playground/validator"
)

type Servicessalary struct {
	rs       repocontract.RepoSalary
	re       repocontract.RepoEmployee
	validate *validator.Validate
}

func NewServiceSalary(rs repocontract.RepoSalary, re repocontract.RepoEmployee) servicecontract.ServiceSalary {
	return &Servicessalary{
		re:       re,
		rs:       rs,
		validate: validator.New(),
	}
}

func (ss *Servicessalary) AddSalary(nip string, newRequest request.RequestSalary) (data request.RequestSalary, err error) {
	nipexist, errexist := ss.re.NipExist(nip)

	if errexist != nil {
		return data, errexist
	}
	newRequest.IDEmployee = uint(nipexist.Id)
	haripertama := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	hariterakhir := haripertama.AddDate(0, 1, -1)
	var totaljam int

	for d := haripertama; !d.After(hariterakhir); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {

			totaljam += 8
		}
	}
	total := totaljam * int(newRequest.Gaji)
	newRequest.Total_Gaji = total

	datarepo, errrepo := ss.rs.AddSalary(newRequest)

	if errrepo != nil {
		return data, errrepo
	}
	return datarepo, nil
}
