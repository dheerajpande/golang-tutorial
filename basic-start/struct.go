package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string) *Person {
	p := Person{name: name}
	fmt.Println(p)
	return &p
}
func main() {
	p := NewPerson("Shubham")
	fmt.Println(*p)
	p1 := Person{name: "New Person", age: 50}
	fmt.Println(p1)
}
