package main

import (
	"fmt"
)

func main() {
	// loop format 1 with start and end condition
	i := 1
	for i = 1; i < 10; i++ {
		fmt.Println("Value of i is", i)
	}

	// for with no condition
	for {
		fmt.Println("Loop with break")
		break
	}

	// or with single condition
	for i < 15 {
		fmt.Println("For loop with single condition ", i)
		i++
	}
	// for loop with continue
	for i = 0; i < 15; i++ {
		if i%2 == 0 {
			// skip even numbers
			continue
		}
		fmt.Println("The number is ", i)
	}
}
