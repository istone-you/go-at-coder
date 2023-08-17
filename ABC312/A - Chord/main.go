package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s == "ACE" || s == "BDF" || s == "CEG" || s == "DFA" || s == "EGB" || s == "FAC" || s == "GBD" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
