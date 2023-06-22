package user

import (
    "strconv"
)

func dataUser(name string, age int)(string) {
	// name := "John"
	// age := 25

	gabungan := "nama = " + name + ", umur = " + strconv.Itoa(age)
	return gabungan
}