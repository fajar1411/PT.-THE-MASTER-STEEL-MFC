package repoemployee

import (
	"errors"
	"master/domain/contract/repocontract"
	"master/domain/model"
	"master/domain/query"
	"master/domain/request"

	"gorm.io/gorm"
)

type Repoemployee struct {
	db *gorm.DB
}

func NewRepoEmployee(db *gorm.DB) repocontract.RepoEmployee {
	return &Repoemployee{
		db: db,
	}
}

// AddEmployee implements repocontract.RepoEmployee.
func (re *Repoemployee) AddEmployee(newRequest request.RequestEmployee) (data request.RequestEmployee, err error) {
	datareqtomodelemp := query.RequesempToModel(newRequest)
	_, errnip := re.NipExist(datareqtomodelemp.Nip)

	_, errexist := re.EmailExist(datareqtomodelemp.Email)

	if errexist == nil || errnip != nil {
		return request.RequestEmployee{}, errors.New("Email Sudah Terdaftar")
	}

	if errnip == nil || errexist != nil {
		return request.RequestEmployee{}, errors.New("nip Sudah Terdaftar")
	}
	if errnip == nil || errexist == nil {
		return request.RequestEmployee{}, errors.New("anda Sudah Terdaftar")
	}
	tx := re.db.Create(&datareqtomodelemp)

	if tx.Error != nil {
		return request.RequestEmployee{}, tx.Error
	}

	datamodeltoreq := query.ModelempToreq(&datareqtomodelemp)

	return datamodeltoreq, nil
}

func (re *Repoemployee) EmailExist(email string) (data request.RequestEmployee, err error) {
	var activ model.Employee

	tx := re.db.Raw("Select employees.id, employees.password, employees.email, employees.nama,employees.nip,employees.role,employees.division from employees WHERE employees.email= ? ", email).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestEmployee{}, tx.Error
	}
	var activcore = query.ModelempToreq(&activ)
	return activcore, nil
}

// NipExist implements repocontract.RepoEmployee.
func (re *Repoemployee) NipExist(nip string) (data request.RequestEmployee, err error) {
	var activ model.Employee

	tx := re.db.Raw("Select employees.id, employees.password, employees.email, employees.nama,employees.nip,employees.role,employees.division from employees WHERE employees.nip= ? ", nip).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestEmployee{}, tx.Error
	}
	var activcore = query.ModelempToreq(&activ)
	return activcore, nil
}

// GetSalary implements repocontract.RepoEmployee.
func (re *Repoemployee) GetSalary(id int) (data request.RequestSalary, err error) {
	var activ model.Salary

	tx := re.db.Raw("Select salaries.id,  salaries.gaji,  salaries.id_employee,  salaries.total_gaji from salaries WHERE salaries.id_employee= ? ", id).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestSalary{}, tx.Error
	}
	var activcore = query.MdelsalarytoReq(activ)
	return activcore, nil
}
