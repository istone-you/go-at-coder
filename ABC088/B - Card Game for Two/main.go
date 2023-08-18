package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var alice int
	var bob int

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for i, v := range a {
		if i%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}
	fmt.Println(alice - bob)
}
