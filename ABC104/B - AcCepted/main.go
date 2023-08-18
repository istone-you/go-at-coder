package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	c := s[2 : len(s)-1]
	count := strings.Count(c, "C")

	var result bool

	if count == 1 && s[0] == 'A' {
		result = true
		c = s[1:]
		c = strings.Replace(c, "C", "", -1)
		for i := 0; i < len(c); i++ {
			if unicode.IsUpper(rune(c[i])) {
				result = false
			}
		}
	}

	if result {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
