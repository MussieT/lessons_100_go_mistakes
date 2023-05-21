package main

import (
	"fmt"
)

// we can have multiple init functions
// don't insert things with concrete implementation (like db connection, or anything that can break the server..) hard to unit test with mock ..
func init() {
	fmt.Println("first init")
}

func init() {
	fmt.Println("second init")
}

func main() {
	// Variable shadowing
	ShadowedVariable()
	UnShadowedVariable()

	// Style
	ConcatStrings("one", "love", 7)

	// interfaces
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	// map_to_struct
}
