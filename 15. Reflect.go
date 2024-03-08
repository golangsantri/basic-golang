package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 25
	reflectValue := reflect.ValueOf(number)
	fmt.Println("Number data type: ", reflectValue.Type())
	fmt.Println("Value of number: ",reflectValue.Interface().(int))
	name := "Muhamad Rizki Arif Fadillah"
	reflectValue = reflect.ValueOf(name)
	fmt.Println("Name data type: ", reflectValue.Type())
	// always get error. I'm still confused how to resolve it
	fmt.Println("Value of number: ",reflectValue.Interface().(string))
	// Access to object
	s1 := &Student{
		id: "20",
		name: "Muhamad Rizki",
	}
	s1.getPropertyInfo()
}
type Student struct {
	id string
	name string
}
func (s *Student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
	
		reflectValue = reflectValue.Elem()
	
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
	
		fmt.Println("nama:", reflectType.Field(i).Name)
	
		fmt.Println("tipe data :", reflectType.Field(i).Type)
	
		fmt.Println("nilai:", reflectValue.Field(i).Interface().(string))
	
		fmt.Println("")
	}
	}