package main

import (
	"fmt"
)

func main() {
	var n, p, q int
	fmt.Scan(&n, &p, &q)

	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}

	var min int = d[0]
	for _, v := range d {
		if min >= v {
			min = v
		}
	}

	if p < q+min {
		fmt.Println(p)
	} else {
		fmt.Println(q + min)
	}

}
