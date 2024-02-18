package main

import "fmt"

func main() {
	// first array
	array := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(array)
	for _, y := range array {
		fmt.Println(y)
	}
	// initiate array withoun the number of elements
	array2 := [...]int{
		3, 4, 5, 6, 7, 8, 9,
	}
	fmt.Println(array2)
	fmt.Println(len(array2))
	// Array multi dimensions
	vector := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(vector)
	fmt.Println(len(vector))
	fmt.Println(len(vector[0]))
	// make
	var array3 = make([]int, 3)
	array3[0] = 1
	array3[1] = 3
	array3[2] = 2

	x := 1
	for {
		if x > 4 {
			break
		}
		array3 = append(array3, x)
		x++

	}
	fmt.Println(array3)
}
