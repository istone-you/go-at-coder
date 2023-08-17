package main

import "fmt"

func main() {
	var pi string = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	var number int

	fmt.Scan(&number)

	fmt.Println(pi[:number+2])
}
