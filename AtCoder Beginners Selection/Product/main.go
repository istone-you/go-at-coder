package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanf("%d %d", &a, &b)

	result := a * b % 2

	if result == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
