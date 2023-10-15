package main

import (
	"unsafe"
)

func main() {
	const (
		Unknown = 0
		Female  = 1
		Male    = 1
	)

	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

	println(Unknown, Female, Male, a, b, c)
}
