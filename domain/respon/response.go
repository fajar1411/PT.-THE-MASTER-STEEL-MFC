package respon

type ResponseLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
type ResponseUser struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
type ResponEmployee struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"nama"`
	Role     string `json:"role"`
	Nip      string `json:"nip"`
	Division string `json:"division"`
}
type ResponSalary struct {
	Gaji       float64 `json:"gaji"`
	IDEmployee uint    `json:"id_emp"`
	Total_Gaji int     `json:"total_gaji"`
}
