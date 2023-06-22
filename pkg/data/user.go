package user

import (
    "strconv"
)

func DataUser() string {
	name := "John"
	age := 25

	gabungan := "nama = " + name + ", umur = " + strconv.Itoa(age)
	return gabungan
}