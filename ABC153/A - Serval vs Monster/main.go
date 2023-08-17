package main

import (
	"fmt"
)

func main() {
	var h, a int
	fmt.Scan(&h, &a)

	var count int

	for h > 0 {
		h -= a
		count++
	}
	fmt.Println(count)
}
