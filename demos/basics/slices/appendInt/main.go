package main

import "fmt"

// This is for demo purposes, you should use built-in `append` instead
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		// z's underlying array is still x
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if len(x) != 0 {
			zcap = 2 * len(x)
		}

		// z's underlying array points to a new array
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}

	// Append y
	z[len(x)] = y

	// Return slice
	return z
}

func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
