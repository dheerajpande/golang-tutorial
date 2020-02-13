package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func factorial(num int) int {
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	return ans
}

func selection_sort(a []int) {
	l := len(a)
	var temp int
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		temp = a[i]
		a[i] = a[min]
		a[min] = temp
	}
}

func get_max_2(a []int) (int, int) {
	m1 := a[0]
	m2 := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m1 {
			m2 = m1
			m1 = a[i]
		} else if a[i] > m2 {
			m2 = a[i]
		}
	}
	return m1, m2
}

func sum_n(a ...int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println("functions.go  demo")
	a := []int{1, 2, 30, 14, 13, 6, 7}
	fmt.Println(a)
	m1, m2 := get_max_2(a)
	fmt.Println(m1, m2)
	fmt.Println(sum_n(1, 2, 3, 5))
}
