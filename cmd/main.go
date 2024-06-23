package main

import (
	user "belajar-golang/pkg/data"
	"fmt"
)

func main() {
	firstName, lastName, age, salary := user.DataUserInference()
	fmt.Printf("firstName: %v, lastName: %v, age: %v, salary: %v\n", firstName, lastName, age, salary)

	ifCondition := user.IfCondision()
	fmt.Printf(ifCondition)

	// Membuat instance dari Customer (struct)
	var customer user.Customer
	dataStruct := user.DataCustomer(customer)
	// Print field satu per satu
	fmt.Println("Name:", dataStruct.Name)
	fmt.Println("Address:", dataStruct.Address)
	fmt.Println("Age:", dataStruct.Age)

	interfaceX := user.InterfaceX()
	fmt.Println(interfaceX)
}
