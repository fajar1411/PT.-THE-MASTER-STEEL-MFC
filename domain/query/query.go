package query

import (
	"master/domain/model"
	"master/domain/request"
	"master/domain/respon"
)

func RequuserToModel(data request.RequestUser) model.User {
	return model.User{

		Email:    data.Email,
		Nama:     data.Name,
		Password: data.Password,
	}
}
func ModelusertoReq(data *model.User) request.RequestUser {
	return request.RequestUser{

		Email:    data.Email,
		Name:     data.Nama,
		Password: data.Password,
		Id:       int(data.ID),
	}
}
func ModelempToreq(data *model.Employee) request.RequestEmployee {
	return request.RequestEmployee{
		Id:       int(data.ID),
		Password: data.Password,
		Email:    data.Email,
		Name:     data.Nama,
		Role:     data.Role,
		Nip:      data.Nip,
		Division: data.Division,
	}
}

func ReqtoRespon(data request.RequestUser) respon.ResponseUser {
	return respon.ResponseUser{
		Id:       data.Id,
		Email:    data.Email,
		Name:     data.Name,
		Password: data.Password,
	}
}
func ReqtoResponLogin(data request.RequestUser, token string) respon.ResponseLogin {
	return respon.ResponseLogin{

		Email:    data.Email,
		Name:     data.Name,
		Token:    token,
		Password: data.Password,
	}
}

func ReqtoResponlogin(data request.RequestEmployee, token string) respon.ResponseLogin {
	return respon.ResponseLogin{

		Email:    data.Email,
		Name:     data.Name,
		Token:    token,
		Password: data.Password,
	}
}
func ListModeluserToReq(data []model.User) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModelusertoReq(&val))
	}
	return datareq
}
func ListrequserToRes(data []request.RequestUser) (datareq []respon.ResponseUser) {
	for _, val := range data {
		datareq = append(datareq, ReqtoRespon(val))
	}
	return datareq
}
func ReqemployetoRespon(data request.RequestEmployee) respon.ResponEmployee {
	return respon.ResponEmployee{
		Id:       data.Id,
		Password: data.Password,
		Email:    data.Email,
		Name:     data.Name,
		Role:     data.Role,
		Nip:      data.Nip,
		Division: data.Division,
	}
}

func RequesempToModel(data request.RequestEmployee) model.Employee {
	return model.Employee{

		Password: data.Password,
		Email:    data.Email,
		Nama:     data.Name,
		Nip:      data.Nip,
		Role:     data.Role,
		Division: data.Division,
	}
}
func ReqsalarytoRespon(data request.RequestSalary) respon.ResponSalary {
	return respon.ResponSalary{
		Id:         data.Id,
		Gaji:       data.Gaji,
		IDEmployee: data.IDEmployee,
		Total_Gaji: data.Total_Gaji,
	}
}
func MdelsalarytoReq(data model.Salary) request.RequestSalary {
	return request.RequestSalary{
		Gaji:       data.Gaji,
		IDEmployee: data.IDEmployee,
		Total_Gaji: data.Total_Gaji,
	}
}
func Reqtomodel(data request.RequestSalary) model.Salary {
	return model.Salary{
		Gaji:       data.Gaji,
		IDEmployee: data.IDEmployee,
		Total_Gaji: data.Total_Gaji,
	}
}
