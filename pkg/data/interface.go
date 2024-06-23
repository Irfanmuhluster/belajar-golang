package data

// Interface
type Speaker interface {
	Speak() string
}

// Struct yang mengimplementasikan interface Speaker
type Person struct {
	Name string
}

// method Contract
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func InterfaceX() string {
	// instance
	// var s Speaker

	p := Person{Name: "John"}
	// s = p
	return p.Speak()
}
