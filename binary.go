package main

import "fmt"

func convBin(n int) {
	if n != 0 {
		convBin(n / 2)
		fmt.Printf("%d", n%2)
	}
	return
}

func main() {
	convBin(6)
}
