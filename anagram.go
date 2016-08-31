// to see if the strings are anagrams

package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	sa1 := strings.Split(strings.Join(strings.Fields(s1), ""), "")
	sa2 := strings.Split(strings.Join(strings.Fields(s2), ""), "")
	sort.Strings(sa1)
	sort.Strings(sa2)
	return reflect.DeepEqual(sa1, sa2)
}

func main() {
	fmt.Println(isAnagram("amit mishra", "mxaria    smi     th "))
}
