package main

//import "math"
//
//func prime(p float64) bool {
//	for i := 2; i <= int(math.Sqrt(p)); i++ {
//		if int(p)%i == 0 {
//			return false
//		}
//	}
//	return true
//}

func recprime2(num,i int) int  {
	if num != 1 {
		if num % i == 0{
			return 1 + recprime2(num/i,i)
		} else {
			return recprime2(num,i+1)
		}
	}
	return 0
}

func main() {
	println(recprime2(16,2))
}
