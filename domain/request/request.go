package request

type RequestUser struct {
	Id       int
	Password string `json:"password" form:"password" validate:"required,min=5"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Name     string `json:"nama" form:"nama" validate:"required,min=5"`
}
type RequestEmployee struct {
	Id       int
	Password string `json:"password" form:"password" validate:"required,min=5"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Name     string `json:"nama" form:"nama" validate:"required,min=5"`
	Role     string `json:"role" form:"role" validate:"required,min=4"`
	Nip      string `json:"nip" form:"nip" validate:"required,min=5"`
	Division string `json:"division" form:"division" validate:"required,min=3"`
}
type RequestSalary struct {
	Id         int
	Gaji       float64 `json:"gaji" form:"gaji" validate:"required"`
	IDEmployee uint
	Total_Gaji int
}
