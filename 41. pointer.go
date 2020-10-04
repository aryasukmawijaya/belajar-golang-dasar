package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	// address2 adalah variabel dgn nilai pointer ke variabel address1
	// deklarasi tipe data menggunakan * dan valuenya menggunakan &
	var address2 *Address = &address1

	// address3 tidak mengarah ke pointer address1
	var address3 Address = address1

	// jika hanya mengubah sebagian, address1 tidak terpengaruh
	address2.City = "Sukabumi"

	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address3 = Address{"Cirebon", "Jawa Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	// sehingga address3 tidak berubah
	fmt.Println(address3)
}
