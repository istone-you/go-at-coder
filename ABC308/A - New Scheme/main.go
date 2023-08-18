package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f, g, h int

	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h)

	s := []int{a, b, c, d, e, f, g, h}

	var result bool = true

	for i := 0; i < 8; i++ {
		if i < 7 && s[i] > s[i+1] {
			result = false
		}

		if s[i] < 100 || s[i] > 675 {
			result = false
		}

		if s[i]%25 != 0 {
			result = false
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
