/*
 Given a string containing of ‘0’, ‘1’ and ‘?’ wildcard characters, generate all binary strings
 that can be formed by replacing each wildcard character by ‘0’ or ‘1’.
*/

package main

import (
	"fmt"
	"strings"
)

var arrstr = make([]string, 0)

func replace(str string, index int, sub string) string {
	as := strings.Split(str, "")
	as[index] = sub
	return strings.Join(as, "")
}

func biper(s string) {
	fmt.Println(s)
	if s == "" {
		return
	}
	index := strings.Index(s, "?")
	if index == -1 {
		return
	} else {
		s = replace(s, index, "0")
		if strings.Index(s, "?") == -1 {
			arrstr = append(arrstr, s)
		}
		biper(s)
		//fmt.Println(s)
		s = replace(s, index, "1")
		if strings.Index(s, "?") == -1 {
			arrstr = append(arrstr, s)
		}
		biper(s)
	}
}

func main() {
	biper("??")
	fmt.Println(arrstr)
}
