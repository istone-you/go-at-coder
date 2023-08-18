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

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	for i, v := range a {
		if i%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}
	fmt.Println(alice - bob)
}
