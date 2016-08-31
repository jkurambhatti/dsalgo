package main

import (
	"fmt"
	//"math"
)
//
//func prime(p float64) bool {
//	for i := 2; i <= int(math.Sqrt(p)); i++ {
//		if int(p)%i == 0 {
//			return false
//		}
//	}
//	return true
//}
//
//func createArrayPrime(N int) []int {
//	var arrprm = make([]int, 0)
//	for i := 2; i < N; i++ {
//		if prime(float64(i)) {
//			arrprm = append(arrprm, i)
//		}
//	}
//	return arrprm
//}
//
//func nooffacts(num int) int  {
//	var nof,i int
//	arrprmno := createArrayPrime(30)
//
//	if num == 0 || num == 1{
//		return nof
//	}
//
//
//	for  {
//		if num % arrprmno[i] == 0 {
//			nof++
//			if num > arrprmno[i]{
//				num = num / arrprmno[i]
//			} else {
//				break
//			}
//		} else {
//			i+=1
//		}
//	}
//	return nof
//}

func recprime1(num,i int) int  {
	if num != 1 {
		if num % i == 0{
			return 1 + recprime1(num/i,i)
		} else {
			return recprime1(num,i+1)
		}
	}
	return 0
}

func findkp(k, N int) {
	fmt.Println(k, N)
	count := 0
	for number := 2;count < N;number+=1 {
		if recprime1(number,2) == k {
			fmt.Println(number)
			count+=1
		}
	}

}

func main() {
	findkp(7, 5)
}
