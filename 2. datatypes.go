package main

import (
	"fmt"
)

func main() {
	// Data types in golang
	// string
	nama := "Muhamad Rizki"
	fmt.Println(nama)
	// int,int8,16,32,64
	// bisa bernilai minus -1 -
	usia := -40
	fmt.Println(usia)
	// uint,uint8,16,32,64
	// tidak bisa minus
	var stokBarang uint = 60
	fmt.Println(stokBarang)
	// float32,64
	var tinggi float32 = 17.6
	fmt.Println(tinggi)
	// bool
	// true/false
	handsome := true
	fmt.Println("Am I handosome?", handsome)
}
