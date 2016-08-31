/*
 Given an array of 0's and 1's your task is to complete the function maxLen which returns
 size of  the  largest sub array  with equal number of 0's and 1's .
 The function maxLen takes 2 arguments .
 The first argument is the array A[] and second argument is the size 'N' of the array A[] .
*/

package main

import (
	"fmt"
)

var largSubArrlen int

func count01s(a []int) {
	var count = make(map[int]int)

	if len(a)%2 == 0 {
		for _, v := range a {
			count[v]++
		}
		if count[0] == count[1] {
			if largSubArrlen < (2 * count[0]) {
				largSubArrlen = 2 * count[0]
				fmt.Println(a)
			}
		}
	}
}

func maxLen(a []int) int {

	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			count01s(a[i : j+1])
		}
	}

	return largSubArrlen
}

func main() {
	var arr = make([]int, 0)
	arr = append(arr, 0, 1, 0, 1, 0, 0, 1)
	fmt.Println(maxLen(arr))
}
