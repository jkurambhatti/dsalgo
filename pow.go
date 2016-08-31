// calculate positive and negative power of a number

package main

import "fmt"

func pow(a, b float64) float64 {
	if b == 0 {
		return 1
	} else if b < 0 {
		return (1 / a) * pow(a, b+1)
	} else {
		return a * pow(a, b-1)
	}
}

func main() {
	fmt.Printf("%.2f", pow(2, -1))
}
