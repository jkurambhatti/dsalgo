// greatest common divisor using euclidian method

package main

import "fmt"

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a%b == 0 {
		return b
	} else if b%a == 0 {
		return a
	} else if a > b {
		return gcd(a%b, b)
	} else {
		return gcd(a, b%a)
	}
}

func main() {
	fmt.Println(gcd(17, 19))
}
