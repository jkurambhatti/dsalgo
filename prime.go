package main

import "math"

func prime(p float64) bool {
	for i := 2; i <= int(math.Sqrt(p)); i++ {
		if int(p)%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	println(prime(11))
}
