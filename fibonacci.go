// sum of first n elements in a fibonacci series
// create an array of n fibonacci elements

package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibarr(n int) []int {
	var a = make([]int, 0)
	for i := 0; i < n; i++ {
		a = append(a, fib(i))
	}
	return a
}

func main() {
	fmt.Println(fibarr(10))
}
