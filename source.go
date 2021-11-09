package godoc_generics_issue

import "fmt"

// Print prints the elements of a slice.
// It should be possible to call this with any slice value.
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Slice is not showing up in Godoc.
func Slice[T any]() []T {
	return nil
}

func Single[T any]() *T {
	return nil
}