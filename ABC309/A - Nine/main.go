package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	var ans bool

	if a == 1 && b == 2 {
		ans = true
	}
	if a == 2 && b == 3 {
		ans = true
	}
	if a == 4 && b == 5 {
		ans = true
	}
	if a == 5 && b == 6 {
		ans = true
	}
	if a == 7 && b == 8 {
		ans = true
	}
	if a == 8 && b == 9 {
		ans = true
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
