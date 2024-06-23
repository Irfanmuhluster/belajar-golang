package data

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

func IfCondision() string {
	// if shortStatement; condition {
	// 	// Blok kode yang dieksekusi jika kondisi benar
	// }
	if age := 25; age < 18 {
		return "Anda masih di bawah umur."
	} else {
		return "Anda sudah dewasa."
	}
}
