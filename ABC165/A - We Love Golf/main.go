package main

import (
	"fmt"
)

func main() {
	var h, a, b int
	fmt.Scan(&h, &a, &b)

	var result string = "NG"

	for i := 0; i*h <= b; i++ {
		if i*h >= a {
			result = "OK"
		}
	}

	fmt.Println(result)
}
