package main

import (
	"fmt"
	"math/rand"
)

func main() {
	value := rand.Int31n(10)
	if value <= 8 && value > 6 {
		fmt.Println(value)
		fmt.Println("Good Job")
	} else if value <= 10 && value > 8 {
		fmt.Println(value)
		fmt.Println("Great Value")
	} else {
		fmt.Println(value)
		fmt.Println("Try again")
	}

	// Temporary variables in golang
	point := 8840
	if percent := point / 100; percent >= 100 {
		fmt.Println("Perfect")
	} else if percent >= 80 && percent < 100 {
		fmt.Println("Good Job")
	} else {
		fmt.Println("Try again")
	}

	// Switch

	count := 6
	switch count {
	case 8:
		fmt.Println("Great")
	case 6:
		fmt.Println("Good job")
	default:
		fmt.Println("Not bad")
	}

	// case with lots of conditions

	count = rand.Intn(10)
	switch count {
	case 10:
		fmt.Println("Perfect")
	case 9, 8, 7, 6, 5:
		fmt.Println("Good Job")
	default:
		fmt.Println("Try again")
	}

	// Nested Condition
	var num int = rand.Intn(10)
	if num >= 7 {
		switch num {
		case 8:
			fmt.Println(num, "Great")
		case 9:
			fmt.Println(num, "Too Great")
		default:
			fmt.Println("Default")
		}
	} else {
		if num == 6 {
			fmt.Println(num)
		} else if num == 5 {
			fmt.Println(num)
		} else {
			fmt.Println(num)
		}
	}
}
