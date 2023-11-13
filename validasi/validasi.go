package validasi

import (
	"fmt"

	"github.com/go-playground/validator"
)

func ValidationErrorHandle(err error) string {
	var message string

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				message = fmt.Sprintf("Ada Tabel yang Kosong, Harap Diisi !!!")
			case "min":
				message = fmt.Sprintf(" Input Minimal 5 character !!!")
			case "max":
				message = fmt.Sprintf(" Input Maksimal 15 character !!!")
			case "lte":
				message = fmt.Sprintf(" Input tidak boleh di bawah 2 !!!")
			case "gte":
				message = fmt.Sprintf(" Input tidak boleh di atas 15 !!!")
			case "numeric":
				message = fmt.Sprintf("Input Harus angka !!!")
			case "url":
				message = fmt.Sprintf(" Input Harus Url !!!")
			case "password":
				message = fmt.Sprintf(" Input value must be filled")
			case "alphanum":
				message = fmt.Sprintf(" input harus angka dan huruf")
			case "alpha":
				message = fmt.Sprintf(" input harus huruf")
			case "email":
				message = fmt.Sprintf("%s Input Harus Email !!!", v.Field())
			}
		}
	}

	return message
}
