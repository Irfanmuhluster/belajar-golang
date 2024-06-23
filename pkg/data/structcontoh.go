package data

// ini krn string tidak boleh dicampur oleh INT

type Customer struct {
	Name, Address string
	Age           int
}

func DataCustomer(eko Customer) Customer {
	eko.Name = "Eko Kurniwan"
	eko.Address = "Jl Gejayan No 82"
	eko.Age = 23

	return eko
}
