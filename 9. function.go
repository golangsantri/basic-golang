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
	// Closure
	getMinMax := func(list []int) (min, max int) {
		for x, y := range list {
			if x == 0 {
				min, max = y, y
			} else {
				if y > max {
					max = y
				}
				if y < min {
					min = y
				}
			}
		}
		return
	}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	min, max := getMinMax(nums)
	// %v for everything, %d for number %T for float
	fmt.Printf("Min %d Max %d", min, max)
	// Immediately-Invoked Function Expression
	newNumbers := func(min int) []int {
		var r []int
		for _, y := range nums {
			if y < min {
				continue
			}
			r = append(r, y)
		}
		return r
	}(3) //input the parameters
	fmt.Println("Numbers: ", nums)
	fmt.Println("New Numbers:", newNumbers)
	// Closure as a return
	lenght, ClosureAsAreturn := ClosureAsAReturn(nums, 4)
	Value := ClosureAsAreturn()
	fmt.Println("Lenght ", lenght, "Closure As A Return: ", Value)
	// Function as a parameter
	data := []string{
		"Kibo",
		"Muhamad",
		"Rizki",
		"Arif",
		"Fadillah",
		"Bokir",
	}
	MoreThenFive := func(word string) bool {
		return len(word) > 5
	}
	result := Filter(data, MoreThenFive)
	fmt.Println("Result: ", result)
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

// Closure as a return
func ClosureAsAReturn(num []int, max int) (int, func() []int) {
	var r []int
	for _, y := range num {
		if y > max {
			r = append(r, y)
		}
	}
	return len(r), func() []int {
		return r
	}
}

// Function as a parameter
func Filter(data []string, callback func(string) bool) []string {
	result := []string{}
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// Alias Schema Closure
type FilterCallBack func(string) bool

// By using the alias we can short our code easely
