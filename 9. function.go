package main

import "fmt"

func main() {
	// Basic function
	HelloWorld()
	// Return function
	fmt.Println(ReturnFunction())
	// Parameter function
	fmt.Println(GetName("Kibo"))
	// Multiple parameters
	fmt.Println(GetAddress("Kibo", "Purwakarta", 24))
	// Multiple return
	fmt.Println(MultipleReturn("Muhamad RIzki", "Kibo", 24))
	// Variadic function
	fmt.Println("Variadid Function", VariadicFunction(1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 12, 13, 14, 15, 16, 17, 19, 10))
}

// basic function
func HelloWorld() {
	fmt.Println("Hello World!")
}

// return function
func ReturnFunction() string {
	return "Hello World"
}

// 1 parameter
func GetName(name string) string {
	return "Hello! " + name
}

// Multiple parameters
func GetAddress(name, address string, age uint) string {
	return name + string(rune(uint(age))) + address
}

// Multiple return
func MultipleReturn(name, nickname string, age int) (string, uint) {
	return name + nickname, uint(age)
}

// Variadic Function
func VariadicFunction(numbers ...int) int {
	result := 0
	if len(numbers) <= 1 {
		return numbers[0]
	}
	for _, value := range numbers {
		result += value
	}
	return result
}
