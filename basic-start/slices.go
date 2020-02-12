package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Slices")
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(len(s), s)
	s = append(s, "Random string")
	fmt.Println(s, len(s), s[1:3])
}
