package main

import (
	"fmt"
)

func main() {
	var n, a, b int

	fmt.Scan(&n, &a, &b)

	var result int

	for i := 1; i <= n; i++ {
		sum := calculateDigitSum(i)
		if a <= sum && sum <= b {
			result += i
		}
	}
	fmt.Println(result)
}

func calculateDigitSum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
