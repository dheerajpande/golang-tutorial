package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

func (r *Rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	fmt.Println("Welcome")
	rect := Rectangle{10, 20}
	rect_ptr := &rect
	fmt.Println("Area of rectangle:", rect.area())
	fmt.Println("Perimeter of rectangle: ", rect_ptr.perimeter())
}
