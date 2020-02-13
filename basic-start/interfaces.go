package main

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Given geometry is: ", reflect.TypeOf(g), g)
	fmt.Println("area of: ", reflect.TypeOf(g), g.area())
	fmt.Println("perimeter of: ", reflect.TypeOf(g), g.perimeter())
}

func main() {
	r := rectangle{3, 4}
	c := circle{5}
	measure(r)
	measure(c)
}
