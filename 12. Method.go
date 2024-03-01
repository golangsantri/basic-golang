package main

import (
	"fmt"
	"strings"
)

func main() {
	owner := Owner{
		Id: 1,
		Name: "Muhamad Rizki Arif Fadillah",
	}
	owner.sayHelloW()
	own := owner.getNameAt(2)
	fmt.Println("Get: ",own)
	// Method Pointer
	fmt.Println("Owner: ", owner)
	owner.changeName("Kibo")
	fmt.Println("Owner: ", owner)
	owner.ChangeNamePointer("susu")
	fmt.Println("Owner: ", owner)

}

type Owner struct {
	Id uint
	Name string
}

func (O Owner) sayHelloW()  {
	fmt.Println("Hello",O.Name)
}
func (o Owner) getNameAt(i int)string  {
	return strings.Split(o.Name," ")[i-1]
}
func (o Owner) changeName(name string)  {
	fmt.Println("Changed name to ->",name)
	o.Name = name
}
func (o *Owner) ChangeNamePointer(name string)  {
	fmt.Println("Changed name to ->",name)
	o.Name = name
}