package main

import "fmt"

func main() {
	// arithmetic operators
	a := 10
	b := 20
	// Addition pertambahan
	fmt.Println(a + b)
	// substraction pengurangan
	fmt.Println(a - b)
	// multiplication
	fmt.Println(a * b)
	// Divide
	fmt.Println(a / b)
	// modulo
	fmt.Println(a % 6)

	// Assignment operator
	count := 10
	// pertambahan
	count += 5
	println(count)
	// pengurangan
	count -= 3
	println(count)
	// perkalian
	count *= 4
	println(count)
	// pembagian
	count /= 4
	println(count)
	// modulo
	count %= 5
	println(count)

	// Comparison Operator.
	var (
		x int = 10
		y int = 15
	)
	// bigger than
	println(x > y)
	// lower than
	println(x < y)
	// bigger than or equal to
	println(x >= y)
	// lower than or equal to
	println(x <= y)
	// equal to
	println(x == y)
	// not equal to
	println(x != y)

	// logic operations
	var (
		X = true
		Y = false
	)
	// And
	println(X && X)
	println(X && Y)
	println(Y && Y)
	// Or
	println(X || X)
	println(X || Y)
	println(Y || Y)
	// negation / kebalikan
	println(!X)
	println(!Y)
}
