package main

type CustomConstraint interface {
	~int | ~string
}

func GetKeys[K CustomConstraint, V any] (m map[K]V) []K {
	// implementation
	return nil
}
