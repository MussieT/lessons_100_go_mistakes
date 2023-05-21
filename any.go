package main

// any is an empty interface .. accepts everything
func UseAny() {
	var i any

	i = 42;

	i = "foo"

	i = struct{
		s string
	} {
		s: "bar",
	}

	i = f

	_ = i
}

func f() {}

// By using any, we lose some of the benefits of Go as a statically type language. Instead
// we should avoid any types and make our signatures explicit as much as possible.


// WHEN IT IS HELPFUL

// The standard encoding/json library, has a marshal function (Anything can be marshalled)
func Marshal(v any) ([]byte, error) {
	// ....
	var incoming any

	return incoming.([]byte), nil
}

