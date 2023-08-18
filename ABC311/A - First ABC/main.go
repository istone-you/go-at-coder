package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	var a, b, c bool

	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			a = true
		} else if s[i] == 'B' {
			b = true
		} else if s[i] == 'C' {
			c = true
		}

		if a && b && c {
			fmt.Println(i + 1)
			break
		}
	}
}
