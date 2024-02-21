package main

import "fmt"

func main() {
	// Pointer is the reference data type
	// asterik
	var NumberA = 5
	var NumberB *int = &NumberA
	fmt.Println("Value of A: ", NumberA)
	fmt.Println("The Memori Address of A:", &NumberA)
	// With asterisk
	fmt.Println("Pointer: Value ", *NumberB)
	fmt.Println("Pointer: Address ", NumberB)
	// Changing the value, Jika salah satu data pointer diubah maka dia merubah kesemua variable yang berkaitan.
	NumberC := &NumberA
	fmt.Printf("A: %v \nB: %v\nC: %v", NumberA, *NumberB, *NumberC)
	*NumberC = 10
	fmt.Printf("\nA: %v \nB: %v\nC: %v", NumberA, *NumberB, *NumberC)
	// Pointer Parameter
	NumD := 10
	Change(&NumD, 20)
}

// Function need pointer parameter
func Change(src *int, value int) {
	*src = value
}
