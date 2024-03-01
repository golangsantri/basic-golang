package main

import (
	"fmt"
)

func main() {
	// first looping
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
	// second looping
	for i := 1; i < 10; {
		fmt.Println(i)
		if i == 3 {
			break
		}
		i++
	}
	// Third looping
	z := 5
	for z < 10 {
		fmt.Println(z)
		z++
	}
	// Looping without argument
	z = 0
	for {
		fmt.Println(z)
		if z == 10 {
			break
		}
		z++
	}
	// Looping with index and value
	name := "Muhamad Rizki Arif Fadillah"
	for i, v := range name {
		fmt.Printf("Index %d, value %v \n", i, string(v))
	}
	// Looping for array or slice
	array := [5]int{1, 2, 3, 4, 5}
	for a, b := range array {
		fmt.Println(a, b)
	}
	// Looping for map
	student := map[string]string{
		"a": "Kibo",
		"b": "Caca",
		"c": "kibo",
	}
	for z, x := range student {
		fmt.Println(z, x)
	}
	// breeak and continue
	for z := 0; z < 20; z++ {
		if z == 5 {
			continue
		}
		if z == 10 {
			break
		}
		fmt.Println(z)
	}
	// Nested looping
	for x := 1; x < 10; x++ {
		for y := x; y < 10; y++ {
			fmt.Print(y)
		}
		fmt.Println()
	}
}
