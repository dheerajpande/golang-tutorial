package main

import (
	"fmt"
)

func main() {
	var a [10]int
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)

	var twoDarray [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twoDarray[i][j] = i*10 + j
		}
	}
	fmt.Println(twoDarray)
}
