package main

import "fmt"

func main() {
	// map is like dictionary in python
	student := map[string]int{
		"Key_1": 1,
		"Key_2": 2,
		"Key_3": 3,
		"Key_4": 4,
		"Key_5": 15,
	}
	fmt.Println("Student map :", student)
	// student2 := make(map[string]string)
	// student3 := *new(map[string]string)
	for x, y := range student {
		fmt.Println("Key", x, "Value", y)
	}
	// delete an item in map
	delete(student, "Key_3")
	fmt.Println("Student", student)
	// Combination between slice and map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["name"], chicken["gender"])
	}
}
