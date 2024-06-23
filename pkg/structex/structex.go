package data

import "strconv"

// ini krn string tidak boleh dicampur oleh INT
// template data utk menggabungkan beberapatipe data (kumpulan data field)
type Customer struct {
	Name, Address string
	Age           int
}

func DataCustomer(eko Customer) string {
	eko.Name = "Eko Kurniwan"
	eko.Address = "Jl Gejayan No 82"
	eko.Age = 23

	return eko.Name + "Alamat = " + eko.Address + "Age = " + strconv.Itoa(eko.Age)

	// p := Customer{"Joko", "Jl Patang Puluhan", 30}
	// fmt.Println(p)
}
