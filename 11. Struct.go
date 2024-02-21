package main

import "fmt"

// private variable
type student struct {
	Id   uint
	Name string
}

// Embed struct
type Employee struct {
	Id uint
	Person
	Status string
	Salary uint
}
type Person struct {
	Name    string
	Address string
}

func main() {
	// Public Variable
	FirstStudent := student{
		Id:   1,
		Name: "Kibo",
	}
	fmt.Println(FirstStudent.Id, FirstStudent.Name, FirstStudent)
	// Pointer in struct
	SecondStudent := &FirstStudent
	fmt.Println(FirstStudent.Id, FirstStudent.Name, FirstStudent)
	fmt.Println(SecondStudent.Id, SecondStudent.Name, SecondStudent)
	SecondStudent.Id = 10
	fmt.Println(FirstStudent.Id, FirstStudent.Name, FirstStudent)
	fmt.Println(SecondStudent.Id, SecondStudent.Name, SecondStudent)
	// Embed struct
	NewEmployee := Employee{
		Id: 1,
		Person: Person{
			Name:    "Kibo",
			Address: "Purwakarta",
		},
		Status: "Contract",
		Salary: 10000000,
	}
	fmt.Println("New Employee: ", NewEmployee)
	fmt.Printf("New Employee:\nID: %v\nName: %v\nAddress: %v\nStatus: %v\nSalary: %v\n", NewEmployee.Id, NewEmployee.Person.Name, NewEmployee.Person.Address, NewEmployee.Status, NewEmployee.Salary)
	// Anonymous struct
	anonStruct := struct {
		id   uint
		name string
	}{
		// input value
		1,
		"kibo",
	}
	fmt.Println("Anonymous Struct: ", anonStruct)
}
