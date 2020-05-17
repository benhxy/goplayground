package landmine

import (
	"fmt"
	"testing"
)

func TestUnderlyingArray(t *testing.T) {
	// When a slice is given capacity, an array with length = capacity is created.
	// When append an element to the slice, as long as the size is within the
	// underlying array's length, the returned slice will keep referencing to the same array.
	emptySliceWithNonZeroCapacity := make([]int, 0, 1)

	a := append(emptySliceWithNonZeroCapacity, 100)
	fmt.Printf("a's address=%p\n", a)

	b := append(emptySliceWithNonZeroCapacity, 200)
	fmt.Printf("b's address=%p\n", b)

	fmt.Printf("a = %v, b = %v\n", a, b)

	// If the slice is assigned 0 capacity, then a new array needs to be created
	// when append() is called. Causing c and d to be slices with different underlying arrays.
	emptySliceWithZeroCapacity := make([]int, 0)

	c := append(emptySliceWithZeroCapacity, 100)
	fmt.Printf("c's address=%p\n", c)

	d := append(emptySliceWithZeroCapacity, 200)
	fmt.Printf("d's address=%p\n", d)

	fmt.Printf("c = %v, d = %v\n", c, d)

	/*
		a's address=0xc0000142f8
		b's address=0xc0000142f8
		a = [200], b = [200]

		c's address=0xc000014330
		d's address=0xc000014338
		c = [100], d = [200]
	*/
}
