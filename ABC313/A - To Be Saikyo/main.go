package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	var max int
	for _, s := range a {
		if max < s {
			max = s
		}
	}
	if max != a[0] {
		fmt.Println((max + 1) - a[0])
	} else {
		var b int
		for i := 1; i < n; i++ {
			if a[0] == a[i] {
				b = 1
				break
			}
		}
		fmt.Println(b)
	}
}
