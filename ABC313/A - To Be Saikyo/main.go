package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	var max int
	for _, v := range p {
		if max < v {
			max = v
		}
	}
	var ans int

	if max == p[0] {
		ans = 0
		for _, v := range p[1:] {
			if max == v {
				ans = 1
				break
			}
		}
	} else {
		ans = (max + 1) - p[0]
	}

	fmt.Println(ans)
}
