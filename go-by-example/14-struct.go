package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name, age}
	return &p
}
func main() {
	fmt.Println(person{"Hari", 20})
	fmt.Println(person{"Shyam", 21})
	fmt.Println(newPerson("subina", 21))
	p := person{name: "syantang", age: 21}
	p.age = 23
	fmt.Println("Name: ", p.name, " Age: ", p.age)
	p1 := &p
	p1.age = 22
	fmt.Println("Name: ", p1.name, " Age: ", p1.age)
	fmt.Println("Name: ", p.name, " Age: ", p.age)
}
