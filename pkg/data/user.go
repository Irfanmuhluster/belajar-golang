package user

import (
	"strconv"
)

func DataUser(name string, age int) string {
	gabungan := "nama = " + name + ", umur = " + strconv.Itoa(age)
	return gabungan
}

/**
*	 type inference
 */
func DataUserInference() (string, string, int, float64) {
	firstName := "Amir"
	lastName := "Sarifudin"
	age := 26
	salary := 50000.25
	return firstName, lastName, age, salary
}
