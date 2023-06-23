package main

import (
    "fmt"
	"belajar-golang/pkg/data"
)

func main(){
	firstName, lastName, age, salary := user.DataUserInference()
	fmt.Printf("firstName: %v, lastName: %v, age: %v, salary: %v\n", firstName, lastName, age, salary)
}