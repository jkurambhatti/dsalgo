// sum of elements of an array using recursion
package main

func sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + sum(a[1:])
}

func main() {
	println(sum([]int{1, 2, 3, 4, 5, 6}))
}
