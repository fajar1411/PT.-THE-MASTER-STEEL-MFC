package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Email    string
	Nama     string
}
type Employee struct {
	gorm.Model
	Password string
	Email    string
	Nama     string
	Nip      string
	Role     string
	Division string
}
