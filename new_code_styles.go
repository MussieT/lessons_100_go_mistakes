package main

import (
	"fmt"
	"strings"
)

// string concatination with str builder
func ConcatStrings(s1 string, s2 string, max int) {
	var concat strings.Builder

	concat.WriteString(s2)
	concat.WriteString(s1)

	fmt.Println("the variable: ", concat.String()[:max])
}
