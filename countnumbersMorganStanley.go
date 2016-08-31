package main

import "fmt"

func isRange(n int) bool {
	switch n {
	case 1:
		return true
	case 2:
		return true
	case 3:
		return true
	case 4:
		return true
	case 5:
		return true
	default:
		return false
	}
}

func count(N int) []int {
	var flag = 0
	var arr = make([]int, 0)
	for i := 1; i <= N; i++ {
		for j := i; j != 0; j = j / 10 {
			if isRange(j % 10) {
				flag = 1
			} else {
				flag = 0
				break
			}
		}
		if flag == 1 {
			arr = append(arr, i)
		}
		flag = 0
	}
	return arr
}

func main() {
	fmt.Println(count(777))
}
