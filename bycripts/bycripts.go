package bycripts

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Bcript(y string) string {
	password := []byte(y)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)

}

func CheckPassword(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return errors.New(err.Error())
	}
	return nil
}
