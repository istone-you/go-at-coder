package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	number := 0

finished:
	for {
		for i := 0; i < len(a); i++ {
			if a[i]%2 != 0 {
				break finished
			}
			a[i] = a[i] / 2
		}
		number++
	}
	fmt.Println(number)
}
