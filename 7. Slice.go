package main

import "fmt"

func main() {
	slice := []string{
		"Muhamad",
		"Rizki",
		"Arif",
		"Fadillah",
	}
	fmt.Println(slice)
	for _, y := range slice {
		fmt.Println(y)
	}
	// indexing
	fmt.Println("Indexing")
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	// Slicing
	fmt.Println("Slicing")
	fmt.Println(slice[2:])
	fmt.Println(slice[:3])
	fmt.Println(slice[:4])

	// Slice is references of data type
	source := []int{1, 2, 3, 4, 5, 6, 7}
	branch_1 := source
	branch_2 := branch_1
	fmt.Println("Source :", source)
	fmt.Println("Branch 1 :", branch_1)
	fmt.Println("Branch 2: ", branch_2)
	branch_2[3] = 10
	fmt.Println("Source :", source)
	fmt.Println("Branch 1 :", branch_1)
	fmt.Println("Branch 2: ", branch_2)
	// if the branch of source is changed. The souce will be changed also

	// len function is to calculate the number of indexs
	fmt.Println("Len", len(slice))
	// cap is to count the maximun of capacity
	fmt.Println("Cap", cap(slice))
	// append is used for insert the value to the slice
	empty := []int{}
	for x := 10; x > 0; x-- {
		empty = append(empty, x)
	}
	fmt.Println("Empty slice: ", empty)
}
