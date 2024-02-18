package main

import "fmt"

func main() {
	// var nama string = "Muhamad Rizki"
	var usia int = 25
	var tinggi float32 = 169.9
	println(usia, tinggi)

	var nama1 = "Kibo"
	var usia2 = 25
	var tinggi2 = 199.9
	println(nama1, usia2, tinggi2)

	nama3 := "Bokir"
	usia3 := 25
	tinggi3 := 175.5
	println(nama3, usia3, tinggi3)

	var nama4 string
	nama4 = "Kibo BOkir"
	println(nama4)

	//multiple variables and values.
	var nama, alamat, kota string = "Kibo", "Purwkarta", "Purwakarta"
	println(nama, alamat, kota)
	var Panjang, Lebar, Tinggi int = 5, 10, 12
	println(Panjang, Lebar, Tinggi)
	var x, y, z float64 = 13.3, 13.5, 13.9
	println(x, y, z)

	Nama, Alamat, Kota := "Kibo Bokir", "Purwakarta", "Purwakarta"
	println(Nama, Alamat, Kota)

	var (
		name4 string  = "Muhamad Rizki"
		age4  int     = 25
		tall4 float32 = 169.9
	)
	fmt.Println(name4, age4, tall4)
}
